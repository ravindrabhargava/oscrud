<!-- omit in toc -->
# Table Of Contents

- [Introduction](#introduction)
- [Getting Started](#getting-started)
- [Oscrud Server](#oscrud-server)
  - [NewOscrud() *Oscrud](#newoscrud-oscrud)
  - [UseOptions(opts ...Options) *Oscrud](#useoptionsopts-options-oscrud)
  - [RegisterBinder(rtype interface{}, bindFb Bind) *Oscrud](#registerbinderrtype-interface-bindfb-bind-oscrud)
  - [RegisterTransport(transports ...Transport) *Oscrud](#registertransporttransports-transport-oscrud)
  - [RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud](#registerendpointmethod-endpoint-string-handler-handler-opts-options-oscrud)
  - [RegisterService(basePath string, service Service, opts ...Options) *Oscrud](#registerservicebasepath-string-service-service-opts-options-oscrud)
  - [Start()](#start)
- [Handler & Context](#handler--context)
- [References & Resources](#references--resources)
  - [Available Options](#available-options)
  - [Transport ( Official / Community )](#transport--official--community-)
  - [Service ( Official / Community )](#service--official--community-)

# Introduction

Oscrud is a golang resftul api wrapper framework. The purpose of the framework is make everything independent like transport, authentication, middleware and parser. So we can change the component to what we want anytime without changing code. This framework is inspired from [FeathersJS](https://feathersjs.com/). Currently framework still under development, any suggestion or PR is welcome.

# Getting Started

Currently package only tested in Go 1.13. To install the package 

```
$ go get -u github.com/Oskang09/oscrud
```

After complete installation, you can Go with your beloved framework and here an hello world example. You can choose your own transport from lists, currently only supported Echo. For future will implement [service discovery](https://github.com/hashicorp/mdns) by default.

```go
package main

import (
    "github.com/Oskang09/oscrud"
)

func main() {
	server := oscrud.NewOscrud()
	server.RegisterTransport(e.NewEcho(echo.New()).UsePort(3000))
	server.RegisterEndpoint("GET", "/test", func(ctx oscrud.Context) oscrud.Context {
		return ctx.String(200, "Hello World").End()
	})
	server.Start()
}

$ curl -v localhost:3000/test

> GET /test HTTP/1.1
> Host: localhost:3000
> User-Agent: curl/7.58.0
> Accept: */*
> 
< HTTP/1.1 200 OK
< Content-Type: text/plain; charset=UTF-8
< Date: Fri, 31 Jan 2020 16:04:49 GMT
< Content-Length: 11
< 
* Connection #0 to host localhost left intact
Hello World
```

# Oscrud Server

`oscrud.Server` is the core struct of the framework, inside server core you can have setup method to initialize components ( middleware, binder, and more ).

## NewOscrud() *Oscrud

For constructing a new instance with some default parameters. *Preferred* to use this instead construct yourself unless you know what you're doing.

```go
package main

server := oscrud.NewOscrud()
```

## UseOptions(opts ...Options) *Oscrud

For apply server-level options ( mean apply to all endpoints ).

```go
package main

server := oscrud.NewOscrud()
middleware := oscrud.MiddlewareOptions{
	Before: []oscrud.Handler{
        func(ctx oscrud.Context) oscrud.Context {
            log.Println("I'm Before Middleware")
            return ctx
        }
    },
}
event := oscrud.EventOptions{
    OnComplete: func(ctx oscrud.Context) {
        log.Println("This running from go-routine as event-drive OnComplete().")
    },
}
// You can just apply all in one line.
server.UseOptions(middleware, event)
```

## RegisterBinder(rtype interface{}, bindFb Bind) *Oscrud 

For register data binding method for specific struct / array / slice. Incoming data can be any type so suggested to be make a switch-case statement with default by handling other type that not supported ( usually just a string from 'header', 'query', 'body' or 'param' ).

```go
package main

// AnyStruct :
type AnyStruct struct {
    Data string
}

server := oscrud.NewOscrud()
server.RegisterBinder(AnyStruct{}, func(raw interface{}) (interface{}, error) {
    str, ok := raw.(string)
    if ok {
        return AnyStruct{str}, nil
    }
    return nil, errors.New("Trying to deserialize non-string to AnyStruct.")
})
```

## RegisterTransport(transports ...Transport) *Oscrud 

For register transport for the server, *must be called before any endpoint registration*. Every transport must be implemented based on interface.

```go
package main

func main() {
	server := oscrud.NewOscrud()
    server.RegisterTransport(e.NewEcho(echo.New()).UsePort(3000))
}
```

## RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud 

For registering endpoint with specified method, endpoint & handler, and also able to apply endpoint level options ( mean only work on the specifed endpoint ).

```go
package main

func main() {
	server := oscrud.NewOscrud()
    middleware := oscrud.MiddlewareOptions{
        Before: []oscrud.Handler{
            func(ctx oscrud.Context) oscrud.Context {
                log.Println("I'm Before Middleware")
                return ctx
            }
        },
    }
    event := oscrud.EventOptions{
        OnComplete: func(ctx oscrud.Context) {
            log.Println("This running from go-routine as event-drive OnComplete().")
        },
    }

    // options can be apply in one line also.
	server.RegisterEndpoint("GET", "/test", func(ctx oscrud.Context) oscrud.Context {
		return ctx.String(200, "Hello World").End()
	}, event, middleware)
}
```

## RegisterService(basePath string, service Service, opts ...Options) *Oscrud

For registering service on a specified path, basically service would includes 6 endpoints. It's have same registering strategy with endpoint.

* GET /basePath - Service.Find
* GET /basePath/:$id - Service.Get
* POST /basePath - Service.Create
* PUT /basePath/:$id - Service.Update
* PATCH /basePath/:$id - Service.Patch
* DELETE /basePath/:$id - Service.Delete

```go
package main

func main() {
	server := oscrud.NewOscrud()
    middleware := oscrud.MiddlewareOptions{
        Before: []oscrud.Handler{
            func(ctx oscrud.Context) oscrud.Context {
                log.Println("I'm Before Middleware")
                return ctx
            }
        },
    }

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

    // User is a query model struct based on oscrud.ServiceModel interface
	server.RegisterService("test", service.ToService("user", new(User)), middleware)
}
```

## Start()

For start the oscrud server, server starting will only start all the registered transport. No transported registered will panic with err. No any setup will be invoke at this step, all would be done when `register`, so internal call / access can be work before or even not calling `Start()`.

```go
func main() {
	server := oscrud.NewOscrud()
	server.Start()
}
```

# Handler & Context

`oscrud.Handler` is just a function `func(Context) Context`. `oscrud.Handler` act as main handler for whole framework, endpoint handler is using it and middleware options also using the same handler.



# References & Resources

Reference and Resource about documentation like list of usable options or third party url.

## Available Options

* oscrud.MiddlewareOptions
* oscrud.EventOptions
* oscrud.TimeoutOptions

## Transport ( Official / Community )

* oscrud-echo

## Service ( Official / Community )

* oscrud-sqlike