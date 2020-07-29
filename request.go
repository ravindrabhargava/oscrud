package oscrud

import (
	"context"

	"github.com/google/uuid"
)

// Request Skip Definition
var (
	skipMiddleware = "ALL"
	skipBefore     = "BEFORE"
	skipAfter      = "AFTER"
	skipNone       = "NONE"
)

// Request :
type Request struct {
	transport Transport
	context   context.Context

	requestID string
	host      string
	method    string
	path      string
	state     map[string]interface{}
	query     map[string]interface{}
	body      map[string]interface{}
	param     map[string]string
	header    map[string]string
	skip      string
}

// NewRequest :
func NewRequest() *Request {
	req := &Request{
		transport: nil,
		requestID: uuid.New().String(),
		skip:      skipNone,
		state:     make(map[string]interface{}),
		query:     make(map[string]interface{}),
		body:      make(map[string]interface{}),
		param:     make(map[string]string),
		header:    make(map[string]string),
	}
	return req
}

// SkipAfter :
func (req *Request) SkipAfter() *Request {
	req.skip = skipAfter
	return req
}

// SkipBefore :
func (req *Request) SkipBefore() *Request {
	req.skip = skipBefore
	return req
}

// SkipMiddleware :
func (req *Request) SkipMiddleware() *Request {
	req.skip = skipMiddleware
	return req
}

// Transport :
func (req *Request) Transport(trs Transport) *Request {
	req.transport = trs
	return req
}

// Context :
func (req *Request) Context(ctx context.Context) *Request {
	req.context = ctx
	return req
}

// SetBody :
func (req *Request) SetBody(body map[string]interface{}) *Request {
	req.body = body
	return req
}

// SetParam :
func (req *Request) SetParam(param map[string]string) *Request {
	req.param = param
	return req
}

// SetQuery :
func (req *Request) SetQuery(query map[string]interface{}) *Request {
	req.query = query
	return req
}

// SetHeader :
func (req *Request) SetHeader(header map[string]string) *Request {
	req.header = header
	return req
}

// SetState :
func (req *Request) SetState(state map[string]interface{}) *Request {
	req.state = state
	return req
}

// SetHost :
func (req *Request) SetHost(host string) *Request {
	req.host = host
	return req
}

// Query :
func (req *Request) Query(key string, value interface{}) *Request {
	req.query[key] = value
	return req
}

// Header :
func (req *Request) Header(key string, value string) *Request {
	req.header[key] = value
	return req
}

// Param :
func (req *Request) Param(key string, value string) *Request {
	req.param[key] = value
	return req
}

// State :
func (req *Request) State(key string, value interface{}) *Request {
	req.state[key] = value
	return req
}
