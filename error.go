package oscrud

import "fmt"

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
