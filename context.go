package oscrud

import (
	"errors"
	"oscrud/util"
	"reflect"
	"strings"
)

// Context :
type Context struct {
	method    string
	path      string
	query     map[string]interface{}
	body      map[string]interface{}
	param     map[string]string
	header    map[string]string
	context   interface{}
	transport Transport
	oscrud    Oscrud

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

// Method :
func (c Context) Method() string {
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
func (c Context) Transport() string {
	if c.transport == nil {
		return ""
	}
	return c.transport.Name()
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
func (c Context) Bind(assign interface{}) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return errors.New("binder interface must be addressable")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)
		var value interface{}

		htag := field.Tag.Get("header")
		if val, ok := c.header[htag]; ok {
			value = val
		}

		qtag := field.Tag.Get("query")
		if val, ok := c.query[qtag]; ok {
			value = val
		}

		btag := field.Tag.Get("body")
		if val, ok := c.body[btag]; ok {
			value = val
		}

		ptag := field.Tag.Get("param")
		if val, ok := c.param[ptag]; ok {
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
		return errors.New("binder interface must be addressable")
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	values := util.MergeMaps(c.header, c.query, c.body, c.param)
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

// End :
func (c Context) End() Context {
	c.sent = true
	return c
}
