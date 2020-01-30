package oscrud

import (
	"errors"

	errs "github.com/pkg/errors"
)

// Error Definition
var (
	ErrNotFound            = errors.New("endpoint or service not found")
	ErrResponseNotComplete = errors.New("response doesn't called end in all handlers")
	ErrResponseFailed      = errors.New("response doesn't return properly in transport")
)

func (c Context) missingEnd() Context {
	c.status = 404
	c.exception = ErrResponseNotComplete
	return c
}

// NotFound :
func (c Context) NotFound() Context {
	c.status = 404
	c.exception = errs.WithStack(ErrNotFound)
	return c
}

// Error :
func (c Context) Error(status int, exception error) Context {
	c.status = status
	c.exception = exception
	return c
}

// Stack :
func (c Context) Stack(status int, exception error) Context {
	c.status = status
	c.exception = errs.WithStack(exception)
	return c
}
