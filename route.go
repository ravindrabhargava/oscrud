package oscrud

import "time"

// Handler :
type Handler func(*Context) *Context

// Route :
type Route struct {
	MiddlewareOptions
	EventOptions
	TimeoutOptions
	TransportOptions

	Method  string
	Path    string
	Handler Handler
}

// Options :
type Options interface{}

// TimeoutOptions :
type TimeoutOptions struct {
	Duration  time.Duration
	OnTimeout Handler
}

// MiddlewareOptions :
type MiddlewareOptions struct {
	Before []Handler
	After  []Handler
}

// EventOptions :
type EventOptions struct {
	OnComplete func(*Context)
}

// TransportOptions :
type TransportOptions struct {
	DisableRegister map[TransportID]bool
}
