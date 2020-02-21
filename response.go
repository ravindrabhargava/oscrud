package oscrud

import (
	"errors"

	errs "github.com/pkg/errors"
)

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

// Set :
func (c Context) Set(key string, value string) Context {
	c.responseHeaders[key] = value
	return c
}

// End :
func (c Context) End() Context {
	c.sent = true
	return c
}

func (c Context) missingEnd() Context {
	c.status = 404
	c.exception = ErrResponseNotComplete
	return c
}

// Error Definition
var (
	ErrNotFound             = errors.New("endpoint or service not found")
	ErrResponseNotComplete  = errors.New("response doesn't called end in all handlers")
	ErrResponseFailed       = errors.New("response doesn't return properly in transport")
	ErrSourceNotAddressable = errors.New("binder source must be addressable")
	ErrRequestTimeout       = errors.New("request timeout")
)

// NotFound :
func (c Context) NotFound() Context {
	c.status = 404
	c.exception = errs.WithStack(ErrNotFound)
	return c
}

// Error :
func (c Context) Error(status int, exception error) Context {
	c.status = status
	c.exception = exception
	return c
}

// Stack :
func (c Context) Stack(status int, exception error) Context {
	c.status = status
	c.exception = errs.WithStack(exception)
	return c
}
