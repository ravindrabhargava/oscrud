# Oscrud

Oscrud is a golang resftul api wrapper framework. The purpose of the framework is make everything independent like transport, authentication, middleware and parser. So we can change the component to what we want anytime without changing code. This framework is inspired from [FeathersJS](https://feathersjs.com/). Currently framework still under development, any suggestion or PR is welcome.

# Planning

Current is what i planned to achive in version one of the framework.

### Transport

* Echo
* Micro
* SocketIO

### Binder

* Standardize `Bind(i interface{})` with reflect.Tag `query`, `param`, `body`.
* Query and Body wiil be `map[string]interface{}` by default
* Param will be `map[string]string` by default

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
	"oscrud/endpoint"
	"oscrud/service"
	ec "oscrud/transport/echo"
	socketio "oscrud/transport/socketio"

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
func (t TestService) Find(service service.Context) error {
	log.Println("You're accessing TestService.Find")
	return nil
}

// Get :
func (t TestService) Get(service service.Context) error {
	log.Println("You're accessing TestService.Get")
	return nil
}

// Create :
func (t TestService) Create(service service.Context) error {
	log.Println("You're accessing TestService.Create")
	return nil
}

// Update :
func (t TestService) Update(service service.Context) error {
	log.Println("You're accessing TestService.Update")
	return nil
}

// Patch :
func (t TestService) Patch(service service.Context) error {
	log.Println("You're accessing TestService.Patch")
	return nil
}

// Remove :
func (t TestService) Remove(service service.Context) error {
	log.Println("You're accessing TestService.Remove")
	return nil
}

// Test2 :
func Test2(ctx endpoint.Context) error {
	var i struct {
		Test0 int    `param:"id"`
		Test1 string `param:"test"`
		Test2 uint64 `query:"test"`
		Test3 int32  `body:"test"`
	}

	err := ctx.Bind(&i)
	log.Println(i, err)
	log.Println("You're accessing Endpoint.")
	ctx.JSON(200, i)
	return nil
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterService("test", "test", NewService())
	server.RegisterEndpoint("test2", "GET", "/test2/:id/:test", Test2)
	server.RegisterTransport(
		ec.NewEcho(echo.New()).UsePort(5001),
		socketio.NewSocket(nil).UsePort(5000),
	)

	req := service.NewRequest()
	server.Service("test").Find(req)

	param := map[string]string{
		"id":   "12",
		"test": "1",
	}
	body := map[string]interface{}{
		"test": 100,
	}
	query := map[string]interface{}{
		"test": 2000,
	}
	req2 := endpoint.NewRequest().SetParam(param).SetBody(body).SetQuery(query)
	server.Endpoint("test2", req2)

	server.Start()
}


[LOG]
[21:20:01][OSCRUD] : 2019/12/26 21:20:01 You're accessing TestService.Find
[21:20:01][OSCRUD] : 2019/12/26 21:20:01 {12 1 2000 100} <nil>
[21:20:01][OSCRUD] : 2019/12/26 21:20:01 You're accessing Endpoint.
[21:20:01][OSCRUD] :    ____    __
[21:20:01][OSCRUD] :   / __/___/ /  ___
[21:20:01][OSCRUD] :  / _// __/ _ \/ _ \
[21:20:01][OSCRUD] : /___/\__/_//_/\___/ v4.1.11
[21:20:01][OSCRUD] : High performance, minimalist Go web framework
[21:20:01][OSCRUD] : https://echo.labstack.com
[21:20:01][OSCRUD] : ____________________________________O/_______
[21:20:01][OSCRUD] :                                     O\
[21:20:01][OSCRUD] : ⇨ http server started on [::]:5001
```
