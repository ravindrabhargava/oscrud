package oscrud

import (
	"reflect"
	"strings"

	"github.com/oscrud/oscrud/util"
)

// Context :
type Context struct {
	request Request
	oscrud  Oscrud

	sent            bool
	contentType     string
	responseHeaders map[string]string
	status          int
	exception       error
	result          interface{}
}

func (c Context) transportResponse() TransportResponse {
	return TransportResponse{
		contentType:     c.contentType,
		responseHeaders: c.responseHeaders,
		status:          c.status,
		exception:       c.exception,
		result:          c.result,
	}
}

// Get :
func (c Context) Get(key string) interface{} {

	if val, ok := c.request.param[key]; ok {
		return val
	}

	if val, ok := c.request.query[key]; ok {
		return val
	}

	if val, ok := c.request.body[key]; ok {
		return val
	}

	if val, ok := c.request.header[key]; ok {
		return val
	}

	return nil
}

// Headers :
func (c Context) Headers() map[string]string {
	return c.request.header
}

// Query :
func (c Context) Query() map[string]interface{} {
	return c.request.query
}

// Params :
func (c Context) Params() map[string]string {
	return c.request.param
}

// Body :
func (c Context) Body() map[string]interface{} {
	return c.request.body
}

// Bind :
func (c Context) Bind(assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return ErrSourceNotAddressable
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)
		var value interface{}

		htag := field.Tag.Get("header")
		if val, ok := c.request.header[htag]; ok {
			value = val
		}

		qtag := field.Tag.Get("query")
		if val, ok := c.request.query[qtag]; ok {
			value = val
		}

		btag := field.Tag.Get("body")
		if val, ok := c.request.body[btag]; ok {
			value = val
		}

		ptag := field.Tag.Get("param")
		if val, ok := c.request.param[ptag]; ok {
			value = val
		}

		if value != nil {
			if err := c.oscrud.binder.Bind(setter.Field(i).Addr().Interface(), value); err != nil {
				return err
			}
		}
	}

	return nil
}

// BindAll :
func (c Context) BindAll(assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return ErrSourceNotAddressable
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	values := util.MergeMaps(
		c.request.header,
		c.request.query,
		c.request.body,
		c.request.param,
	)
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)
		var key string

		json := field.Tag.Get("json")
		if json != "" {
			key = strings.Split(json, ",")[0]
		}

		qm := field.Tag.Get("qm")
		if qm != "" {
			key = qm
		}

		if key != "" {
			if value, ok := values[key]; ok {
				if err := c.oscrud.binder.Bind(setter.Field(i).Addr().Interface(), value); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

// Log :
func (c Context) Log(operation string, content string) {
	c.oscrud.Log(operation, content)
}

// State :
func (c Context) State() map[string]interface{} {
	return c.request.state
}

// RequestID  :
func (c Context) RequestID() string {
	return c.request.requestID
}

// Method :
func (c Context) Method() string {
	return c.request.method
}

// Path :
func (c Context) Path() string {
	return c.request.path
}

// Transport :
func (c Context) Transport() string {
	return c.request.transport.Name()
}
