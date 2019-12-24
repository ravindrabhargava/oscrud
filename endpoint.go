package oscrud

// EndpointContext :
type EndpointContext struct {
	Method string
	Path   string
	Param  map[string]string
	Body   map[string]interface{}
	Query  map[string]interface{}
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

// GetParams :
func (c EndpointContext) GetParams() map[string]string {
	return c.Param
}

// GetParam :
func (c EndpointContext) GetParam(key string) string {
	return c.Param[key]
}

// GetBody :
func (c EndpointContext) GetBody() map[string]interface{} {
	return c.Body
}

// Bind :
func (c EndpointContext) Bind(i interface{}) error {
	return BindEndpoint(c.Param, c.Body, c.Query, i)
}

// String :
func (c EndpointContext) String(status int, text string) error {
	return nil
}

// HTML :
func (c EndpointContext) HTML(status int, html string) error {
	return nil
}

// JSON :
func (c EndpointContext) JSON(status int, i interface{}) error {
	return nil
}

// XML :
func (c EndpointContext) XML(status int, i interface{}) error {
	return nil
}

// Redirect :
func (c EndpointContext) Redirect(status int, url string) error {
	return nil
}
