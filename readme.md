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

func main() {
	server := oscrud.NewOscrud()

	// Register transport
	server.RegisterTransport(
		ec.NewEcho(echo.New()).UsePort(5001),
		sc.NewSocket(nil).UsePort(3000),
	)

	// Register data binding for specific struct / slice / array.
	server.RegisterBinder(
		AnyStruct{},
		func(raw interface{}) (interface{}, error) {
			str, ok := raw.(string)
			if ok {
				return AnyStruct{str}, nil
			}
			return nil, errors.New("received data isn't a string")
		},
	)

	// Endpoint options definition ( usually be middleware )
	event := oscrud.EventOptions{
		OnComplete: func(ctx oscrud.Context) {
			log.Println("This running from go-routine as event-drive OnComplete().")
		},
	}
	middleware := oscrud.MiddlewareOptions{
		Before: []oscrud.Handler{Before},
		After:  []oscrud.Handler{After},
	}

	// Register Endpoint
	server.RegisterEndpoint("GET", "/test2/:id/test", Test2, event, middleware)

	// Internal Call
	req := oscrud.NewRequest().Query("any_struct", "223")
	res := server.Endpoint("GET", "/test2/1/test", req)
	log.Println(res.Result(), res.Error())

	res = server.Endpoint("GET", "/test2/0/test", oscrud.NewRequest())
	log.Println(res.Result(), res.Error())

	// Sqlike database conn initialize
	client := sql.MustConnect("mysql",
		options.Connect().
			SetHost("localhost").
			SetPort("3306").
			SetUsername("root").
			SetPassword("test"),
	)
	client.SetPrimaryKey("Key")

	// Service Definition
	service := sqlike.NewService(client).Database("test")
	server.RegisterService("test", service.ToService("user", new(User)))

	// Everything done? Start the server.
	server.Start()
}


[LOG]
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 I'm Before
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 Binder AnyStruct :  {223}
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 I'm After
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 You're accessing endpoint. <nil>
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 I'm Before
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 I'm After
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 <nil> ID should bigger than 0
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 Connect to : root:test@tcp(localhost:3306)/?parseTime=true&charset=utf8mb4&collation=utf8mb4_unicode_ci
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 This running from go-routine as event-drive OnComplete().
[00:22:30][OSCRUD] : 2020/01/31 00:22:30 This running from go-routine as event-drive OnComplete().
[00:22:30][OSCRUD] :    ____    __
[00:22:30][OSCRUD] :   / __/___/ /  ___
[00:22:30][OSCRUD] :  / _// __/ _ \/ _ \
[00:22:30][OSCRUD] : /___/\__/_//_/\___/ v4.1.11
[00:22:30][OSCRUD] : High performance, minimalist Go web framework
[00:22:30][OSCRUD] : https://echo.labstack.com
[00:22:30][OSCRUD] : ____________________________________O/_______
[00:22:30][OSCRUD] :                                     O\
[00:22:30][OSCRUD] : ⇨ http server started on [::]:5001
```
