package echo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oscrud/action"
	"strings"

	"github.com/labstack/echo/v4"
)

// Transport :
type Transport struct {
	Port int
	Echo *echo.Echo
}

// NewEcho :
func NewEcho(echo *echo.Echo) *Transport {
	return &Transport{
		Echo: echo,
		Port: 3000,
	}
}

// UsePort :
func (t *Transport) UsePort(port int) *Transport {
	t.Port = port
	return t
}

// RegisterService :
func (t *Transport) RegisterService(service, method, path string, handler action.ServiceHandler) {
	t.Echo.Add(
		strings.ToUpper(method), path,
		func(e echo.Context) error {
			bytes, err := ioutil.ReadAll(e.Request().Body)
			if err != nil {
				panic(err)
			}

			body := make(map[string]interface{})
			if e.Request().Method != "GET" {
				err = json.Unmarshal(bytes, &body)
				if err != nil {
					panic(err)
				}
			}

			query := make(map[string]interface{})
			for key, value := range e.Request().URL.Query() {
				if len(value) == 1 {
					query[key] = value[0]
				} else {
					query[key] = value
				}
			}

			ctx := ServiceContext{
				Context: e,
				Type:    service,
				ID:      e.Param("id"),
				Body:    body,
				Query:   query,
			}
			return handler(ctx)
		},
	)
}

// RegisterEndpoint :
func (t *Transport) RegisterEndpoint(method, path string, handler action.EndpointHandler) {
	t.Echo.Add(
		strings.ToUpper(method), path,
		func(e echo.Context) error {
			bytes, err := ioutil.ReadAll(e.Request().Body)
			if err != nil {
				panic(err)
			}

			body := make(map[string]interface{})
			if e.Request().Method != "GET" {
				err = json.Unmarshal(bytes, &body)
				if err != nil {
					panic(err)
				}
			}

			query := make(map[string]interface{})
			for key, value := range e.Request().URL.Query() {
				if len(value) == 1 {
					query[key] = value[0]
				} else {
					query[key] = value
				}
			}

			param := make(map[string]string)
			values := e.ParamValues()
			for index, name := range e.ParamNames() {
				param[name] = values[index]
			}

			ctx := EndpointContext{
				Context: e,
				Param:   param,
				Body:    body,
				Query:   query,
			}
			return handler(ctx)
		},
	)
}

// Start :
func (t *Transport) Start() error {
	port := fmt.Sprintf(":%d", t.Port)
	return t.Echo.Start(port)
}
