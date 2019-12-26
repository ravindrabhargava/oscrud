package oscrud

import (
	"fmt"
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

// Endpoint :
func (server *Oscrud) Endpoint(endpoint string, ctx *endpoint.Request) (*endpoint.Response, error) {
	routeKey := endpoint
	route, ok := server.Endpoints[routeKey]
	if !ok {
		return nil, fmt.Errorf("Endpoint '%s' not found, maybe you call before endpoint registration?", endpoint)
	}
	return route.Call(ctx)
}

// Service :
func (server *Oscrud) Service(service string) Service {
	return Service{server, service}
}

func serviceCall(s Service, ctx *service.Request, action string) (*service.Response, error) {
	routeKey := s.service + "." + action
	service, ok := s.server.Services[routeKey]
	if !ok {
		return nil, fmt.Errorf("Service '%s.%s' not found, maybe you call before service registration?", s.service, action)
	}
	return service.Call(ctx)
}

// Get :
func (s Service) Get(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "get")
}

// Find :
func (s Service) Find(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "find")
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
