package oscrud

import (
	"oscrud/action"
	"oscrud/transport"
	"reflect"
)

// Oscrud :
type Oscrud struct {
	Transports []transport.Transport
	Routes     map[string]map[string]interface{}
}

// NewOscrud :
func NewOscrud() *Oscrud {
	return &Oscrud{
		Transports: make([]transport.Transport, 0),
		Routes: map[string]map[string]interface{}{
			"GET":    map[string]interface{}{},
			"POST":   map[string]interface{}{},
			"PUT":    map[string]interface{}{},
			"PATCH":  map[string]interface{}{},
			"DELETE": map[string]interface{}{},
		},
	}
}

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...transport.Transport) *Oscrud {
	for _, trs := range transports {
		server.Transports = append(server.Transports, trs)
	}
	return server
}

// RegisterEndpoint : ( Even Index )
func (server *Oscrud) RegisterEndpoint(method string, path string, endpoint action.Endpoint) *Oscrud {
	server.Routes[method][path] = endpoint.Action
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(basePath string, service action.Service) *Oscrud {
	server.Routes["GET"][basePath] = service.Find
	server.Routes["POST"][basePath] = service.Create
	server.Routes["GET"][basePath+"/:id"] = service.Get
	server.Routes["PUT"][basePath+"/:id"] = service.Update
	server.Routes["PATCH"][basePath+"/:id"] = service.Patch
	server.Routes["DELETE"][basePath+"/:id"] = service.Remove
	return server
}

// Handler :
func (server *Oscrud) Handler(handler interface{}) action.Handler {
	return func(ctx action.EndpointContext, srv action.ServiceContext) {
		fn := reflect.ValueOf(handler)
		in := make([]reflect.Value, 0)
		if fn.Type().NumIn() == 1 {
			in = append(in, reflect.ValueOf(ctx))
		} else {
			in = append(in, reflect.ValueOf(srv))
			in = append(in, reflect.ValueOf(ctx))
		}
		fn.Call(in)
	}
}

// Start :
func (server *Oscrud) Start() {
	for _, trs := range server.Transports {

		for method, pathHandler := range server.Routes {
			for path, handler := range pathHandler {
				trs.Register(method, path, server.Handler(handler))
			}
		}

		go func(t transport.Transport) {
			err := t.Start(server.Handler)
			if err != nil {
				panic(err)
			}
		}(trs)
	}
	select {}
}
