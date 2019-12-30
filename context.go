package oscrud

// Context :
type Context struct {
	method    string
	transport string
	path      string
	header    map[string]interface{}
	query     map[string]interface{}
	body      map[string]interface{}
	param     map[string]string
	sent      bool
	result    *ResultResponse
	exception *ErrorResponse
}

// GetMethod :
func (c Context) GetMethod() string {
	return c.method
}

// GetTransport :
func (c Context) GetTransport() string {
	return "INTERNAL"
}

// GetPath :
func (c Context) GetPath() string {
	return c.path
}

// GetHeaders :
func (c Context) GetHeaders() map[string]interface{} {
	return c.header
}

// GetQuery :
func (c Context) GetQuery() map[string]interface{} {
	return c.query
}

// GetParams :
func (c Context) GetParams() map[string]string {
	return c.param
}

// GetBody :
func (c Context) GetBody() map[string]interface{} {
	return c.body
}

// Bind :
func (c Context) Bind(i interface{}) error {
	return Bind(c.header, c.param, c.body, c.query, i)
}

// Header :
func (c Context) Header(key string) interface{} {
	return c.header[key]
}

// Query :
func (c Context) Query(key string) interface{} {
	return c.query[key]
}

// Body :
func (c Context) Body(key string) interface{} {
	return c.body[key]
}

// Param :
func (c Context) Param(key string) string {
	return c.param[key]
}

// End :
func (c Context) End() Context {
	c.sent = true
	return c
}
