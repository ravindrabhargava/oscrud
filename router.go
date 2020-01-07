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

// MiddlewareOptions :
type MiddlewareOptions struct {
	Before []Handler
	After  []Handler
}

// EventOptions :
type EventOptions struct {
	OnComplete func(*ResultResponse, *ErrorResponse)
}
