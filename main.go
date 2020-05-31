package oscrud

import (
	"context"
	"reflect"
	"strings"
	"time"

	uuid "github.com/google/uuid"
	"github.com/oscrud/binder"
)

// Oscrud :
type Oscrud struct {
	MiddlewareOptions
	EventOptions
	TimeoutOptions

	transports []Transport
	routes     []Route
	logger     []Logger
	binder     *binder.Binder
}

// NewOscrud :
func NewOscrud() *Oscrud {
	return &Oscrud{
		transports: make([]Transport, 0),
		routes:     make([]Route, 0),
		logger:     make([]Logger, 0),
		binder:     binder.NewBinder(),
	}
}

// Log :
func (server *Oscrud) Log(operation string, content string) {
	for _, logger := range server.logger {
		logger.Log(operation, content)
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
func (server *Oscrud) RegisterBinder(ftype interface{}, ttype interface{}, bindFn binder.Bind) *Oscrud {
	server.binder.Register(ftype, ttype, bindFn)
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

	server.routes = append(server.routes, *route)
	for _, transport := range server.transports {
		transport.Register(method, endpoint, server.transportHandler(route))
	}
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(basePath string, service Service, opts ...Options) *Oscrud {
	server.RegisterEndpoint("get", basePath, service.Find, opts)
	server.RegisterEndpoint("post", basePath, service.Create, opts)
	server.RegisterEndpoint("get", basePath+"/:$id", service.Get, opts)
	server.RegisterEndpoint("put", basePath+"/:$id", service.Update, opts)
	server.RegisterEndpoint("patch", basePath+"/:$id", service.Patch, opts)
	server.RegisterEndpoint("delete", basePath+"/:$id", service.Delete, opts)
	return server

}

// Start :
func (server *Oscrud) Start() {
	for _, trs := range server.transports {
		go func(t Transport) {
			err := t.Start()
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

	if req.requestID == "" {
		req.requestID = uuid.New().String()
	}

	if req.context == nil {
		req.context = context.Background()
	}

	ctx := Context{
		oscrud:    *server,
		route:     *route,
		request:   *req,
		sent:      false,
		result:    nil,
		exception: nil,
	}

	gr := make(chan Context, 1)
	go server.invokeHandler(ctx, gr)

	duration := 30 * time.Second
	if server.Duration != 0 {
		duration = server.Duration
	}
	if route.Duration != 0 {
		duration = route.Duration
	}

	select {
	case <-time.After(duration):
		gr <- ctx

		if route.OnTimeout != nil {
			return route.OnTimeout(ctx)
		}
		if server.OnTimeout != nil {
			return server.OnTimeout(ctx)
		}
		return ctx.Error(408, ErrRequestTimeout).End()
	case ctx = <-gr:
		return ctx
	}
}

func (server *Oscrud) invokeHandler(ctx Context, gr chan Context) {

	for _, logger := range server.logger {
		logger.StartRequest(ctx)
	}

	req := ctx.request
	route := ctx.route

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

		if len(gr) > 0 {
			return
		}

		ctx = handler(ctx)
		if ctx.sent {
			// EventOptions :
			if route.OnComplete != nil {
				go route.OnComplete(ctx)
			}

			if server.OnComplete != nil {
				go server.OnComplete(ctx)
			}

			for _, logger := range server.logger {
				logger.EndRequest(ctx)
			}
			gr <- ctx
			return
		}
	}
	gr <- ctx.missingEnd()
	return
}
