package echo

import (
	"github.com/labstack/echo/v4"
	"oscrud"
)

// ServiceContext :
type ServiceContext struct {
	Context echo.Context
	Type    string
	ID      string
	Body    map[string]interface{}
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
func (c ServiceContext) GetBody() map[string]interface{} {
	return c.Body
}

// GetQuery :
func (c ServiceContext) GetQuery() map[string]interface{} {
	return c.Query
}

// Bind :
func (c ServiceContext) Bind(i interface{}) error {
	return oscrud.BindService(c.ID, c.Body, c.Query, i)
}