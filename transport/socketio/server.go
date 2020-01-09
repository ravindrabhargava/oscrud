package socketio

import (
	"encoding/json"
	"fmt"
	"net/http"
	"oscrud"

	engineio "github.com/googollee/go-engine.io"
	socketio "github.com/googollee/go-socket.io"
)

// Transport Definition
var (
	TransportName = "SOCKETIO"
)

// Transport :
type Transport struct {
	Port   int
	Socket *socketio.Server
}

// SocketObject :
type SocketObject struct {
	Body   map[string]interface{} `json:"body"`
	Query  map[string]interface{} `json:"query"`
	Header map[string]string      `json:"headers"`
}

// NewSocket :
func NewSocket(opts *engineio.Options) *Transport {
	socket, err := socketio.NewServer(opts)
	if err != nil {
		panic(err)
	}

	return &Transport{
		Socket: socket,
		Port:   3000,
	}
}

// UsePort :
func (t *Transport) UsePort(port int) *Transport {
	t.Port = port
	return t
}

// Name :
func (t *Transport) Name() string {
	return TransportName
}

// Register :
func (t *Transport) Register(method string, endpoint string, handler oscrud.TransportHandler) {
	// log.Println(endpoint)
	// t.Socket.OnEvent(
	// 	"/", endpoint,
	// 	func(socket socketio.Conn, object string) string {
	// 		sobject := new(SocketObject)
	// 		if err := json.Unmarshal([]byte(object), sobject); err != nil {
	// 			panic(err)
	// 		}

	// 		req := oscrud.NewRequest(method, endpoint).
	// 			Transport(t).
	// 			Context(socket).
	// 			SetBody(sobject.Body).
	// 			SetQuery(sobject.Query).
	// 			SetHeader(sobject.Header)

	// 		result, exception := handler(req)
	// 		if exception != nil {
	// 			return parseError(exception)
	// 		}
	// 		return parseResult(result)
	// 	},
	// )
}

// Start :
func (t *Transport) Start(handler oscrud.TransportHandler) error {
	defer t.Socket.Close()
	go t.Socket.Serve()

	t.Socket.OnEvent(
		"/", "endpoint",
		func(socket socketio.Conn, endpoint string, object string) string {
			sobject := new(SocketObject)
			if err := json.Unmarshal([]byte(object), sobject); err != nil {
				panic(err)
			}

			req := oscrud.NewRequest("*", endpoint).
				Transport(t).
				Context(socket).
				SetBody(sobject.Body).
				SetQuery(sobject.Query).
				SetHeader(sobject.Header)

			response := handler(req)
			return parseResponse(response)
		},
	)

	http.Handle("/socket.io/", t.Socket)
	port := fmt.Sprintf(":%d", t.Port)
	return http.ListenAndServe(port, nil)
}
