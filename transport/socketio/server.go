package socketio

import (
	"fmt"
	"net/http"
	"oscrud/action"

	engineio "github.com/googollee/go-engine.io"
	socketio "github.com/googollee/go-socket.io"
)

// Transport :
type Transport struct {
	Port   int
	Socket *socketio.Server
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
func (t *Transport) RegisterService(service string, route action.ServiceRoute) {

}

// RegisterEndpoint :
func (t *Transport) RegisterEndpoint(endpoint string, route action.EndpointRoute) {

}

// Start :
func (t *Transport) Start() error {
	defer t.Socket.Close()
	go t.Socket.Serve()
	http.Handle("/socket.io/", t.Socket)
	port := fmt.Sprintf(":%d", t.Port)
	return http.ListenAndServe(port, nil)
}
