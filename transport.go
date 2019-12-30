package oscrud

// TransportHandler :
type TransportHandler func(req *Request) (*ResultResponse, *ErrorResponse)

// Transport :
type Transport interface {
	Register(string, string, TransportHandler)
	Start(TransportHandler) error
}
