package oscrud

import (
	"errors"
	"fmt"
	"strings"

	errs "github.com/pkg/errors"
)

// Response :
type Response struct {
	contentType     string
	responseHeaders map[string]string
	status          int
	exception       error
	result          interface{}
}

// Error Definition
var (
	ErrNotFound             = errors.New("endpoint or service not found")
	ErrResponseNotComplete  = errors.New("response doesn't called end in all handlers")
	ErrResponseFailed       = errors.New("response doesn't return properly in transport")
	ErrSourceNotAddressable = errors.New("binder source must be addressable")
	ErrRequestTimeout       = errors.New("request timeout")
)

// ContentType Definition
var (
	ContentTypePlainText = "text/plain"
	ContentTypeHTML      = "text/html"
	ContentTypeJSON      = "application/json"
	ContentTypeXML       = "application/xml"
)

// ContentType :
func (c Response) ContentType() string {
	return c.contentType
}

// ResponseHeaders :
func (c Response) ResponseHeaders() map[string]string {
	return c.responseHeaders
}

// Status :
func (c Response) Status() int {
	return c.status
}

// Exception :
func (c Response) Exception() error {
	return c.exception
}

// Result :
func (c Response) Result() interface{} {
	return c.result
}

// ErrorMap :
func (c Response) ErrorMap() map[string]interface{} {
	err := make(map[string]interface{})
	err["error"] = c.exception.Error()

	stack := strings.Split(strings.ReplaceAll(fmt.Sprintf("%+v", c.exception), "\t", ""), "\n")
	if len(stack) > 1 {
		err["stack"] = stack[2:]
	}
	return err
}

func (c Context) missingEnd() Context {
	c.response.status = 404
	c.response.exception = ErrResponseNotComplete
	return c
}

// Response :
func (c Context) Response() Response {
	return c.response
}

// NoContent :
func (c Context) NoContent() Context {
	c.response.status = 204
	c.response.result = nil
	return c
}

// String :
func (c Context) String(status int, text string) Context {
	c.response.status = status
	c.response.result = text
	c.response.contentType = ContentTypePlainText
	return c
}

// HTML :
func (c Context) HTML(status int, html string) Context {
	c.response.status = status
	c.response.result = html
	c.response.contentType = ContentTypeHTML
	return c
}

// JSON :
func (c Context) JSON(status int, i interface{}) Context {
	c.response.status = status
	c.response.result = i
	c.response.contentType = ContentTypeJSON
	return c
}

// XML :
func (c Context) XML(status int, i interface{}) Context {
	c.response.status = status
	c.response.result = i
	c.response.contentType = ContentTypeXML
	return c
}

// Send :
func (c Context) Send(status int, contentType string, i interface{}) Context {
	c.response.status = status
	c.response.result = i
	c.response.contentType = contentType
	return c
}

// Set :
func (c Context) Set(key string, value string) Context {
	c.response.responseHeaders[key] = value
	return c
}

// Error :
func (c Context) Error(status int, exception error) Context {
	c.response.status = status
	c.response.exception = exception
	return c
}

// Stack :
func (c Context) Stack(status int, exception error) Context {
	c.response.status = status
	c.response.exception = errs.WithStack(exception)
	return c
}

// End :
func (c Context) End() Context {
	c.sent = true
	return c
}

// NotFound :
func (c Context) NotFound() Context {
	c.response.status = 404
	c.response.exception = errs.WithStack(ErrNotFound)
	return c
}
