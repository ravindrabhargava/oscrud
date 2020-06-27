package oscrud

// TransportHandler :
type TransportHandler func(req *Request) Response

// TransportID :
type TransportID string

// Transport :
type Transport interface {
	Register(string, string, TransportHandler)
	Start() error
	Name() TransportID
}
