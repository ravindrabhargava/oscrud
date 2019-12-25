package oscrud

// EndpointContext :
type EndpointContext struct {
	Endpoint string
	Method   string
	Path     string
	Param    map[string]string
	Body     map[string]interface{}
	Query    map[string]interface{}
	Result   *EndpointResult
}

// EndpointResult :
type EndpointResult struct {
	Status      int
	ContentType string
	Result      interface{}
}

// GetContext :
func (c EndpointContext) GetContext() interface{} {
	return nil
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
	c.Result = &EndpointResult{
		Status:      status,
		Result:      text,
		ContentType: "text/plain",
	}
	return nil
}

// HTML :
func (c EndpointContext) HTML(status int, html string) error {
	c.Result = &EndpointResult{
		Status:      status,
		Result:      html,
		ContentType: "text/html",
	}
	return nil
}

// JSON :
func (c EndpointContext) JSON(status int, i interface{}) error {
	c.Result = &EndpointResult{
		Status:      status,
		Result:      i,
		ContentType: "application/json",
	}
	return nil
}

// XML :
func (c EndpointContext) XML(status int, i interface{}) error {
	c.Result = &EndpointResult{
		Status:      status,
		Result:      i,
		ContentType: "application/xml",
	}
	return nil
}
