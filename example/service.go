// package main

// import (
// 	"errors"
// 	"log"
// 	"oscrud"
// 	"time"

// 	"oscrud/service/sqlike"
// 	ec "oscrud/transport/echo"

// 	sql "github.com/si3nloong/sqlike/sqlike"
// 	"github.com/si3nloong/sqlike/sqlike/options"

// 	_ "github.com/go-sql-driver/mysql"
// 	"github.com/labstack/echo/v4"
// )

// // AnyStruct :
// type AnyStruct struct {
// 	Data string
// }

// // Test2 :
// func Test2(ctx oscrud.Context) oscrud.Context {
// 	var i struct {
// 		Test0  int       `param:"id"`
// 		Test2  uint64    `query:"test"`
// 		Test3  int32     `body:"test"`
// 		Struct AnyStruct `query:"any_struct"`
// 	}

// 	err := ctx.Bind(&i)
// 	if err != nil {
// 		return ctx.Stack(500, err)
// 	}

// 	if i.Test0 == 0 {
// 		return ctx.Error(500, errors.New("ID should bigger than 0"))
// 	}
// 	log.Println("Binder AnyStruct : ", i.Struct)
// 	return ctx.String(200, "You're accessing endpoint.")
// }

// func main() {
// 	server := oscrud.NewOscrud()

// 	// Register transport
// 	server.RegisterTransport(
// 		ec.NewEcho(echo.New()).UsePort(3001),
// 	)

// 	// Register data binding for specific struct / slice / array.
// 	server.RegisterBinder(
// 		AnyStruct{},
// 		func(raw interface{}) (interface{}, error) {
// 			str, ok := raw.(string)
// 			if ok {
// 				return AnyStruct{str}, nil
// 			}
// 			return nil, errors.New("received data isn't a string")
// 		},
// 	)

// 	// Endpoint options definition ( usually be middleware )
// timeout := oscrud.TimeoutOptions{
// 	Duration: 1 * time.Microsecond,
// 	OnTimeout: func(ctx oscrud.Context) oscrud.Context {
// 		return ctx.Error(408, errors.New("Another requestimeout")).End()
// 	},
// }

// event := oscrud.EventOptions{
// 	OnComplete: func(ctx oscrud.Context) {
// 		log.Println("This running from go-routine as event-drive OnComplete().")
// 	},
// }
// middleware := oscrud.MiddlewareOptions{
// 	Before: []oscrud.Handler{Before},
// 	After:  []oscrud.Handler{After},
// }

// 	// Server level options registration
// 	server.UseOptions(timeout)

// 	// Register Endpoint
// 	server.RegisterEndpoint("GET", "/test2/:id/test", Test2, event, middleware)

// 	// Internal Call
// 	req := oscrud.NewRequest().Query("any_struct", "223")
// 	res := server.Endpoint("GET", "/test2/1/test", req)
// 	log.Println(res.Result(), res.Error())

// 	res = server.Endpoint("GET", "/test2/0/test", oscrud.NewRequest())
// 	log.Println(res.Result(), res.Error())

// 	// Sqlike database conn initialize
// 	client := sql.MustConnect("mysql",
// 		options.Connect().
// 			SetHost("localhost").
// 			SetPort("3306").
// 			SetUsername("root").
// 			SetPassword("test"),
// 	)
// 	client.SetPrimaryKey("Key")

// 	// Service Definition
// 	service := sqlike.NewService(client).Database("test")
// 	server.RegisterService("test", service.ToService("user", new(User)))

// 	// Everything done? Start the server.
// 	server.Start()
// }

package main

import (
	"github.com/labstack/echo/v4"
	"oscrud"
	e "oscrud/transport/echo"
)

func main() {
	server := oscrud.NewOscrud()
	server.RegisterTransport(
		e.NewEcho(echo.New()).UsePort(3000),
	)
	server.RegisterEndpoint("GET", "/test", func(ctx oscrud.Context) oscrud.Context {
		return ctx.String(200, "Hello World").End()
	})
	server.Start()
}
