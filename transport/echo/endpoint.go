package echo

import (
	"github.com/labstack/echo/v4"
)

// EndpointContext :
type EndpointContext struct {
	Context echo.Context
	Query   map[string]interface{}
	Body    []byte
}

// Bind :
func (c EndpointContext) Bind(i interface{}) error {
	return c.Context.Bind(i)
}

// GetParam :
func (c EndpointContext) GetParam(key string) string {
	return c.Context.Param(key)
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
func (c EndpointContext) GetBody() string {
	return string(c.Body)
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

// Redirect :
func (c EndpointContext) Redirect(status int, url string) error {
	return c.Context.Redirect(status, url)
}
