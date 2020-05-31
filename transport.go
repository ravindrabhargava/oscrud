package oscrud

// TransportHandler :
type TransportHandler func(req *Request) Response

// Transport :
type Transport interface {
	Register(string, string, TransportHandler)
	Start() error
	Name() string
}
