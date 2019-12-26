package transport

import (
	"oscrud/endpoint"
	"oscrud/service"
)

// Transport :
type Transport interface {
	RegisterEndpoint(endpoint string, route endpoint.Route)
	RegisterService(service string, route service.Route)
	Start() error
}
