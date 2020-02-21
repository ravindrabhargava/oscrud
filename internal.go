package main

// Endpoint :
func (server *Oscrud) Endpoint(method, endpoint string, req *Request) TransportResponse {
	req.method = method
	req.path = endpoint
	ctx := server.lookupHandler(nil, req)
	return ctx.transportResponse()
}
