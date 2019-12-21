package http

import "oscrud/parser"

// ServiceContext :
type ServiceContext struct {
	Method string
	Path   string
	URL    string
	Body   []byte
	Param  map[string]string
	Query  map[string]interface{}
	Parser []parser.Parser
}
