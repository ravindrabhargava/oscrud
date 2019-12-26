package socketio

import (
	"fmt"
	"net/http"
	"oscrud/endpoint"
	"oscrud/service"

	engineio "github.com/googollee/go-engine.io"
	socketio "github.com/googollee/go-socket.io"
)

// Transport :
type Transport struct {
	Port   int
	Socket *socketio.Server
}

// SocketContext :
type SocketContext struct {
	Query  map[string]interface{} `json:"query"`
	Body   map[string]interface{} `json:"body"`
	Header map[string]interface{} `json:"header"`
	Param  map[string]string      `json:"param"`
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

// RegisterService :
func (t *Transport) RegisterService(srv string, route service.Route) {
	t.Socket.OnEvent(
		"/", srv+"."+route.Action,
		func(socket socketio.Conn, object string) string {
			return "SERVICE"
		},
	)
}

// RegisterEndpoint :
func (t *Transport) RegisterEndpoint(endpoint string, route endpoint.Route) {
	t.Socket.OnEvent(
		"/", endpoint,
		func(socket socketio.Conn, object string) string {
			return "ENDPOINT"
		},
	)
}

// Start :
func (t *Transport) Start() error {
	defer t.Socket.Close()
	go t.Socket.Serve()
	http.Handle("/socket.io/", t.Socket)
	port := fmt.Sprintf(":%d", t.Port)
	return http.ListenAndServe(port, nil)
}
