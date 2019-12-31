package oscrud

import (
	"errors"
	"fmt"
)

// Error Definition
var (
	ErrNotFound            = errors.New("endpoint or service not found")
	ErrResponseNotComplete = errors.New("response doesn't called end in all handlers")
)

func (c Context) missingEnd() Context {
	c.exception = &ErrorResponse{
		status: 404,
		stack:  ErrResponseNotComplete,
	}
	return c
}

// ErrorResponse :
type ErrorResponse struct {
	status int
	stack  error
	result interface{}
}

// Status :
func (c ErrorResponse) Status() int {
	return c.status
}

// Stack :
func (c ErrorResponse) Stack() error {
	return c.stack
}

// Result :
func (c ErrorResponse) Result() interface{} {
	return c.result
}

// NotFound :
func (c Context) NotFound() Context {
	c.exception = &ErrorResponse{
		status: 404,
		stack:  ErrNotFound,
	}
	return c
}

// Error :
func (c Context) Error(status int, stack error) Context {
	c.exception = &ErrorResponse{
		status: status,
		stack:  stack,
	}
	return c
}

// Message :
func (c Context) Message(status int, message string) Context {
	c.exception = &ErrorResponse{
		status: status,
		stack:  fmt.Errorf(message),
	}
	return c
}
