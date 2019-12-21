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

Still under development, any PR or suggestion is welcome in format

* Cases
* Example