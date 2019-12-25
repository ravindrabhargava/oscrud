package oscrud

import (
	"fmt"
	"oscrud/action"
	"oscrud/transport"
	"strings"
)

// Oscrud :
type Oscrud struct {
	Transports []transport.Transport
	Services   map[string]action.ServiceRoute
	Endpoints  map[string]action.EndpointRoute
}

// NewOscrud :
func NewOscrud() *Oscrud {
	return &Oscrud{
		Transports: make([]transport.Transport, 0),
		Services:   make(map[string]action.ServiceRoute),
		Endpoints:  make(map[string]action.EndpointRoute),
	}
}

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...transport.Transport) *Oscrud {
	for _, trs := range transports {
		server.Transports = append(server.Transports, trs)
	}
	return server
}

// CallService :
func (server *Oscrud) CallService(ctx ServiceContext) (*ServiceResult, error) {
	routeKey := ctx.Service + "." + ctx.Action
	service, ok := server.Services[routeKey]
	if !ok {
		return nil, fmt.Errorf("Service '%s.%s' not found, maybe you call before service registration?", ctx.Service, ctx.Action)
	}
	err := service.Handler(ctx)
	if err != nil {
		return nil, err
	}
	return ctx.Result, nil
}

// CallEndpoint :
func (server *Oscrud) CallEndpoint(ctx EndpointContext) (*EndpointResult, error) {
	routeKey := ctx.Endpoint
	route, ok := server.Endpoints[routeKey]
	if !ok {
		return nil, fmt.Errorf("Endpoint '%s' not found, maybe you call before endpoint registration?", routeKey)
	}

	ctx.Method = route.Method
	ctx.Path = route.Path
	err := route.Handler(ctx)
	if err != nil {
		return nil, err
	}
	return ctx.Result, nil
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(key, method, route string, handler action.EndpointHandler) *Oscrud {
	server.Endpoints[key] = action.EndpointRoute{
		Method:  method,
		Path:    strings.TrimPrefix(route, "/"),
		Handler: handler,
	}
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(key, basepath string, service action.Service) *Oscrud {
	path := strings.TrimPrefix(basepath, "/")
	findKey := key + ".find"
	createKey := key + ".create"
	getKey := key + ".get"
	updateKey := key + ".update"
	patchKey := key + ".patch"
	deleteKey := key + ".remove"

	server.Services[findKey] = action.ServiceRoute{
		Action:  "find",
		Method:  "get",
		Path:    path,
		Handler: service.Find,
	}

	server.Services[getKey] = action.ServiceRoute{
		Action:  "get",
		Method:  "get",
		Path:    path,
		Handler: service.Get,
	}

	server.Services[createKey] = action.ServiceRoute{
		Action:  "create",
		Method:  "post",
		Path:    path,
		Handler: service.Create,
	}

	server.Services[updateKey] = action.ServiceRoute{
		Action:  "update",
		Method:  "put",
		Path:    path,
		Handler: service.Update,
	}

	server.Services[patchKey] = action.ServiceRoute{
		Action:  "patch",
		Method:  "patch",
		Path:    path,
		Handler: service.Patch,
	}

	server.Services[deleteKey] = action.ServiceRoute{
		Action:  "remove",
		Method:  "delete",
		Path:    path,
		Handler: service.Patch,
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
