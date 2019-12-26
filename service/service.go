package service

// Handler :
type Handler func(Context) error

// Context :
type Context interface {
	GetTransport() string
	GetService() string
	GetAction() string
	GetID() string
	GetBody() map[string]interface{}
	GetQuery() map[string]interface{}
	GetContext() interface{}
	Bind(i interface{}) error

	String(status int, text string) error
	HTML(status int, html string) error
	JSON(status int, i interface{}) error
	XML(status int, i interface{}) error
}

// Service :
type Service interface {
	Find(Context) error
	Get(Context) error
	Create(Context) error
	Update(Context) error
	Patch(Context) error
	Remove(Context) error
}

// Route :
type Route struct {
	Service string
	Action  string
	Method  string
	Path    string
	Handler Handler
}

// Call :
func (r Route) Call(ctx *Request) (*Response, error) {
	ctx.service = r.Service
	ctx.action = r.Action

	err := r.Handler(*ctx)
	if err != nil {
		return nil, err
	}
	return ctx.result, nil
}
