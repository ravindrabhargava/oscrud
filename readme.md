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

```go
package main

import (
	"log"
	"oscrud"

	ec "oscrud/transport/echo"
	sc "oscrud/transport/socketio"

	"github.com/labstack/echo/v4"
)

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
	if err != nil {
		return ctx.Stack(500, err)
	}

	if i.Test0 == 0 {
		return ctx.Error(500, "ID should bigger than 0")
	}

	log.Println(i, err)
	log.Println("You're accessing Endpoint.")
	return ctx.String(200, "TestValue")
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
		sc.NewSocket(nil).UsePort(3000),
	)

	event := oscrud.EventOptions{
		OnComplete: func(ctx oscrud.Context) {
			log.Println("This running from go-routine as event-drive OnComplete().")
		},
	}
	middleware := oscrud.MiddlewareOptions{
		Before: []oscrud.Handler{Before},
		After:  []oscrud.Handler{After},
	}

	server.RegisterEndpoint("GET", "/test2/:id/test", Test2, event, middleware)

	res := server.Endpoint("GET", "/test2/1/test", oscrud.NewRequest())
	log.Println(res.Result(), res.Error())

	res = server.Endpoint("GET", "/test2/0/test", oscrud.NewRequest())
	log.Println(res.Result(), res.Error())

	server.Start()
}


[LOG]
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 I'm Before
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 {1 0 0} <nil>
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 You're accessing Endpoint.
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 I'm After
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 &{200 text/plain TestValue} <nil>
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 I'm After
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 <nil> &{500 <nil> ID should bigger than 0}
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 This running from go-routine as event-drive OnComplete().
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 This running from go-routine as event-drive OnComplete().
[11:36:53][OSCRUD] :    ____    __
[11:36:53][OSCRUD] :   / __/___/ /  ___
[11:36:53][OSCRUD] :  / _// __/ _ \/ _ \
[11:36:53][OSCRUD] : /___/\__/_//_/\___/ v4.1.11
[11:36:53][OSCRUD] : High performance, minimalist Go web framework
[11:36:53][OSCRUD] : https://echo.labstack.com
[11:36:53][OSCRUD] : ____________________________________O/_______
[11:36:53][OSCRUD] : 2020/01/08 11:36:53 I'm Before
[11:36:53][OSCRUD] :                                     O\
[11:36:53][OSCRUD] : ⇨ http server started on [::]:5001
```
