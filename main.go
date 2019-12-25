package oscrud

import (
	"fmt"
	"oscrud/action"
	"oscrud/transport"
	"oscrud/util"
	"strings"
)

// Oscrud :
type Oscrud struct {
	Transports []transport.Transport
	Routes     []string
	Services   map[string]action.ServiceHandler
	Endpoints  map[string]action.EndpointHandler
}

// NewOscrud :
func NewOscrud() *Oscrud {
	return &Oscrud{
		Transports: make([]transport.Transport, 0),
		Routes:     make([]string, 0),
		Services:   make(map[string]action.ServiceHandler),
		Endpoints:  make(map[string]action.EndpointHandler),
	}
}

// CallService :
func (server *Oscrud) CallService(ctx ServiceContext) (*ServiceResult, error) {
	method := util.GetMethodByAction(ctx.Action)
	basePath := strings.TrimPrefix(ctx.Path, "/")
	serviceKey := strings.ToLower("service." + util.TransformPath(basePath, ctx.Action) + "." + method + "." + ctx.Action)
	serviceFn, ok := server.Services[serviceKey]
	if !ok {
		return nil, fmt.Errorf("Service '%s.%s' not found, maybe you call before service registration?", basePath, ctx.Action)
	}
	err := serviceFn(ctx)
	if err != nil {
		return nil, err
	}
	return ctx.Result, nil
}

// CallEndpoint :
func (server *Oscrud) CallEndpoint(ctx EndpointContext) (*EndpointResult, error) {
	routeKey := strings.ToLower("endpoint." + strings.TrimPrefix(ctx.Path, "/") + "." + ctx.Method)
	routeFn, ok := server.Endpoints[routeKey]
	if !ok {
		return nil, fmt.Errorf("Endpoint '%s %s' not found, maybe you call before endpoint registration?", strings.ToUpper(ctx.Method), ctx.Path)
	}
	err := routeFn(ctx)
	if err != nil {
		return nil, err
	}
	return ctx.Result, nil
}

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...transport.Transport) *Oscrud {
	for _, trs := range transports {
		server.Transports = append(server.Transports, trs)
	}
	return server
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(method string, basePath string, endpoint action.EndpointHandler) *Oscrud {
	path := strings.TrimPrefix(basePath, "/")
	routeKey := strings.ToLower("endpoint." + path + "." + method)
	server.Routes = append(server.Routes, routeKey)
	server.Endpoints[routeKey] = endpoint
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(basePath string, service action.Service) *Oscrud {
	path := strings.TrimPrefix(basePath, "/")

	findKey := "service." + path + ".get.find"
	createKey := "service." + path + ".post.create"
	getKey := "service." + path + "/:id.get.get"
	updateKey := "service." + path + "/:id.put.update"
	patchKey := "service." + path + "/:id.patch.patch"
	deleteKey := "service." + path + "/:id.delete.remove"

	server.Routes = append(server.Routes, findKey)
	server.Services[findKey] = service.Find

	server.Routes = append(server.Routes, getKey)
	server.Services[getKey] = service.Get

	server.Routes = append(server.Routes, createKey)
	server.Services[createKey] = service.Create

	server.Routes = append(server.Routes, updateKey)
	server.Services[updateKey] = service.Update

	server.Routes = append(server.Routes, patchKey)
	server.Services[patchKey] = service.Get

	server.Routes = append(server.Routes, deleteKey)
	server.Services[deleteKey] = service.Get

	return server
}

// Start :
func (server *Oscrud) Start() {
	for _, trs := range server.Transports {
		for _, route := range server.Routes {
			setting := strings.Split(route, ".")

			path := setting[1]
			method := setting[2]
			if setting[0] == "service" {
				action := setting[3]
				trs.RegisterService(action, method, path, server.Services[route])
			} else if setting[0] == "endpoint" {
				trs.RegisterEndpoint(method, path, server.Endpoints[route])
			}
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
