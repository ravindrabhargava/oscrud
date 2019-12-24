package action

// EndpointHandler :
type EndpointHandler func(EndpointContext) error

// EndpointContext :
type EndpointContext interface {
	GetMethod() string
	GetTransport() string
	GetPath() string
	GetParam(key string) string
	GetParams() map[string]string
	GetBody() map[string]interface{}
	GetQuery() map[string]interface{}
	Bind(i interface{}) error

	String(status int, text string) error
	HTML(status int, html string) error
	JSON(status int, i interface{}) error
	XML(status int, i interface{}) error
	Redirect(status int, url string) error
}
