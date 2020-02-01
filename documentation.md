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
  - [Endpoint(method, path string, req *Request) TransportResponse](#endpointmethod-path-string-req-request-transportresponse)
  - [NewRequest(args ...string) *Request](#newrequestargs-string-request)
  - [Start()](#start)
- [Handler & Context](#handler--context)
  - [Example](#example)
  - [Context](#context)
  - [Request Methods](#request-methods)
  - [Response Methods](#response-methods)
  - [Internal Access / Call](#internal-access--call)
    - [Example](#example-1)
    - [Transport Implementation](#transport-implementation)
- [Binder](#binder)
  - [All Binding](#all-binding)
  - [Specific Binding](#specific-binding)
  - [Register New Binder](#register-new-binder)
- [Service](#service)
- [Transport](#transport)
- [References & Resources](#references--resources)
  - [Content Type List](#content-type-list)
  - [Errors List](#errors-list)
  - [Available Options](#available-options)
    - [MiddlewareOptions](#middlewareoptions)
    - [EventOptions](#eventoptions)
    - [TimeoutOptions](#timeoutoptions)
  - [Transport ( Official / Community )](#transport--official--community-)
  - [Service ( Official / Community )](#service--official--community-)

# Introduction

Oscrud is a golang resftul api wrapper framework. The purpose of the framework is make everything independent like transport, authentication, middleware and parser. So we can change the component to what we want anytime without changing code. This framework is inspired from [FeathersJS](https://feathersjs.com/). Currently framework still under development, any suggestion or PR is welcome.

# Getting Started

Currently package only tested in Go 1.13. To install the package 

```
$ go get -u github.com/oscrud/oscrud
```

After complete installation, you can Go with your beloved framework and here an hello world example. You can choose your own transport from lists, currently only supported Echo. For future will implement [service discovery](https://github.com/hashicorp/mdns) by default.

```go
package main

import (
    "github.com/oscrud/oscrud"
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

## Endpoint(method, path string, req *Request) TransportResponse

Mainly is for transport implementation purpose, but you can also use it as internal access call.

```go
server := oscrud.NewOscrud()
req := oscrud.NewRequest().
    Transport(t).
    Context(e).
    SetBody(body).
    SetQuery(query).
    SetHeader(header).
    SetParam(param)
res := server.Endpoint("GET", "/test", req)
```

## NewRequest(args ...string) *Request

Constructing new request for access endpoint.

```go
req := oscrud.NewRequest().
    Transport(t).
    Context(e).
    SetBody(body).
    SetQuery(query).
    SetHeader(header).
    SetParam(param)
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

## Example

For endpoint, must have a handler that calling `End()` method, if not will throwing error due to context haven't end yet. Why with this design, because of we make handler can be invoke until `After` middleware options and can be apply for profiler, logger about `context`.

```go
func Example(ctx oscrud.Context) oscrud.Context {
    return ctx.String(200, "Example Handler").End()
}

func main() {
    server := oscrud.NewOscrud()
    server.RegisterEndpoint("GET", "/example", Example)
}
```

## Context

`oscrud.Context` will have all of the information about `query`, `header`, `body` and more, usually is the information from incoming requests. Also `oscrud.Context` contains method for return data to outgoing response.

## Request Methods

For retrieving data from requests and some data binding.

| Method                         | Description                                                                                                                                  |
| ------------------------------ | -------------------------------------------------------------------------------------------------------------------------------------------- |
| Method() string                | Return request method, default to be smaller case `get`, `post`.                                                                             |
| Get(key string) interface{}    | Get value by key from  `param`, `query`, `body`, `header`, order respectively.                                                               |
| Context() interface{}          | Get request context, usually will be transport's context instance like `echo.Echo` instance.                                                 |
| Transport() string             | Get transport name, `INTERNAL` will be default if no transport.                                                                              |
| Path() string                  | Return request path                                                                                                                          |
| Headers() map[string]string    | Return request headers                                                                                                                       |
| Query() map[string]interface{} | Return request queries                                                                                                                       |
| Params() map[string]string     | Return request params                                                                                                                        |
| Body() map[string]interface{}  | Return request body                                                                                                                          |
| Bind(src interface{}) error    | Bind data from `map` specify by `reflect.Tag`. More information please look at [Specific Binding](#specific-binding).                        |
| BindAll(src interface{}) error | Bind data from `param`, `query`, `body`, `header` based on `json` and `qm` tag. More information please look at [All Binding](#all-binding). |


## Response Methods

For updating response information like data, headers.

| Method                                                      | Description                                                                                                                                           |
| ----------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| Set(key string, value string) Context                       | Append header to response, will override if key exists                                                                                                |
| NoContent() Context                                         | Response with status 204, and empty result.                                                                                                           |
| NotFounf() Context                                          | Response with status 404, and not found error.                                                                                                        |
| String(status int, text string) Context                     | Response with status, and raw string. Content type will be set as `text/plain`                                                                        |
| HTML(status int, html string) Context                       | Response with status, and html string. Content type will be set as `text/html`.                                                                       |
| JSON(status int, i interface{}) Context                     | Response with status, and interface. Content type will be set as `application/json`.                                                                  |
| XML(status int, i interface{}) Context                      | Response with status, and interface. Content type will be set as `application/xml`.                                                                   |
| Send(status int, contentType string, i interface{}) Context | Response with status, content type and interface. Transport should be handle on this if not will return error `ErrResponseFailed`.                    |
| Error(status int, exception error)                          | Response with status, and given error                                                                                                                 |
| Stack(status int, exception error)                          | Response with status, and given error ( stack will be provided ). `errors.WithStack()`.                                                               |
| End() Context                                               | A signal to tell handler, the flow is reach end already. If `End()` didn't call until end of the handler, will return error `ErrResponseNotComplete`. |

## Internal Access / Call

Mainly is for transport implementation, you can use it on mock tests or other internal access too, we not suggest to call the other `handler func` directly with given `context` unless you know what you're doing. You have to construct with `oscrud.NewRequest()` and pass `request` to `server.Endpoint()`

| Method                                         | Description                    |
| ---------------------------------------------- | ------------------------------ |
| SkipAfter() *Request                           | Skip after middleware          |
| SkipBefore() *Request                          | Skip before middleware         |
| SkipMiddleware() *Request                      | Skip all middleware            |
| Transport(trs Transport) *Request              | Set transport for request      |
| Context(ctx interface{}) *Request              | Set context for request        |
| SetBody(body map[string]interface{}) *Request  | Set body data for request      |
| SetParam(body map[string]string) *Request      | Set param data for request     |
| SetQuery(body map[string]interface{}) *Request | Set query data for request     |
| SetHeader(body map[string]header) *Request     | Set header data for request    |
| Header(key string, value string) *Request      | Append header data for request |
| Query(key string, value interface{}) *Request  | Append query data for request  |

### Example

Example is just for internal access / call, so you must have core struct `oscrud.Server` only you can access.

```go
server := oscrud.NewOscrud()
req := oscrud.NewRequest().
    Transport(t).
    Context(e).
    SetBody(body).
    SetQuery(query).
    SetHeader(header).
    SetParam(param)
res := server.Endpoint("GET", "/test", req)
```

### Transport Implementation

Transport implementation, based on incoming requests pass required data to request. Then using `handler` parameters to receive response from server.

```go
func (t *Transport) Register(method string, endpoint string, handler oscrud.TransportHandler) {
    req := oscrud.NewRequest(method, endpoint).
        Transport(t).
        Context(e).
        SetBody(body).
        SetQuery(query).
        SetHeader(header).
        SetParam(param)
    res := handler(req)
}
```

# Binder

Binder is for data binding when transform incoming requests data to a specified struct / type / array or slice. Currently customization only supported on `struct`, `slice` and `array`, primitive type will set by `reflect.Set()`.

## All Binding

All binding will get key from `reflect.Tag` and retrieve value from all maps, if exists will bind to struct. `qm` tag have higher priority then `json` so if `qm` tag exists, will take `qm` instead of `json`. Binding struct when calling must have addressable, if not will return error `ErrSourceNotAddressable`.

```go
var i struct {
    Id string `qm:"id"`
    Token string `json:"token" qm:"x-authorization"`
}

ctx.BindAll(&i)
```

## Specific Binding

Specific binding will bind value based on specified tag. If you want to bind from all with specified key, you can use all binding instead of specific binding. Binding struct when calling must have addressable, if not will return error `ErrSourceNotAddressable`.

* `header` will target to Header. 
* `query` will target to Query. 
* `body` will target to Body. 
* `param` will target to Param.

```go
var i struct {
    Token string `header:"token"`
    IsNew string `query:"isNew"`
    Username string `body:"username"`
    Password string `body:"password"`
    Id string `param:"id"`
}

ctx.Bind(&i)
```

## Register New Binder

Registering new binder for sepcific type ( struct, array or slice ), primitive not supported yet. 

```go
type Example struct {
    Line1 string
    Line2 string
}

func main() {
    // You use binder independently also
    binder := oscrud.NewBinder()
    binder.Register(new(Example), func(raw interface{}) (interface{}, error) {
        str := fmt.Sprintf("%v", raw)
        if strings.Contains(raw, ",") { 
            split := strings.Split(raw, ",")
            return Example{raw[0], raw[1]}, nil
        }
        return nil, fmt.Errorf("Invalid data "%v" for deserialize to Example", raw)
    })

    example := new(Example)
    err := binder.Bind(&example, "line1,line2")
    log.Println(example, err) // { line1, line2 }, <nil>


    err := binder.Bind(&example, "line1-line2")
    log.Println(example, err) // nil, "Invalid data line-line2 for deserialize to Example"
}
```

# Service

# Transport

# References & Resources

Reference and Resource about documentation like list of usable options or third party url.

## Content Type List

```go
// ContentType Definition
var (
	ContentTypePlainText = "text/plain"
	ContentTypeHTML      = "text/html"
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
)
```

## Errors List

| Error                   | Description                                                                             | Message                                       |
| ----------------------- | --------------------------------------------------------------------------------------- | --------------------------------------------- |
| ErrNotFound             | default endpoint not found error message & `ctx.NotFound()` error message               | endpoint or service not found                 |
| ErrResponseNotComplete  | error message when `End()` doesn't called and reach end of handler.                     | response doesn't called end in all handlers   |
| ErrResponseFailed       | error message when `Send()`'s information, transport unable to handler or not supported | response doesn't return properly in transport |
| ErrSourceNotAddressable | error message when binding source is not addressable                                    | binder source must be addressable             |
| ErrRequestTimeout       | default request timeout error message                                                   | request timeout                               |

## Available Options

### MiddlewareOptions

Middleware options is for applying `before` and `after` lifecycle to endpoint or server. So if both server and endpoint middleware specified, will run based on : ***Incoming Request --> `server.Before` -> `endpoint.Before` -> Main handler -> `endpoint.After` -> `server.After` -> Outgoing Response***.


```go
// MiddlewareOptions :
type MiddlewareOptions struct {
	Before []Handler
	After  []Handler
}
```

### EventOptions

Event options is for applying event-driven functionality, like `OnComplete`. For some of the transport may have bidirectional communicate, so have to remind when every requests complete. And this function will invoke using go-routine. As order will run `endpoint options` and only we run `server options`.

```go
event := oscrud.EventOptions{
	OnComplete: func(ctx oscrud.Context) {
		log.Println("This running from go-routine as event-drive OnComplete().")
	},
}
```

### TimeoutOptions

Timeout options is apply timeout for endpoint or server. Priority will take endpoint timeout options, if not specified will take server's timeout options. If none of them specified, will have a *default 30 seconds timeout* for every request.

```go
timeout := oscrud.TimeoutOptions{
	Duration: 1 * time.Microsecond,
	OnTimeout: func(ctx oscrud.Context) oscrud.Context {
		return ctx.Error(408, errors.New("Another requestimeout")).End()
	},
}
```


## Transport ( Official / Community )

* oscrud-echo

## Service ( Official / Community )

* oscrud-sqlike