package echo

import (
	"oscrud/binder"

	"github.com/labstack/echo/v4"
)

// EndpointContext :
type EndpointContext struct {
	context  echo.Context
	endpoint string
	param    map[string]string
	header   map[string]interface{}
	query    map[string]interface{}
	body     map[string]interface{}
}

// Bind :
func (c EndpointContext) Bind(i interface{}) error {
	return binder.BindEndpoint(c.header, c.param, c.body, c.query, i)
}

// GetContext :
func (c EndpointContext) GetContext() interface{} {
	return c.context
}

// GetEndpoint :
func (c EndpointContext) GetEndpoint() string {
	return c.endpoint
}

// Query :
func (c EndpointContext) Query(key string) interface{} {
	return c.query[key]
}

// Body :
func (c EndpointContext) Body(key string) interface{} {
	return c.body[key]
}

// Header :
func (c EndpointContext) Header(key string) interface{} {
	return c.header[key]
}

// Param :
func (c EndpointContext) Param(key string) string {
	return c.param[key]
}

// GetHeaders :
func (c EndpointContext) GetHeaders() map[string]interface{} {
	return c.header
}

// GetParams :
func (c EndpointContext) GetParams() map[string]string {
	return c.param
}

// GetMethod :
func (c EndpointContext) GetMethod() string {
	return c.context.Request().Method
}

// GetQuery :
func (c EndpointContext) GetQuery() map[string]interface{} {
	return c.query
}

// GetTransport :
func (c EndpointContext) GetTransport() string {
	return "ECHO"
}

// GetPath :
func (c EndpointContext) GetPath() string {
	return c.context.Path()
}

// GetBody :
func (c EndpointContext) GetBody() map[string]interface{} {
	return c.body
}

// String :
func (c EndpointContext) String(status int, text string) error {
	return c.context.String(status, text)
}

// HTML :
func (c EndpointContext) HTML(status int, html string) error {
	return c.context.HTML(status, html)
}

// JSON :
func (c EndpointContext) JSON(status int, i interface{}) error {
	return c.context.JSON(status, i)
}

// XML :
func (c EndpointContext) XML(status int, i interface{}) error {
	return c.context.XML(status, i)
}
