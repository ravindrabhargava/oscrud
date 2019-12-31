package oscrud

// ContentType Definition
var (
	ContentTypePlainText = "text/plain"
	ContentTypeHTML      = "text/html"
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
)

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

// NoContent :
func (c Context) NoContent(status int) Context {
	c.result = &ResultResponse{
		status: status,
		result: nil,
	}
	return c
}

// String :
func (c Context) String(status int, text string) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      text,
		contentType: ContentTypePlainText,
	}
	return c
}

// HTML :
func (c Context) HTML(status int, html string) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      html,
		contentType: ContentTypeHTML,
	}
	return c
}

// JSON :
func (c Context) JSON(status int, i interface{}) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      i,
		contentType: ContentTypeJSON,
	}
	return c
}

// XML :
func (c Context) XML(status int, i interface{}) Context {
	c.result = &ResultResponse{
		status:      status,
		result:      i,
		contentType: ContentTypeXML,
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
