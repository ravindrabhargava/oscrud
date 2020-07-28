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

	state      map[string]interface{}
	transports map[TransportID]Transport
	routes     []Route
	logger     []Logger
	binder     *binder.Binder
}

// NewOscrud :
func NewOscrud() *Oscrud {
	return &Oscrud{
		transports: make(map[TransportID]Transport),
		routes:     make([]Route, 0),
		logger:     make([]Logger, 0),
		binder:     binder.NewBinder(),
	}
}

// GetTransport :
func (server *Oscrud) GetTransport(key TransportID) (Transport, bool) {
	t, ok := server.transports[key]
	return t, ok
}

// SetState :
func (server *Oscrud) SetState(key string, value interface{}) {
	server.state[key] = value
}

// GetState :
func (server *Oscrud) GetState(key string) interface{} {
	return server.state[key]
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
		server.transports[transport.Name()] = transport
	}
	return server
}

// RegisterLogger :
func (server *Oscrud) RegisterLogger(loggers ...Logger) *Oscrud {
	for _, logger := range loggers {
		server.logger = append(server.logger, logger)
	}
	return server
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud {
	if len(server.transports) == 0 {
		panic("oscrud: register endpoint should be called after registered transports")
	}

	route := &Route{
		Method:  strings.ToLower(method),
		Path:    endpoint,
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

	if route.TransportOptions.DisableRegister == nil {
		route.TransportOptions.DisableRegister = make(map[TransportID]bool)
	}

	server.routes = append(server.routes, *route)
	for _, transport := range server.transports {
		if isDisabled, ok := route.TransportOptions.DisableRegister[transport.Name()]; ok && isDisabled {
			continue
		}
		transport.Register(method, endpoint, server.transportHandler(route))
	}
	return server
}

// RegisterService :
func (server *Oscrud) RegisterService(basePath string, service Service, serviceOptions *ServiceOptions, opts ...Options) *Oscrud {

	if serviceOptions == nil {
		serviceOptions = new(ServiceOptions)
	}

	if !serviceOptions.DisableCreate {
		server.RegisterEndpoint("post", basePath, service.Create, opts)
	}

	if !serviceOptions.DisableFind {
		server.RegisterEndpoint("get", basePath, service.Find, opts)
	}

	if !serviceOptions.DisableGet {
		server.RegisterEndpoint("get", basePath+"/:$id", service.Get, opts)
	}

	if !serviceOptions.DisablePatch {
		server.RegisterEndpoint("patch", basePath+"/:$id", service.Patch, opts)
	}

	if !serviceOptions.DisableUpdate {
		server.RegisterEndpoint("put", basePath+"/:$id", service.Update, opts)
	}

	if !serviceOptions.DisableDelete {
		server.RegisterEndpoint("delete", basePath+"/:$id", service.Delete, opts)
	}

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
	return func(req *Request) Response {
		ctx := server.lookupHandler(route, req)
		return ctx.response
	}
}

func (server *Oscrud) lookupHandler(route *Route, req *Request) Context {

	req.path = route.Path
	req.method = route.Method

	if req.requestID == "" {
		req.requestID = uuid.New().String()
	}

	if req.context == nil {
		req.context = context.Background()
	}

	ctx := Context{
		oscrud:   *server,
		request:  *req,
		response: Response{},
		sent:     false,
	}

	for _, logger := range server.logger {
		go logger.StartRequest(ctx)
	}

	gr := make(chan Context, 1)
	go server.invokeHandler(ctx, req, route, gr)

	duration := 30 * time.Second
	if server.Duration != 0 {
		duration = server.Duration
	}
	if route.Duration != 0 {
		duration = route.Duration
	}

	select {
	case ctx = <-gr:
	case <-time.After(duration):
		if route.OnTimeout != nil {
			gr <- route.OnTimeout(ctx)
			break
		}
		if server.OnTimeout != nil {
			gr <- server.OnTimeout(ctx)
			break
		}
		gr <- ctx.Error(408, ErrRequestTimeout)
		break
	}

	for _, logger := range server.logger {
		go logger.EndRequest(ctx)
	}

	if ctx.response.status == 0 {
		return ctx.NoContent()
	}
	return ctx
}

func (server *Oscrud) invokeHandler(ctx Context, req *Request, route *Route, gr chan Context) {
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
			break
		}
		ctx = handler(ctx)
	}

	if route.OnComplete != nil {
		go route.OnComplete(ctx)
	}

	if server.OnComplete != nil {
		go server.OnComplete(ctx)
	}

	gr <- ctx
	return
}
