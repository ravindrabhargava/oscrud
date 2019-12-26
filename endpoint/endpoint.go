package endpoint

// Handler :
type Handler func(Context) error

// Context :
type Context interface {
	GetEndpoint() string
	GetMethod() string
	GetTransport() string
	GetPath() string

	Header(key string) interface{}
	Query(key string) interface{}
	Param(key string) string
	Body(key string) interface{}

	GetHeaders() map[string]interface{}
	GetParams() map[string]string
	GetBody() map[string]interface{}
	GetQuery() map[string]interface{}
	GetContext() interface{}
	Bind(i interface{}) error

	String(status int, text string) error
	HTML(status int, html string) error
	JSON(status int, i interface{}) error
	XML(status int, i interface{}) error
}

// Route :
type Route struct {
	Endpoint string
	Method   string
	Path     string
	Handler  Handler
}

// Call :
func (r Route) Call(ctx *Request) (*Response, error) {
	ctx.method = r.Method
	ctx.path = r.Path
	ctx.endpoint = r.Endpoint

	err := r.Handler(*ctx)
	if err != nil {
		return nil, err
	}
	return ctx.result, nil
}
