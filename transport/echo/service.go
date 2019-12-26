package echo

import (
	"oscrud/binder"

	"github.com/labstack/echo/v4"
)

// ServiceContext :
type ServiceContext struct {
	context echo.Context
	service string
	action  string
	id      string
	header  map[string]interface{}
	body    map[string]interface{}
	query   map[string]interface{}
}

// Bind :
func (c ServiceContext) Bind(i interface{}) error {
	return binder.BindService(c.id, c.body, c.query, c.header, i)
}

// GetService :
func (c ServiceContext) GetService() string {
	return c.service
}

// GetContext :
func (c ServiceContext) GetContext() interface{} {
	return c.context
}

// GetTransport :
func (c ServiceContext) GetTransport() string {
	return "ECHO"
}

// GetAction :
func (c ServiceContext) GetAction() string {
	return c.action
}

// GetID :
func (c ServiceContext) GetID() string {
	return c.id
}

// GetBody :
func (c ServiceContext) GetBody() map[string]interface{} {
	return c.body
}

// GetQuery :
func (c ServiceContext) GetQuery() map[string]interface{} {
	return c.query
}

// String :
func (c ServiceContext) String(status int, text string) error {
	return c.context.String(status, text)
}

// HTML :
func (c ServiceContext) HTML(status int, html string) error {
	return c.context.HTML(status, html)
}

// JSON :
func (c ServiceContext) JSON(status int, i interface{}) error {
	return c.context.JSON(status, i)
}

// XML :
func (c ServiceContext) XML(status int, i interface{}) error {
	return c.context.XML(status, i)
}
