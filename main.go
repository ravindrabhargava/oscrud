package oscrud

import (
	"reflect"
	"strings"

	"github.com/gbrlsnchs/radix"
)

// Oscrud :
type Oscrud struct {
	MiddlewareOptions
	EventOptions

	transports []Transport
	binder     Binder
	router     *radix.Tree
}

// NewOscrud :
func NewOscrud() *Oscrud {
	tree := radix.New(radix.Tsafe)
	tree.SetBoundaries(':', '/')
	return &Oscrud{
		transports: make([]Transport, 0),
		binder:     NewBinder(),
		router:     tree,
	}
}

// UseOptions :
func (server *Oscrud) UseOptions(opts ...Options) *Oscrud {
	serverElem := reflect.ValueOf(server).Elem()
	for _, iopt := range opts {
		name := reflect.TypeOf(iopt).Name()
		opt := reflect.ValueOf(iopt)
		field := serverElem.FieldByName(name)
		if field.CanSet() {
			field.Set(opt)
		}
	}
	return server
}

// RegisterBinder :
func (server *Oscrud) RegisterBinder(rtype interface{}, bindFn Bind) *Oscrud {
	server.binder.Register(rtype, bindFn)
	return server
}

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...Transport) *Oscrud {
	for _, transport := range transports {
		server.transports = append(server.transports, transport)
	}
	return server
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud {
	radix := radixPath(method, endpoint)
	route := &Route{
		Method:  strings.ToLower(method),
		Route:   endpoint,
		Handler: handler,
	}

	routeElem := reflect.ValueOf(route).Elem()
	for _, iopt := range opts {
		name := reflect.TypeOf(iopt).Name()
		opt := reflect.ValueOf(iopt)
		field := routeElem.FieldByName(name)
		if field.CanSet() {
			field.Set(opt)
		}
	}

	server.router.Add(radix, route)
	for _, transport := range server.transports {
		transport.Register(method, endpoint, server.transportHandler(route))
	}
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(basePath string, service Service, opts ...Options) *Oscrud {
	server.RegisterEndpoint("get", basePath, service.Find, opts)
	server.RegisterEndpoint("post", basePath, service.Create, opts)
	return server

}

// Start :
func (server *Oscrud) Start() {
	for _, trs := range server.transports {
		go func(t Transport) {
			err := t.Start(server.transportHandler(nil))
			if err != nil {
				panic(err)
			}
		}(trs)
	}
	select {}
}

func (server *Oscrud) transportHandler(route *Route) TransportHandler {
	return func(req *Request) TransportResponse {
		ctx := server.lookupHandler(route, req)
		return ctx.transportResponse()
	}
}

func (server *Oscrud) lookupHandler(route *Route, req *Request) Context {
	ctx := Context{
		method:    req.method,
		transport: req.transport,
		path:      req.path,
		param:     req.param,
		header:    req.header,
		query:     req.query,
		body:      req.body,
		oscrud:    *server,
		sent:      false,
		result:    nil,
		exception: nil,
	}

	if route == nil {
		node, params := server.router.Get(radixPath(req.method, req.path))
		if node == nil {
			return ctx.NotFound().End()
		}

		route = node.Value.(*Route)
		ctx.param = params
	}

	// MiddlewareOptions :
	handlers := make([]Handler, 0)
	if req.skip != skipMiddleware && req.skip != skipBefore && route.Before != nil {
		handlers = append(handlers, server.Before...)
		handlers = append(handlers, route.Before...)
	}
	handlers = append(handlers, route.Handler)
	if req.skip != skipMiddleware && req.skip != skipAfter && route.After != nil {
		handlers = append(handlers, route.After...)
		handlers = append(handlers, server.After...)
	}

	for _, handler := range handlers {
		ctx = handler(ctx)
		if ctx.sent {
			// EventOptions :
			if route.OnComplete != nil {
				go route.OnComplete(ctx)
			}

			if server.OnComplete != nil {
				go server.OnComplete(ctx)
			}

			return ctx.End()
		}
	}

	return ctx.missingEnd().End()
}

func radixPath(method, path string) string {
	return strings.ToLower(method) + fixPath(path)
}

func fixPath(path string) string {
	if !strings.HasPrefix(path, "/") {
		path = "/" + path
	}

	if strings.HasSuffix(path, "/") {
		path = strings.TrimSuffix(path, "/")
	}

	return path
}
