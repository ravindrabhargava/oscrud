package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"oscrud/action"
	"oscrud/parser"
)

// Transport :
type Transport struct {
	Port   int
	Parser []parser.Parser
}

// NewHTTP :
func NewHTTP() *Transport {
	return &Transport{
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

}

// Start :
func (t *Transport) Start(h action.RequestHandler) error {
	port := fmt.Sprintf(":%d", t.Port)
	return http.ListenAndServe(port, &httpHandler{t, h})
}

type httpHandler struct {
	setting *Transport
	handler action.RequestHandler
}

func (h httpHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err)
	}

	query := make(map[string]interface{})
	for key, value := range req.URL.Query() {
		if len(value) == 1 {
			query[key] = value[0]
		} else {
			query[key] = value
		}
	}

	srv := ServiceContext{}
	ctx := EndpointContext{
		Method: req.Method,
		Parser: h.setting.Parser,
		Path:   req.URL.Path,
		URL:    req.URL.String(),
		Body:   body,
		Query:  query,
	}
	h.handler(nil)(ctx, srv)
	header := res.Header()
	header.Add("Content-Type", "application/json")
	res.WriteHeader(200)
}
