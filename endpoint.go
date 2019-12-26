package oscrud

import (
	"fmt"
	"oscrud/endpoint"
)

// Endpoint :
func (server *Oscrud) Endpoint(endpoint string, ctx *endpoint.Request) (*endpoint.Response, error) {
	routeKey := endpoint
	route, ok := server.Endpoints[routeKey]
	if !ok {
		return nil, fmt.Errorf("Endpoint '%s' not found, maybe you call before endpoint registration?", endpoint)
	}
	return route.Call(ctx)
}
