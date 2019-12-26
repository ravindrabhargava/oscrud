package endpoint

import "oscrud/binder"

// Request :
type Request struct {
	endpoint string
	method   string
	path     string
	param    map[string]string
	body     map[string]interface{}
	query    map[string]interface{}
	header   map[string]interface{}
	result   *Response
}

// Response :
type Response struct {
	status      int
	contentType string
	result      interface{}
}

// NewRequest :
func NewRequest() *Request {
	return &Request{}
}

// SetParam :
func (c *Request) SetParam(param map[string]string) *Request {
	c.param = param
	return c
}

// SetBody :
func (c *Request) SetBody(body map[string]interface{}) *Request {
	c.body = body
	return c
}

// SetQuery :
func (c *Request) SetQuery(query map[string]interface{}) *Request {
	c.query = query
	return c
}

// SetHeader :
func (c *Request) SetHeader(header map[string]interface{}) *Request {
	c.header = header
	return c
}

// GetEndpoint :
func (c Request) GetEndpoint() string {
	return c.endpoint
}

// GetContext :
func (c Request) GetContext() interface{} {
	return nil
}

// GetMethod :
func (c Request) GetMethod() string {
	return c.method
}

// GetTransport :
func (c Request) GetTransport() string {
	return "INTERNAL"
}

// GetPath :
func (c Request) GetPath() string {
	return c.path
}

// GetHeaders :
func (c Request) GetHeaders() map[string]interface{} {
	return c.header
}

// GetQuery :
func (c Request) GetQuery() map[string]interface{} {
	return c.query
}

// GetParams :
func (c Request) GetParams() map[string]string {
	return c.param
}

// GetBody :
func (c Request) GetBody() map[string]interface{} {
	return c.body
}

// Bind :
func (c Request) Bind(i interface{}) error {
	return binder.BindEndpoint(c.header, c.param, c.body, c.query, i)
}

// Header :
func (c Request) Header(key string) interface{} {
	return c.header[key]
}

// Query :
func (c Request) Query(key string) interface{} {
	return c.query[key]
}

// Body :
func (c Request) Body(key string) interface{} {
	return c.body[key]
}

// Param :
func (c Request) Param(key string) string {
	return c.param[key]
}

// String :
func (c Request) String(status int, text string) error {
	c.result = &Response{
		status:      status,
		result:      text,
		contentType: "text/plain",
	}
	return nil
}

// HTML :
func (c Request) HTML(status int, html string) error {
	c.result = &Response{
		status:      status,
		result:      html,
		contentType: "text/html",
	}
	return nil
}

// JSON :
func (c Request) JSON(status int, i interface{}) error {
	c.result = &Response{
		status:      status,
		result:      i,
		contentType: "application/json",
	}
	return nil
}

// XML :
func (c Request) XML(status int, i interface{}) error {
	c.result = &Response{
		status:      status,
		result:      i,
		contentType: "application/xml",
	}
	return nil
}
