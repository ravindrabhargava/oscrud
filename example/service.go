package main

import (
	"log"
	"oscrud"
	"oscrud/action"
	ec "oscrud/transport/echo"

	"github.com/labstack/echo/v4"
)

// TestService :
type TestService struct {
}

// NewService :
func NewService() TestService {
	return TestService{}
}

// Find :
func (t TestService) Find(service action.ServiceContext) error {
	log.Println("You're accessing TestService.Find")
	return nil
}

// Get :
func (t TestService) Get(service action.ServiceContext) error {
	log.Println("You're accessing TestService.Get")
	return nil
}

// Create :
func (t TestService) Create(service action.ServiceContext) error {
	log.Println("You're accessing TestService.Create")
	return nil
}

// Update :
func (t TestService) Update(service action.ServiceContext) error {
	log.Println("You're accessing TestService.Update")
	return nil
}

// Patch :
func (t TestService) Patch(service action.ServiceContext) error {
	log.Println("You're accessing TestService.Patch")
	return nil
}

// Remove :
func (t TestService) Remove(service action.ServiceContext) error {
	log.Println("You're accessing TestService.Remove")
	return nil
}

// Test2 :
func Test2(ctx action.EndpointContext) error {
	log.Println("You're accessing Endpoint.")
	return nil
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterService("test", NewService())
	server.RegisterEndpoint("GET", "/test2", Test2)

	server.RegisterTransport(
		ec.NewEcho(echo.New()).UsePort(5001),
	)

	server.CallEndpoint(
		oscrud.EndpointContext{
			Method: "GET",
			Path:   "/test2",
		},
	)

	server.Start()
}
