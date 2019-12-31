package oscrud

import (
	"errors"
	"fmt"
	"strings"

	errs "github.com/pkg/errors"
)

// Error Definition
var (
	ErrNotFound            = errors.New("endpoint or service not found")
	ErrResponseNotComplete = errors.New("response doesn't called end in all handlers")
	ErrResponseFailed      = errors.New("response doesn't return properly in transport")
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

// ErrorMap :
func (c ErrorResponse) ErrorMap() map[string]interface{} {
	err := make(map[string]interface{})
	if c.stack != nil {
		err["message"] = c.stack.Error()
		err["stack"] = strings.Split(strings.ReplaceAll(fmt.Sprintf("%+v", c.stack), "\t", ""), "\n")
	}
	if c.result != nil {
		err["error"] = c.result
	}
	return err
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
		stack:  errs.WithStack(ErrNotFound),
	}
	return c
}

// ErrStack :
func (c Context) ErrStack(status int, result interface{}, stack error) Context {
	c.exception = &ErrorResponse{
		status: status,
		result: result,
		stack:  stack,
	}
	return c
}

// Error
func (c Context) Error(status int, result interface{}) Context {
	c.exception = &ErrorResponse{
		status: status,
		result: result,
	}
	return c
}

// Stack :
func (c Context) Stack(status int, stack error) Context {
	c.exception = &ErrorResponse{
		status: status,
		stack:  errs.WithStack(stack),
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
