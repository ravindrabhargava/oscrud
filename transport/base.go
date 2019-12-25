package transport

import "oscrud/action"

// Transport :
type Transport interface {
	RegisterEndpoint(endpoint string, route action.EndpointRoute)
	RegisterService(service string, route action.ServiceRoute)
	Start() error
}
