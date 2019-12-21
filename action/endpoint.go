package action

// EndpointContext :
type EndpointContext interface {
	GetMethod() string
	GetURL() string
	GetPath() string
	GetTransport() string
	ParseBody(body interface{}) error
	ParseQuery(query interface{}) error
	GetBody() string
	GetQuery() map[string]interface{}
}

// Endpoint :
type Endpoint interface {
	Action(ctx EndpointContext) error
}
