package echo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"oscrud"
	"strings"

	"github.com/labstack/echo/v4"
)

// Transport Definition
var (
	TransportName = "ECHO"
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

// Name :
func (t *Transport) Name() string {
	return TransportName
}

// Register :
func (t *Transport) Register(method string, endpoint string, handler oscrud.TransportHandler) {
	t.Echo.Add(
		strings.ToUpper(method), endpoint,
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

			header := make(map[string]string)
			for key, value := range e.Request().Header {
				if len(value) == 1 {
					header[key] = value[0]
				} else {
					header[key] = strings.Join(value, ",")
				}
			}

			param := make(map[string]string)
			values := e.ParamValues()
			for index, name := range e.ParamNames() {
				param[name] = values[index]
			}

			req := oscrud.NewRequest(method, endpoint).
				Transport(t).
				SetBody(body).
				SetQuery(query).
				SetHeader(header).
				SetParam(param)

			response := handler(req)

			for key, val := range response.Headers() {
				e.Response().Header().Add(key, val)
			}

			if response.Error() != nil {
				return e.JSON(response.Status(), response.ErrorMap())
			}

			if response.Result() == nil {
				return e.NoContent(response.Status())
			}
			if response.ContentType() == oscrud.ContentTypePlainText {
				return e.String(response.Status(), response.Result().(string))
			}
			if response.ContentType() == oscrud.ContentTypeHTML {
				return e.HTML(response.Status(), response.Result().(string))
			}
			if response.ContentType() == oscrud.ContentTypeXML {
				return e.XML(response.Status(), response.Result())
			}
			if response.ContentType() == oscrud.ContentTypeJSON {
				return e.JSON(response.Status(), response.Result())
			}
			return oscrud.ErrResponseFailed
		},
	)
}

// Start :
func (t *Transport) Start(handler oscrud.TransportHandler) error {
	port := fmt.Sprintf(":%d", t.Port)
	return t.Echo.Start(port)
}
