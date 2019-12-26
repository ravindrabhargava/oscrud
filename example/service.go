package main

import (
	"log"
	"oscrud"
	"oscrud/endpoint"
	"oscrud/service"
	ec "oscrud/transport/echo"
	socketio "oscrud/transport/socketio"

	"github.com/labstack/echo/v4"
)

// TestService :
type TestService struct {
}

// NewService :
func NewService() TestService {
	return TestService{}
}

// Find :
func (t TestService) Find(service service.Context) error {
	log.Println("You're accessing TestService.Find")
	return nil
}

// Get :
func (t TestService) Get(service service.Context) error {
	log.Println("You're accessing TestService.Get")
	return nil
}

// Create :
func (t TestService) Create(service service.Context) error {
	log.Println("You're accessing TestService.Create")
	return nil
}

// Update :
func (t TestService) Update(service service.Context) error {
	log.Println("You're accessing TestService.Update")
	return nil
}

// Patch :
func (t TestService) Patch(service service.Context) error {
	log.Println("You're accessing TestService.Patch")
	return nil
}

// Remove :
func (t TestService) Remove(service service.Context) error {
	log.Println("You're accessing TestService.Remove")
	return nil
}

// Test2 :
func Test2(ctx endpoint.Context) error {
	var i struct {
		Test0 int    `param:"id"`
		Test1 string `param:"test"`
		Test2 uint64 `query:"test"`
		Test3 int32  `body:"test"`
	}

	err := ctx.Bind(&i)
	log.Println(i, err)
	log.Println("You're accessing Endpoint.")
	ctx.JSON(200, i)
	return nil
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterService("test", "test", NewService())
	server.RegisterEndpoint("test2", "GET", "/test2/:id/:test", Test2)
	server.RegisterTransport(
		ec.NewEcho(echo.New()).UsePort(5001),
		socketio.NewSocket(nil).UsePort(5000),
	)

	req := service.NewRequest()
	server.Service("test").Find(req)

	param := map[string]string{
		"id":   "12",
		"test": "1",
	}
	body := map[string]interface{}{
		"test": 100,
	}
	query := map[string]interface{}{
		"test": 2000,
	}
	req2 := endpoint.NewRequest().SetParam(param).SetBody(body).SetQuery(query)
	server.Endpoint("test2", req2)

	server.Start()
}
