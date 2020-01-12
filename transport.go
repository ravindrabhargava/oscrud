package oscrud

import (
	"fmt"
	"strings"
)

// TransportResponse :
type TransportResponse struct {
	contentType     string
	responseHeaders map[string]string
	status          int
	exception       error
	result          interface{}
}

// TransportHandler :
type TransportHandler func(req *Request) TransportResponse

// Transport :
type Transport interface {
	Register(string, string, TransportHandler)
	Start(TransportHandler) error
	Name() string
}

// ContentType :
func (t TransportResponse) ContentType() string {
	return t.contentType
}

// Headers :
func (t TransportResponse) Headers() map[string]string {
	return t.responseHeaders
}

// Status :
func (t TransportResponse) Status() int {
	return t.status
}

// Error :
func (t TransportResponse) Error() error {
	return t.exception
}

// Result :
func (t TransportResponse) Result() interface{} {
	return t.result
}

// ErrorMap :
func (t TransportResponse) ErrorMap() map[string]interface{} {
	err := make(map[string]interface{})
	err["error"] = t.exception.Error()
	err["stack"] = strings.Split(strings.ReplaceAll(fmt.Sprintf("%+v", t.exception), "\t", ""), "\n")[2:]
	return err
}
