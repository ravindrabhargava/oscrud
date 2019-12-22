package echo

import (
	"fmt"
	"io/ioutil"
	"oscrud/action"
	"oscrud/parser"

	"github.com/labstack/echo/v4"
)

// Transport :
type Transport struct {
	Port   int
	Echo   *echo.Echo
	Parser []parser.Parser
}

// NewEcho :
func NewEcho(echo *echo.Echo) *Transport {
	return &Transport{
		Echo:   echo,
		Port:   3000,
		Parser: make([]parser.Parser, 0),
	}
}

// UsePort :
func (t *Transport) UsePort(port int) *Transport {
	t.Port = port
	return t
}

// UseParser :
func (t *Transport) UseParser(parser parser.Parser) *Transport {
	t.Parser = append(t.Parser, parser)
	return t
}

// RegisterService :
func (t *Transport) RegisterService(service, method, path string, handler action.ServiceHandler) {
	t.Echo.Add(
		method, path,
		func(e echo.Context) error {
			body, err := ioutil.ReadAll(e.Request().Body)
			if err != nil {
				panic(err)
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
				Type:   service,
				ID:     e.Param("id"),
				Body:   body,
				Parser: t.Parser,
				Echo:   t.Echo,
			}
			return handler(ctx)
		},
	)
}

// RegisterEndpoint :
func (t *Transport) RegisterEndpoint(method, path string, handler action.EndpointHandler) {
	t.Echo.Add(
		method, path,
		func(e echo.Context) error {
			body, err := ioutil.ReadAll(e.Request().Body)
			if err != nil {
				panic(err)
			}

			query := make(map[string]interface{})
			for key, value := range e.Request().URL.Query() {
				if len(value) == 1 {
					query[key] = value[0]
				} else {
					query[key] = value
				}
			}
			ctx := EndpointContext{
				Body:    body,
				Query:   query,
				Parser:  t.Parser,
				Context: e,
				Echo:    t.Echo,
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
