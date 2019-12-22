package action

// EndpointHandler :
type EndpointHandler func(EndpointContext) error

// EndpointContext :
type EndpointContext interface {
	GetMethod() string
	GetTransport() string
	GetPath() string
	GetParam(key string) string
	GetBody() string
	GetQuery() map[string]interface{}
	ParseBody(body interface{}) error
	ParseQuery(query interface{}) error
}
