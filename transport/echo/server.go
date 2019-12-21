package echo

import (
	"fmt"
	"io/ioutil"
	"oscrud/action"
	"oscrud/parser"

	"github.com/labstack/echo"
)

// Transport :
type Transport struct {
	Echo   *echo.Echo
	Port   int
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

// Register :
func (t *Transport) Register(method string, path string, handler action.Handler) {
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

			srv := ServiceContext{}
			ctx := EndpointContext{
				Body:    body,
				Query:   query,
				Parser:  t.Parser,
				Context: e,
				Echo:    t.Echo,
			}
			handler(ctx, srv)
			return nil
		},
	)
}

// Start :
func (t *Transport) Start(h action.RequestHandler) error {
	port := fmt.Sprintf(":%d", t.Port)
	return t.Echo.Start(port)
}
