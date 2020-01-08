package oscrud

import (
	"errors"
	"fmt"

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

// ErrStack :
func (c Context) ErrStack(status int, result interface{}, stack error) Context {
	c.status = status
	c.result = result
	c.exception = stack
	return c
}

// Error
func (c Context) Error(status int, result interface{}) Context {
	c.status = status
	c.result = result
	return c
}

// Stack :
func (c Context) Stack(status int, stack error) Context {
	c.status = status
	c.exception = errs.WithStack(stack)
	return c
}

// Message :
func (c Context) Message(status int, message string) Context {
	c.status = status
	c.exception = fmt.Errorf(message)
	return c
}
