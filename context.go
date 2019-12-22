package oscrud

import (
	"encoding/json"
	"oscrud/parser"
)

// EndpointContext :
type EndpointContext struct {
	Method string
	Path   string
	Param  map[string]string
	Body   map[string]interface{}
	Query  map[string]interface{}
	Parser []parser.Parser
}

// GetMethod :
func (c EndpointContext) GetMethod() string {
	return c.Method
}

// GetTransport :
func (c EndpointContext) GetTransport() string {
	return "INTERNAL"
}

// GetPath :
func (c EndpointContext) GetPath() string {
	return c.Path
}

// GetQuery :
func (c EndpointContext) GetQuery() map[string]interface{} {
	return c.Query
}

// GetParam :
func (c EndpointContext) GetParam(key string) string {
	return c.Param[key]
}

// GetBody :
func (c EndpointContext) GetBody() string {
	bytes, err := json.Marshal(c.Body)
	if err != nil {
		return ""
	}
	return string(bytes)
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
	bytes, err := json.Marshal(c.Body)
	if err != nil {
		return err
	}
	return json.Unmarshal(bytes, body)
}
