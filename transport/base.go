package transport

import "oscrud/action"

// Transport :
type Transport interface {
	RegisterEndpoint(method, path string, handler action.EndpointHandler)
	RegisterService(service, method, path string, handler action.ServiceHandler)
	Start() error
}
