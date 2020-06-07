<!-- omit in toc -->
# Introduction

Oscrud is a golang resftul api wrapper framework. The purpose of the framework is make everything independent like transport, authentication, middleware and parser. So we can change the component to what we want anytime without changing code. This framework is inspired from [FeathersJS](https://feathersjs.com/). Currently framework still under development, any suggestion or PR is welcome.

<!-- omit in toc -->
# Documentation

- [Getting Started](#getting-started)
- [Oscrud Server](#oscrud-server)
	- [NewOscrud() *Oscrud](#newoscrud-oscrud)
	- [UseOptions(opts ...Options) *Oscrud](#useoptionsopts-options-oscrud)
	- [RegisterBinder(rtype interface{}, bindFb Bind) *Oscrud](#registerbinderrtype-interface-bindfb-bind-oscrud)
	- [RegisterTransport(transports ...Transport) *Oscrud](#registertransporttransports-transport-oscrud)
	- [RegisterLogger(loggers ...Logger) *Oscrud](#registerloggerloggers-logger-oscrud)
	- [RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud](#registerendpointmethod-endpoint-string-handler-handler-opts-options-oscrud)
	- [RegisterService(basePath string, service Service, opts ...Options) *Oscrud](#registerservicebasepath-string-service-service-opts-options-oscrud)
	- [NewRequest() *Request](#newrequest-request)
	- [SetState(key string, value interface{})](#setstatekey-string-value-interface)
	- [GetState(key string) interface{}](#getstatekey-string-interface)
	- [Log(operation string, content string)](#logoperation-string-content-string)
	- [Start()](#start)
- [Request Lifecycle](#request-lifecycle)
- [Handler & Context](#handler--context)
	- [Example](#example)
	- [Context](#context)
	- [Request Methods](#request-methods)
	- [Response Methods](#response-methods)
	- [Response Object](#response-object)
- [Binder](#binder)
	- [All Binding](#all-binding)
	- [Specific Binding](#specific-binding)
	- [Register New Binder](#register-new-binder)
- [Service](#service)
	- [Create Own Service](#create-own-service)
		- [Data Model](#data-model)
- [Transport](#transport)
	- [Create Own Transport](#create-own-transport)
- [References & Resources](#references--resources)
	- [Content Type List](#content-type-list)
	- [Errors List](#errors-list)
	- [Available Options](#available-options)
		- [MiddlewareOptions](#middlewareoptions)
		- [EventOptions](#eventoptions)
		- [TimeoutOptions](#timeoutoptions)
	- [Transport ( Official / Community )](#transport--official--community-)
	- [Service ( Official / Community )](#service--official--community-)
- [Discussion](#discussion)
	- [PR & Suggestion](#pr--suggestion)
	- [Issues](#issues)

# Getting Started

Currently package only tested in Go 1.13. To install the package 

```
$ go get -u github.com/oscrud/oscrud
```

After complete installation, you can Go with your beloved framework and here an hello world example. You can choose your own transport from lists, currently only supported Echo. For future will implement [service discovery](https://github.com/hashicorp/mdns) by default.

```go
package oscrud

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
package oscrud

server := oscrud.NewOscrud()
```

## UseOptions(opts ...Options) *Oscrud

For apply server-level options ( mean apply to all endpoints ).

```go
package oscrud

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
package oscrud

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
package oscrud

func main() {
	server := oscrud.NewOscrud()
    server.RegisterTransport(e.NewEcho(echo.New()).UsePort(3000))
}
```

## RegisterLogger(loggers ...Logger) *Oscrud

For register logger for the server. Every request made will be run all logger in goroutine, prevent for slowing down requests.

```go
package oscrud

// Logger :
type Logger struct {
}

// Log :
func (l Logger) Log(operation string, content string) {
	log.Println("Operation - ", operation)
	log.Println("Content - ", content)
}

// StartRequest :
func (l Logger) StartRequest(ctx oscrud.Context) {
	log.Println("**************************************")
	log.Println("RequestID - ", ctx.RequestID())
	log.Println("Method - ", ctx.Method())
	log.Println("Path - ", ctx.Path())
	log.Println("State - ", ctx.State())
	log.Println("Header - ", ctx.Headers())
	log.Println("Query - ", ctx.Query())
	log.Println("Body - ", ctx.Body())
	log.Println("**************************************")
}

// EndRequest :
func (l Logger) EndRequest(ctx oscrud.Context) {
	log.Println("**************************************")
	log.Println("RequestID - ", ctx.RequestID())
	log.Println("Method - ", ctx.Method())
	log.Println("Path - ", ctx.Path())
	log.Println("State - ", ctx.State())
	log.Println("Header - ", ctx.Headers())
	log.Println("Query - ", ctx.Query())
	log.Println("Body - ", ctx.Body())
	log.Println("**************************************")
}

func main() {
	server := oscrud.NewOscrud()
    server.RegisterLogger(Logger{})
}
```

## RegisterEndpoint(method, endpoint string, handler Handler, opts ...Options) *Oscrud 

For registering endpoint with specified method, endpoint & handler, and also able to apply endpoint level options ( mean only work on the specifed endpoint ).

```go
package oscrud

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
package oscrud

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

    // User is a query model struct based on oscrud.DataModel interface
	server.RegisterService("test", service.ToService("user", new(User)), middleware)
}
```

## NewRequest() *Request

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
## SetState(key string, value interface{})

Set application level state, you will have it when u have the server instance.

```go
oscrud.SetState("state", some_state)
```

## GetState(key string) interface{}

Get application level state, you will have it when u have the server instance.

```go
state := oscrud.GetState("state")
```

## Log(operation string, content string) 

Logging some operation with content, server context will run through loggers and call log function.

```go
oscrud.Log("Request", "request started")
```

## Start()

For start the oscrud server, server starting will only start all the registered transport. No transported registered will panic with err. No any setup will be invoke at this step, all would be done when `register`, so internal call / access can be work before or even not calling `Start()`.

```go
func main() {
	server := oscrud.NewOscrud()
	server.Start()
}
```

# Request Lifecycle

Basically a request firstly will come to Transport. Transport will do the basic handling to construct a `oscrud.Request` and only bring request to `oscrud`. 

1. Transport

Usually process incoming request and construct request for pass to handler.

2. Start Request Logger

3. Timeout Handler

Construct timeout handler & run handler using go-routine. When timeout reach will just return timeout error.

4. Oscrud

Lookup route, if exists construct middleware handler & main handler which required for the route.

5. Before Middleware Handler
6. Main Handler
7. After Middleware Handler
8. Event onComplete()
9. End Request Logger

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

| Method                                  | Description                                                                                                                                           |
| --------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| Method() string                         | Return request method, default to be smaller case `get`, `post`.                                                                                      |
| Get(key string) interface{}             | Get value by key from  `param`, `query`, `body`, `header`, order respectively.                                                                        |
| Context() context.Context               | Get request context                                                                                                                                   |
| Transport() string                      | Get transport name                                                                                                                                    |
| Path() string                           | Return request path                                                                                                                                   |
| RequestID() string                      | Return request id                                                                                                                                     |
| State() map[string]interface{}          | Return request state                                                                                                                                  |
| Headers() map[string]string             | Return request headers                                                                                                                                |
| Query() map[string]interface{}          | Return request queries                                                                                                                                |
| Params() map[string]string              | Return request params                                                                                                                                 |
| Body() map[string]interface{}           | Return request body                                                                                                                                   |
| Bind(src interface{}) error             | Bind data from `map` specify by `reflect.Tag`. More information please look at [Specific Binding](#specific-binding).                                 |
| BindAll(src interface{}) error          | Bind data from `param`, `query`, `body`, `header`, `state` based on `json` and `qm` tag. More information please look at [All Binding](#all-binding). |
| SetState(key string, value interface{}) | Set data to request level state                                                                                                                       |


## Response Methods

For updating response information like data, headers.

| Method                                                      | Description                                                                                                                                           |
| ----------------------------------------------------------- | ----------------------------------------------------------------------------------------------------------------------------------------------------- |
| Set(key string, value string) Context                       | Append header to response, will override if key exists                                                                                                |
| NoContent() Context                                         | Response with status 204, and empty result.                                                                                                           |
| NotFound() Context                                          | Response with status 404, and not found error.                                                                                                        |
| String(status int, text string) Context                     | Response with status, and raw string. Content type will be set as `text/plain`                                                                        |
| HTML(status int, html string) Context                       | Response with status, and html string. Content type will be set as `text/html`.                                                                       |
| JSON(status int, i interface{}) Context                     | Response with status, and interface. Content type will be set as `application/json`.                                                                  |
| XML(status int, i interface{}) Context                      | Response with status, and interface. Content type will be set as `application/xml`.                                                                   |
| Send(status int, contentType string, i interface{}) Context | Response with status, content type and interface. Transport should be handle on this if not will return error `ErrResponseFailed`.                    |
| Error(status int, exception error)                          | Response with status, and given error                                                                                                                 |
| Stack(status int, exception error)                          | Response with status, and given error ( stack will be provided ). `errors.WithStack()`.                                                               |
| End() Context                                               | A signal to tell handler, the flow is reach end already. If `End()` didn't call until end of the handler, will return error `ErrResponseNotComplete`. |
| Response() Response                                         | Return response object                                                                                                                                |

## Response Object

| Method                              | Description                  |
| ----------------------------------- | ---------------------------- |
| ContentType() string                | Return response content type |
| ResponseHeaders() map[string]string | Return response headers      |
| Status() int                        | Return response status       |
| Exception() error                   | Return response error        |
| Result() interface{}                | Return response result       |
| ErrorMap() map[string]interface{}   | Return pretty error          |


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
* `state` will target to Request Level State

```go
var i struct {
    Token string `header:"token"`
    IsNew string `query:"isNew"`
    Username string `body:"username"`
    Password string `body:"password"`
	Id string `param:"id"`
	State string `state:"state"`
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

Service has 6 action following [CRUD](https://en.wikipedia.org/wiki/Create,_read,_update_and_delete) standard. Basically service will registering 6 endpoints by default. Currently creating a service may required some basic knowledge on `reflect` package, we trying to minimize usage of reflect when creating own service.

* GET /basePath - Service.Find
* GET /basePath/:$id - Service.Get
* POST /basePath - Service.Create
* PUT /basePath/:$id - Service.Update
* PATCH /basePath/:$id - Service.Patch
* DELETE /basePath/:$id - Service.Delete

## Create Own Service

For creating own service, you must have implement methods based on `oscrud.Service` interface. There built in have 2 Query Struct ( `oscrud.Query`, `oscrud.QueryOne` ) for `bind` data from incoming requests, mainly for standardize query naming.

```go
type Service struct {}
func (service Service) Create(ctx oscrud.Context) oscrud.Context {
	// Createing data
}

func (service Service) Find(ctx oscrud.Context) oscrud.Context {
	// List data
}

func (service Service) Get(ctx oscrud.Context) oscrud.Context {
	// Get data ( should be one result always )
}

func (service Service) Update(ctx oscrud.Context) oscrud.Context {
	// update data
}

func (service Service) Patch(ctx oscrud.Context) oscrud.Context {
	// patch data
}

func (service Service) Delete(ctx oscrud.Context) oscrud.Context {
	// delete data
}
```

### Data Model

Service model is a model struct usually will be a table from database. Service model must have implmenet method from `oscrud.DataModel`. So when creating own service, we can use method to filter result or returning data even prevent toxic data injection. `$id` tag will automatically assign input value from endpoint, such as `GET /test/:$id` for a `Get` action.

| Method                 | Usage                                                                                                           |
| ---------------------- | --------------------------------------------------------------------------------------------------------------- |
| ToQuery() interface{}  | For returning query syntax based on service requirement, for sqlike is expr cosntruct from their query builder. |
| ToUpdate() interface{} | For construct model and return for update                                                                       |
| ToCreate() interface{} | For construct model and return for create                                                                       |
| ToResult() interface{} | For construct model and return for find / get                                                                   |


```go
// User :
type User struct {
	Key  int64  `json:"id" qm:"$id"`
	Name string `json:"name"`
}

func (user User) ToQuery() interface{} {
	var query interface{}

	if user.Key != 0 {
		query = expr.Equal("Key", user.Key)
	}

	return query
}

func (user *User) ToUpdate() interface{} {
	return user
}

func (user *User) ToCreate() interface{} {
	user.Name += "-NEW"
	return user
}

func (user *User) ToResult() interface{} {
	return user
}

```

# Transport

Transport act as communicate tools between client and server. Transport would have own name to determine request is from which transport.

## Create Own Transport

For creating own transport, you have to implement method from `oscrud.Transport`. Basically have 2 type registraion way. Way 1 is in `Register` for those transport which support endpoint routing like Echo. Way 2 is in `Start` for those transport doesn't have a router support like `net/http` package.

```go
type Transport struct {}

func (t *Transport) Name() string {
	return "TransportName"
}

func (t *Transport) Register(method string, endpoint string, handler oscrud.TransportHandler) {
	// Every endpoint registration will call this method.
}

func (t *Transport) Start() error {
	// Transport start receiving request
}
```

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

* oscrud/echo

## Service ( Official / Community )

* oscrud/sqlike

# Discussion

Since we're still new, any suggestion and pull requests are welcomed. Currently we haven't setup any social for disscussion will be update later. Any discussion should follow template by providing a clearer information for everyone.

## PR & Suggestion

* Cases
* Example
* Solving Issues

## Issues

* Simple example
* Way to reproduce issues
* Issues about 
* Version