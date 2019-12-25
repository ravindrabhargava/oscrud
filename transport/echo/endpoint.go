package echo

import (
	"github.com/labstack/echo/v4"
	"oscrud"
)

// EndpointContext :
type EndpointContext struct {
	Context echo.Context
	Param   map[string]string
	Query   map[string]interface{}
	Body    map[string]interface{}
}

// Bind :
func (c EndpointContext) Bind(i interface{}) error {
	return oscrud.BindEndpoint(c.Param, c.Body, c.Query, i)
}

// GetContext :
func (c EndpointContext) GetContext() interface{} {
	return c.Context
}

// GetParams :
func (c EndpointContext) GetParams() map[string]string {
	return c.Param
}

// GetParam :
func (c EndpointContext) GetParam(key string) string {
	return c.Param[key]
}

// GetMethod :
func (c EndpointContext) GetMethod() string {
	return c.Context.Request().Method
}

// GetQuery :
func (c EndpointContext) GetQuery() map[string]interface{} {
	return c.Query
}

// GetTransport :
func (c EndpointContext) GetTransport() string {
	return "ECHO"
}

// GetPath :
func (c EndpointContext) GetPath() string {
	return c.Context.Path()
}

// GetBody :
func (c EndpointContext) GetBody() map[string]interface{} {
	return c.Body
}

// String :
func (c EndpointContext) String(status int, text string) error {
	return c.Context.String(status, text)
}

// HTML :
func (c EndpointContext) HTML(status int, html string) error {
	return c.Context.HTML(status, html)
}

// JSON :
func (c EndpointContext) JSON(status int, i interface{}) error {
	return c.Context.JSON(status, i)
}

// XML :
func (c EndpointContext) XML(status int, i interface{}) error {
	return c.Context.XML(status, i)
}
