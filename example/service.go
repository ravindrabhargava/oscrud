package main

import (
	"log"
	"oscrud"

	ec "oscrud/transport/echo"

	"github.com/labstack/echo/v4"
)

// // TestService :
// type TestService struct {
// }

// // NewService :
// func NewService() TestService {
// 	return TestService{}
// }

// // Find :
// func (t TestService) Find(service service.Context) error {
// 	log.Println("You're accessing TestService.Find")
// 	return nil
// }

// // Get :
// func (t TestService) Get(service service.Context) error {
// 	log.Println("You're accessing TestService.Get")
// 	return nil
// }

// // Create :
// func (t TestService) Create(service service.Context) error {
// 	log.Println("You're accessing TestService.Create")
// 	return nil
// }

// // Update :
// func (t TestService) Update(service service.Context) error {
// 	log.Println("You're accessing TestService.Update")
// 	return nil
// }

// // Patch :
// func (t TestService) Patch(service service.Context) error {
// 	log.Println("You're accessing TestService.Patch")
// 	return nil
// }

// // Remove :
// func (t TestService) Remove(service service.Context) error {
// 	log.Println("You're accessing TestService.Remove")
// 	return nil
// }

// Before :
func Before(ctx oscrud.Context) oscrud.Context {
	log.Println("I'm Before")
	return ctx
}

// Test2 :
func Test2(ctx oscrud.Context) oscrud.Context {
	var i struct {
		Test0 int    `param:"id"`
		Test2 uint64 `query:"test"`
		Test3 int32  `body:"test"`
	}

	err := ctx.Bind(&i)
	if err != nil {
		return ctx.Stack(500, err)
	}

	if i.Test0 == 0 {
		return ctx.Error(500, "ID should bigger than 0")
	}

	log.Println(i, err)
	log.Println("You're accessing Endpoint.")
	return ctx.String(200, "TestValue")
}

// After :
func After(ctx oscrud.Context) oscrud.Context {
	log.Println("I'm After")
	return ctx.End()
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterTransport(ec.NewEcho(echo.New()).UsePort(5001))

	event := oscrud.EventOptions{
		OnComplete: func(res *oscrud.ResultResponse, response *oscrud.ErrorResponse) {
			log.Println("This running from go-routine as event-drive OnComplete().")
		},
	}
	middleware := oscrud.MiddlewareOptions{
		Before: []oscrud.Handler{Before},
		After:  []oscrud.Handler{After},
	}

	server.RegisterEndpoint("GET", "/test2/:id/test", Test2, event, middleware)

	res, err := server.Endpoint("GET", "/test2/1/test", oscrud.NewRequest())
	log.Println(res, err)

	server.Start()
}
