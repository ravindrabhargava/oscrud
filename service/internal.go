package service

import "oscrud/binder"

// Request :
type Request struct {
	service    string
	action     string
	identifier string
	body       map[string]interface{}
	query      map[string]interface{}
	header     map[string]interface{}
	result     *Response
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

// SetID :
func (c *Request) SetID(id string) *Request {
	c.identifier = id
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

// GetService :
func (c Request) GetService() string {
	return c.service
}

// GetContext :
func (c Request) GetContext() interface{} {
	return nil
}

// GetTransport :
func (c Request) GetTransport() string {
	return "INTERNAL"
}

// GetAction :
func (c Request) GetAction() string {
	return c.action
}

// GetID :
func (c Request) GetID() string {
	return c.identifier
}

// GetBody :
func (c Request) GetBody() map[string]interface{} {
	return c.body
}

// GetQuery :
func (c Request) GetQuery() map[string]interface{} {
	return c.query
}

// Bind :
func (c Request) Bind(i interface{}) error {
	return binder.BindService(c.identifier, c.body, c.query, c.header, i)
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
