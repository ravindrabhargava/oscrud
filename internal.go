package oscrud

// Endpoint :
func (server *Oscrud) Endpoint(method, endpoint string, req *Request) (*ResultResponse, *ErrorResponse) {
	req.method = method
	req.path = endpoint
	ctx := server.lookupHandler(nil, req)
	if ctx.exception != nil {
		return nil, ctx.exception
	}
	return ctx.result, nil
}
