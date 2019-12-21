package echo

import "oscrud/parser"

import "github.com/labstack/echo/v4"

// ServiceContext :
type ServiceContext struct {
	Type   string
	ID     string
	Body   []byte
	Echo   *echo.Echo
	Parser []parser.Parser
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
