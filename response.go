package oscrud

// ContentType Definition
var (
	ContentTypePlainText = "text/plain"
	ContentTypeHTML      = "text/html"
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
)

// NoContent :
func (c Context) NoContent() Context {
	c.status = 204
	c.result = nil
	return c
}

// String :
func (c Context) String(status int, text string) Context {
	c.status = status
	c.result = text
	c.contentType = ContentTypePlainText
	return c
}

// HTML :
func (c Context) HTML(status int, html string) Context {
	c.status = status
	c.result = html
	c.contentType = ContentTypeHTML
	return c
}

// JSON :
func (c Context) JSON(status int, i interface{}) Context {
	c.status = status
	c.result = i
	c.contentType = ContentTypeJSON
	return c
}

// XML :
func (c Context) XML(status int, i interface{}) Context {
	c.status = status
	c.result = i
	c.contentType = ContentTypeXML
	return c
}

// Send :
func (c Context) Send(status int, contentType string, i interface{}) Context {
	c.status = status
	c.result = i
	c.contentType = contentType
	return c
}
