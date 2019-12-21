package http

import (
	"encoding/json"
	"oscrud/parser"
)

// EndpointContext :
type EndpointContext struct {
	Method string
	Path   string
	URL    string
	Body   []byte
	Param  map[string]string
	Query  map[string]interface{}
	Parser []parser.Parser
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
	return c.Method
}

// GetQuery :
func (c EndpointContext) GetQuery() map[string]interface{} {
	return c.Query
}

// GetURL :
func (c EndpointContext) GetURL() string {
	return c.URL
}

// GetTransport :
func (c EndpointContext) GetTransport() string {
	return "GOHTTP"
}

// GetPath :
func (c EndpointContext) GetPath() string {
	return c.Path
}

// ParseBody :
func (c EndpointContext) ParseBody(body interface{}) error {
	return json.Unmarshal(c.Body, body)
}

// GetBody :
func (c EndpointContext) GetBody() string {
	return string(c.Body)
}
