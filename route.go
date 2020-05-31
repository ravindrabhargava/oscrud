package oscrud

import "time"

// Handler :
type Handler func(Context) Context

// Route :
type Route struct {
	MiddlewareOptions
	EventOptions
	TimeoutOptions

	Method  string
	Route   string
	Path    string
	Handler Handler
}

// Options :
type Options interface{}

// TimeoutOptions :
type TimeoutOptions struct {
	Duration  time.Duration
	OnTimeout func(Context) Context
}

// MiddlewareOptions :
type MiddlewareOptions struct {
	Before []Handler
	After  []Handler
}

// EventOptions :
type EventOptions struct {
	OnComplete func(Context)
}
