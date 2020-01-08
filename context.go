package oscrud

// Context :
type Context struct {
	method          string
	path            string
	query           map[string]interface{}
	body            map[string]interface{}
	param           map[string]string
	header          map[string]string
	sent            bool
	context         interface{}
	transport       Transport
	responseHeaders map[string]string
	result          *ResultResponse
	exception       *ErrorResponse
}

// GetMethod :
func (c Context) GetMethod() string {
	return c.method
}

// Get :
func (c Context) Get(key string) interface{} {

	if val, ok := c.param[key]; ok {
		return val
	}

	if val, ok := c.query[key]; ok {
		return val
	}

	if val, ok := c.body[key]; ok {
		return val
	}

	if val, ok := c.header[key]; ok {
		return val
	}

	return nil
}

// Set :
func (c Context) Set(key string, value string) Context {
	c.responseHeaders[key] = value
	return c
}

// Context :
func (c Context) Context() interface{} {
	return c.context
}

// Transport :
func (c Context) Transport() Transport {
	return c.transport
}

// Path :
func (c Context) Path() string {
	return c.path
}

// Headers :
func (c Context) Headers() map[string]string {
	return c.header
}

// Query :
func (c Context) Query() map[string]interface{} {
	return c.query
}

// Params :
func (c Context) Params() map[string]string {
	return c.param
}

// Body :
func (c Context) Body() map[string]interface{} {
	return c.body
}

// Bind :
func (c Context) Bind(i interface{}) error {
	return bind(c.header, c.param, c.body, c.query, i)
}

// End :
func (c Context) End() Context {
	c.sent = true
	return c
}
