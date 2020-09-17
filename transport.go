package oscrud

// TransportHandler :
type TransportHandler func(req *Request) Response

// TransportID :
type TransportID string

// Transport :
type Transport interface {
	Register(string, string, TransportHandler)
	Request(*Request, interface{}) error
	Start() error
	Name() TransportID
}

// function 'create_uuid' will generate the new uniqe id everytime whenever func will be call

func create_uuid() {
	fmt.Println(uuid.New().String())
}