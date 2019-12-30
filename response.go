package oscrud

// ResultResponse :
type ResultResponse struct {
	status      int
	contentType string
	result      interface{}
}

// Status :
func (c ResultResponse) Status() int {
	return c.status
}

// ContentType :
func (c ResultResponse) ContentType() string {
	return c.contentType
}

// Result :
func (c ResultResponse) Result() interface{} {
	return c.result
}

// String :
func (c Context) String(status int, text string) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      text,
		contentType: "text/plain",
	}
	return c
}

// HTML :
func (c Context) HTML(status int, html string) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      html,
		contentType: "text/html",
	}
	return c
}

// JSON :
func (c Context) JSON(status int, i interface{}) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      i,
		contentType: "application/json",
	}
	return c
}

// XML :
func (c Context) XML(status int, i interface{}) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      i,
		contentType: "application/xml",
	}
	return c
}

// Send :
func (c Context) Send(status int, contentType string, i interface{}) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      i,
		contentType: contentType,
	}
	return c
}
