# Oscrud

Oscrud is a golang resftul api wrapper framework. The purpose of the framework is make everything independent like transport, authentication, middleware and parser. So we can change the component to what we want anytime without changing code. This framework is inspired from [FeathersJS](https://feathersjs.com/). Currently framework still under development, any suggestion or PR is welcome.

# Planning

Current is what i planned to achive in version one of the framework.

### Transport

* Go-Http
* Echo
* SocketIO

### Binder

* Standardize `Bind(i interface{})` with reflect.Tag `query`, `param`, `body`.
* Query, Body, Header wiil be `map[string]interface{}` by default
* Param will be `map[string]string` by default

### Service

* CRUD Endpoints
* Standardize ORM Support

### Endpoints

* Single Action

# Context

* Get & Set ( map[string]interface{} )

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

	ec "oscrud/transport/echo"

	"github.com/labstack/echo/v4"
)

// // TestService :
// type TestService struct {
// }

// // NewService :
// func NewService() TestService {
// 	return TestService{}
// }

// // Find :
// func (t TestService) Find(service service.Context) error {
// 	log.Println("You're accessing TestService.Find")
// 	return nil
// }

// // Get :
// func (t TestService) Get(service service.Context) error {
// 	log.Println("You're accessing TestService.Get")
// 	return nil
// }

// // Create :
// func (t TestService) Create(service service.Context) error {
// 	log.Println("You're accessing TestService.Create")
// 	return nil
// }

// // Update :
// func (t TestService) Update(service service.Context) error {
// 	log.Println("You're accessing TestService.Update")
// 	return nil
// }

// // Patch :
// func (t TestService) Patch(service service.Context) error {
// 	log.Println("You're accessing TestService.Patch")
// 	return nil
// }

// // Remove :
// func (t TestService) Remove(service service.Context) error {
// 	log.Println("You're accessing TestService.Remove")
// 	return nil
// }

// Before :
func Before(ctx oscrud.Context) oscrud.Context {

	log.Println("I'm Before")

	return ctx
}

// Test2 :
func Test2(ctx oscrud.Context) oscrud.Context {
	var i struct {
		Test0 int    `param:"id"`
		Test2 uint64 `query:"test"`
		Test3 int32  `body:"test"`
	}

	err := ctx.Bind(&i)
	log.Println(i, err)
	log.Println("You're accessing Endpoint.")
	return ctx.JSON(200, "Value")
}

// After :
func After(ctx oscrud.Context) oscrud.Context {

	log.Println("I'm After")

	return ctx.End()
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterTransport(
		ec.NewEcho(echo.New()).UsePort(5001),
	)
	server.RegisterEndpoint("GET", "/test2/:id/test", Before, Test2, After)
	server.Start()
}



[LOG]
[00:57:54][OSCRUD] :    ____    __
[00:57:54][OSCRUD] :   / __/___/ /  ___
[00:57:54][OSCRUD] :  / _// __/ _ \/ _ \
[00:57:54][OSCRUD] : /___/\__/_//_/\___/ v4.1.11
[00:57:54][OSCRUD] : High performance, minimalist Go web framework
[00:57:54][OSCRUD] : https://echo.labstack.com
[00:57:54][OSCRUD] : ____________________________________O/_______
[00:57:54][OSCRUD] :                                     O\
[00:57:54][OSCRUD] : ⇨ http server started on [::]:5001
[00:57:57][OSCRUD] : 2019/12/31 00:57:57 I'm Before
[00:57:57][OSCRUD] : 2019/12/31 00:57:57 {12 123 0} <nil>
[00:57:57][OSCRUD] : 2019/12/31 00:57:57 You're accessing Endpoint.
[00:57:57][OSCRUD] : 2019/12/31 00:57:57 I'm After
```
