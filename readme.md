# Oscrud

Oscrud is a golang resftul api wrapper framework. The purpose of the framework is make everything independent like transport, authentication, middleware and parser. So we can change the component to what we want anytime without changing code. This framework is inspired from [FeathersJS](https://feathersjs.com/). Currently framework still under development, any suggestion or PR is welcome.

# Planning

Current is what i planned to achive in version one of the framework.

### Transport

* Echo
* Micro
* SocketIO

### Service

* CRUD Endpoints
* Standardize ORM Support

### Endpoints

* Single Action

### Authentication

* Local ( username & password )
* Key ( api key )
* OAuth
* JWT

### Middleware

* Before & After

### Parser

* ParseQuery
* ParseBody
* ParseValue

# Start Server

```
$ git clone https://github.com/Oskang09/oscrud.git
$ go get https://github.com/oxequa/realize // If you don't have 'realize'
$ realize start
```

# PR & Suggestion 

* Cases
* Example

# Example

You can view more [examples](https://github.com/Oskang09/oscrud/tree/master/example) at `example` folder.

```
package main

import (
	"log"
	"oscrud"
	"oscrud/action"
	"oscrud/parser/basic"
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
		ec.NewEcho(echo.New()).UsePort(5001).UseParser(basic.NewParser()),
	)
	server.Start()
}
```
