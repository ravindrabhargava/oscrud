package oscrud

import (
	"oscrud/util"
	"strings"

	"github.com/gbrlsnchs/radix"
)

// Oscrud :
type Oscrud struct {
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

// RegisterTransport :
func (server *Oscrud) RegisterTransport(transports ...Transport) *Oscrud {
	for _, transport := range transports {
		server.transports = append(server.transports, transport)
	}
	return server
}

// RegisterEndpoint :
func (server *Oscrud) RegisterEndpoint(method, endpoint string, handler ...Handler) *Oscrud {
	radix := util.RadixPath(method, endpoint)
	route := &Route{
		Method:  strings.ToLower(method),
		Route:   endpoint,
		Handler: handler,
	}

	server.router.Add(radix, route)
	for _, transport := range server.transports {
		transport.Register(
			method, endpoint,
			func(req *Request) (*ResultResponse, *ErrorResponse) {
				ctx := server.lookupHandler(route, req)
				return ctx.result, ctx.exception
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
				func(req *Request) (*ResultResponse, *ErrorResponse) {
					ctx := server.lookupHandler(nil, req)
					return ctx.result, ctx.exception
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

	for _, handler := range route.Handler {
		ctx = handler(ctx)
		if ctx.sent {
			return ctx.End()
		}
	}

	return ctx.missingEnd().End()
}
