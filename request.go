package oscrud

// Request Skip Definition
var (
	skipMiddleware = "ALL"
	skipBefore     = "BEFORE"
	skipAfter      = "AFTER"
	skipNone       = "NONE"
)

// Request :
type Request struct {
	method    string
	transport Transport
	path      string
	query     map[string]interface{}
	body      map[string]interface{}
	param     map[string]string
	header    map[string]string
	skip      string
	context   interface{}
}

// NewRequest :
func NewRequest(args ...string) *Request {
	req := &Request{transport: nil, skip: skipNone}
	if len(args) == 2 {
		req.method = args[0]
		req.path = args[1]
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
func (req *Request) Context(ctx interface{}) *Request {
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
