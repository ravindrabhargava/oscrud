package oscrud

import (
	"fmt"
	"oscrud/service"
)

// Service :
func (server *Oscrud) Service(service string) Service {
	return Service{server, service}
}

func serviceCall(s Service, ctx *service.Request, action string) (*service.Response, error) {
	routeKey := s.service + "." + action
	service, ok := s.server.Services[routeKey]
	if !ok {
		return nil, fmt.Errorf("Service '%s.%s' not found, maybe you call before service registration?", s.service, action)
	}
	return service.Call(ctx)
}

// Get :
func (s Service) Get(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "get")
}

// Find :
func (s Service) Find(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "find")
}

// Create :
func (s Service) Create(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "create")
}

// Update :
func (s Service) Update(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "update")
}

// Patch :
func (s Service) Patch(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "patch")
}

// Remove :
func (s Service) Remove(ctx *service.Request) (*service.Response, error) {
	return serviceCall(s, ctx, "remove")
}
