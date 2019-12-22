package echo

import (
	"oscrud/parser"

	"github.com/labstack/echo/v4"
)

// EndpointContext :
type EndpointContext struct {
	Echo    *echo.Echo
	Context echo.Context
	Query   map[string]interface{}
	Body    []byte
	Parser  []parser.Parser
}

// ParseQuery :
func (c EndpointContext) ParseQuery(assign interface{}) error {
	for index, parser := range c.Parser {
		err := parser.ParseQuery(c.Query, assign)
		if err == nil {
			return nil
		}

		if index == len(c.Parser) {
			return err
		}
	}
	return nil
}

// ParseBody :
func (c EndpointContext) ParseBody(body interface{}) error {
	return c.Context.Bind(body)
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
