package oscrud

// TransportResponse :
type TransportResponse struct {
	Result  *ResultResponse
	Error   *ErrorResponse
	Headers map[string]string
}

// TransportHandler :
type TransportHandler func(req *Request) TransportResponse

// Transport :
type Transport interface {
	Register(string, string, TransportHandler)
	Start(TransportHandler) error
}
