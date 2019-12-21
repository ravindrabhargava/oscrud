package transport

import "oscrud/action"

// Transport :
type Transport interface {
	RegisterEndpoint(method string, path string, handler action.EndpointHandler)
	RegisterService(service string, method, path string, handler action.ServiceHandler)
	Start() error
}
