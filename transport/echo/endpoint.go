package echo

import (
	"oscrud/parser"

	"github.com/labstack/echo"
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

// GetMethod :
func (c EndpointContext) GetMethod() string {
	return c.Context.Request().Method
}

// GetQuery :
func (c EndpointContext) GetQuery() map[string]interface{} {
	return c.Query
}

// GetURL :
func (c EndpointContext) GetURL() string {
	return c.Context.Request().RequestURI
}

// GetTransport :
func (c EndpointContext) GetTransport() string {
	return "ECHO"
}

// GetPath :
func (c EndpointContext) GetPath() string {
	return c.Context.Path()
}

// ParseBody :
func (c EndpointContext) ParseBody(body interface{}) error {
	return c.Context.Bind(body)
}

// GetBody :
func (c EndpointContext) GetBody() string {
	return string(c.Body)
}
