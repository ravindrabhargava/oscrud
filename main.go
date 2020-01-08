package oscrud

import (
	"oscrud/util"
	"reflect"
	"strings"

	"github.com/gbrlsnchs/radix"
)

// Oscrud :
type Oscrud struct {
	MiddlewareOptions
	EventOptions

	transports []Transport
	router     *radix.Tree
}

// NewOscrud :
func NewOscrud() *Oscrud {
	tree := radix.New(radix.Tsafe)
	tree.SetBoundaries(':', '/')
	return &Oscrud{
		transports: make([]Transport, 0),
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

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...Transport) *Oscrud {
	for _, transport := range transports {
		server.transports = append(server.transports, transport)
	}
	return server
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud {
	radix := util.RadixPath(method, endpoint)
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
		transport.Register(
			method, endpoint,
			func(req *Request) TransportResponse {
				ctx := server.lookupHandler(route, req)
				return ctx.transportResponse()
			},
		)
	}
	return server
}

// Start :
func (server *Oscrud) Start() {
	for _, trs := range server.transports {
		go func(t Transport) {
			err := t.Start(
				func(req *Request) TransportResponse {
					ctx := server.lookupHandler(nil, req)
					return ctx.transportResponse()
				},
			)
			if err != nil {
				panic(err)
			}
		}(trs)
	}
	select {}
}

// lookupHandler :
func (server *Oscrud) lookupHandler(route *Route, req *Request) Context {
	ctx := Context{
		method:    req.method,
		transport: req.transport,
		path:      req.path,
		param:     req.param,
		header:    req.header,
		query:     req.query,
		body:      req.body,
		sent:      false,
		result:    nil,
		exception: nil,
	}

	if route == nil {
		node, params := server.router.Get(util.RadixPath(req.method, req.path))
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
