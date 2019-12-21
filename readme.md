# Components

* Transport
* Service & Endpoint
* Authentication
* Middleware / Lifecycle [ Before, After ]
* Parser ( Query, Body(?) )

# Endpoints

* Method, Path => GET item

- GET /item => endpoint.Action

# Service

* BasePath => item

- GET /item -> service.find
- POST /item -> service.create
- GET /item/:id -> service.get
- PUT /item/:id -> service.update
- PATCH /item/:id -> service.patch
- DELETE /item/:id -> service.delete

# Example

Http still not working you may try using echo instead gohttp.

```
$ go run example/service.go
$ realize start
```

# Status

Still under development, any PR or suggestion is welcome.

# PR / Discussion / Suggestion 

* Cases
* Example

# Current Structure

* Parser current bind on Transport because Parser would be use when getting `body`, `query` from transport layer.
* Endpoint is a single action api.
* Service is a restful api standard api mostly CRUD endpoints.
* Route is currently like 

```json
{
    "GET": {
        "item/:id": (handler)
    },
    "POST": {
        "item": (handler)
    }
}
```

# Developer Experience

This still not my expected structure, any suggesiton or enhancement is welcome.

* Endpoint with EndpointContext mostly contain Request & Response related.
* Service with ServiceContext mostly contain Service related ( id, query, body ).