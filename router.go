package oscrud

// Handler :
type Handler func(Context) Context

// Route :
type Route struct {
	MiddlewareOptions
	EventOptions

	Method  string
	Route   string
	Path    string
	Handler Handler
}

// Options :
type Options interface{}

// MiddlewareOptions :
type MiddlewareOptions struct {
	Before []Handler
	After  []Handler
}

// EventOptions :
type EventOptions struct {
	OnComplete func(Context)
}
