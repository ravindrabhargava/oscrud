package echo

import (
	"github.com/labstack/echo/v4"
)

// ServiceContext :
type ServiceContext struct {
	Context echo.Context
	Type    string
	ID      string
	Body    []byte
	Query   map[string]interface{}
}

// GetTransport :
func (c ServiceContext) GetTransport() string {
	return "ECHO"
}

// GetType :
func (c ServiceContext) GetType() string {
	return c.Type
}

// GetID :
func (c ServiceContext) GetID() string {
	return c.ID
}

// GetBody :
func (c ServiceContext) GetBody() string {
	return string(c.Body)
}

// Bind :
func (c ServiceContext) Bind(body interface{}) error {
	return c.Context.Bind(body)
}

// GetQuery :
func (c ServiceContext) GetQuery() map[string]interface{} {
	return c.Query
}
