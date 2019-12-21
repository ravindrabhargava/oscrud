package main

import (
	"log"
	"oscrud"
	"oscrud/action"
	"oscrud/parser/basic"
	ec "oscrud/transport/echo"
	"oscrud/transport/http"

	"github.com/labstack/echo"
)

// TestService :
type TestService struct {
}

// NewService :
func NewService() TestService {
	return TestService{}
}

// Find :
func (t TestService) Find(service action.ServiceContext, ctx action.EndpointContext) error {
	log.Println("You're accessing TestService.Find")
	return nil
}

// Get :
func (t TestService) Get(service action.ServiceContext, ctx action.EndpointContext) error {
	log.Println("You're accessing TestService.Get")
	return nil
}

// Create :
func (t TestService) Create(service action.ServiceContext, ctx action.EndpointContext) error {
	log.Println("You're accessing TestService.Create")
	return nil
}

// Update :
func (t TestService) Update(service action.ServiceContext, ctx action.EndpointContext) error {
	log.Println("You're accessing TestService.Update")
	return nil
}

// Patch :
func (t TestService) Patch(service action.ServiceContext, ctx action.EndpointContext) error {
	log.Println("You're accessing TestService.Patch")
	return nil
}

// Remove :
func (t TestService) Remove(service action.ServiceContext, ctx action.EndpointContext) error {
	log.Println("You're accessing TestService.Remove")
	return nil
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterService("test", NewService())
	e := echo.New()
	server.RegisterTransport(
		ec.NewEcho(e).UsePort(5001).UseParser(basic.NewParser()),
		http.NewHTTP().UsePort(5000).UseParser(basic.NewParser()),
	)
	server.Start()
}
