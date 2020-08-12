package oscrud

import (
	"context"
	"encoding/json"
	"mime/multipart"
	"net/url"
	"reflect"
	"strings"

	"github.com/oscrud/oscrud/util"
)

// Context :
type Context struct {
	oscrud   Oscrud
	request  Request
	response Response
	sent     bool
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

	if val, ok := c.request.form[key]; ok {
		return val
	}

	return nil
}

// ParseForm :
func (c Context) ParseForm(multipart bool) error {
	return c.request.formHandler(multipart)
}

// File :
func (c Context) File(key string) (*multipart.FileHeader, error) {
	return c.request.fileHandler(key)
}

// Context :
func (c Context) Context() context.Context {
	return c.request.context
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

// Form :
func (c Context) Form() url.Values {
	return c.request.form
}

// BinderSetting :
type BinderSetting struct {
	UseJSON bool
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

		stag := field.Tag.Get("state")
		if val, ok := c.request.param[stag]; ok {
			value = val
		}

		ftag := field.Tag.Get("form")
		if val, ok := c.request.param[ftag]; ok {
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
func (c Context) BindAll(assign interface{}, setting ...BinderSetting) error {
	t := reflect.TypeOf(assign)
	if t.Kind() != reflect.Ptr && t.Elem().Kind() != reflect.Struct {
		return ErrSourceNotAddressable
	}

	values := util.MergeMaps(
		c.request.header,
		c.request.query,
		c.request.body,
		c.request.param,
		c.request.state,
		c.request.form,
	)

	if len(setting) > 0 && setting[0].UseJSON {
		bytes, err := json.Marshal(values)
		if err != nil {
			return err
		}
		return json.Unmarshal(bytes, assign)
	}

	setter := reflect.ValueOf(assign).Elem()
	npt := t.Elem()
	for i := 0; i < npt.NumField(); i++ {
		field := npt.Field(i)
		var key string

		json := field.Tag.Get("json")
		if json != "" {
			key = strings.Split(json, ",")[0]
		}

		oscrudTag := field.Tag.Get("oscrud")
		if oscrudTag != "" {
			key = oscrudTag
		}

		if key != "" {
			if value, ok := values[key]; ok && value != nil {
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

// SetState :
func (c Context) SetState(key string, value interface{}) {
	c.request.State(key, value)
}

// GetState :
func (c Context) GetState(key string) interface{} {
	return c.request.state[key]
}

// State :
func (c Context) State() map[string]interface{} {
	return c.request.state
}

// RequestID  :
func (c Context) RequestID() string {
	return c.request.requestID
}

// Host :
func (c Context) Host() string {
	return c.request.host
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
func (c Context) Transport() TransportID {
	return c.request.transport.Name()
}

// Oscrud :
func (c Context) Oscrud() Oscrud {
	return c.oscrud
}
