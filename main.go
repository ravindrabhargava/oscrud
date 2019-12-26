package oscrud

import (
	"oscrud/endpoint"
	"oscrud/service"
	"oscrud/transport"
	"strings"
)

// Oscrud :
type Oscrud struct {
	Transports []transport.Transport
	Services   map[string]service.Route
	Endpoints  map[string]endpoint.Route
}

// Service :
type Service struct {
	server  *Oscrud
	service string
}

// NewOscrud :
func NewOscrud() *Oscrud {
	return &Oscrud{
		Transports: make([]transport.Transport, 0),
		Services:   make(map[string]service.Route),
		Endpoints:  make(map[string]endpoint.Route),
	}
}

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...transport.Transport) *Oscrud {
	for _, trs := range transports {
		server.Transports = append(server.Transports, trs)
	}
	return server
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(key, method, route string, handler endpoint.Handler) *Oscrud {
	server.Endpoints[key] = endpoint.Route{
		Endpoint: key,
		Method:   method,
		Path:     strings.TrimPrefix(route, "/"),
		Handler:  handler,
	}
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(key, basepath string, class service.Service) *Oscrud {
	path := strings.TrimPrefix(basepath, "/")
	findKey := key + ".find"
	createKey := key + ".create"
	getKey := key + ".get"
	updateKey := key + ".update"
	patchKey := key + ".patch"
	deleteKey := key + ".remove"

	server.Services[findKey] = service.Route{
		Service: key,
		Action:  "find",
		Method:  "get",
		Path:    path,
		Handler: class.Find,
	}

	server.Services[getKey] = service.Route{
		Service: key,
		Action:  "get",
		Method:  "get",
		Path:    path,
		Handler: class.Get,
	}

	server.Services[createKey] = service.Route{
		Service: key,
		Action:  "create",
		Method:  "post",
		Path:    path,
		Handler: class.Create,
	}

	server.Services[updateKey] = service.Route{
		Service: key,
		Action:  "update",
		Method:  "put",
		Path:    path,
		Handler: class.Update,
	}

	server.Services[patchKey] = service.Route{
		Service: key,
		Action:  "patch",
		Method:  "patch",
		Path:    path,
		Handler: class.Patch,
	}

	server.Services[deleteKey] = service.Route{
		Service: key,
		Action:  "remove",
		Method:  "delete",
		Path:    path,
		Handler: class.Patch,
	}

	return server
}

// Start :
func (server *Oscrud) Start() {
	for _, trs := range server.Transports {

		for key, service := range server.Services {
			srv := strings.Split(key, ".")
			trs.RegisterService(srv[0], service)
		}

		for key, endpoint := range server.Endpoints {
			trs.RegisterEndpoint(key, endpoint)
		}

		go func(t transport.Transport) {
			err := t.Start()
			if err != nil {
				panic(err)
			}
		}(trs)
	}
	select {}
}
