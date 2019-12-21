package transport

import "oscrud/action"

// Tag Definitions
var (
	QueryTag = "query"
)

// Transport :
type Transport interface {
	Register(method string, path string, handler action.Handler)
	Start(handler action.RequestHandler) error
}
