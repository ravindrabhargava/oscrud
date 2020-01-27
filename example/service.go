package main

import (
	"errors"
	"log"
	"oscrud"

	"oscrud/service/sqlike"
	ec "oscrud/transport/echo"
	sc "oscrud/transport/socketio"

	sql "github.com/si3nloong/sqlike/sqlike"
	"github.com/si3nloong/sqlike/sqlike/options"

	_ "github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
)

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
		return ctx.Error(500, errors.New("ID should bigger than 0"))
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

type AnyStruct struct {
	Data string
}

func main() {
	server := oscrud.NewOscrud()
	server.RegisterTransport(
		ec.NewEcho(echo.New()).UsePort(5001),
		sc.NewSocket(nil).UsePort(3000),
	)

	event := oscrud.EventOptions{
		OnComplete: func(ctx oscrud.Context) {
			log.Println("This running from go-routine as event-drive OnComplete().")
		},
	}
	middleware := oscrud.MiddlewareOptions{
		Before: []oscrud.Handler{Before},
		After:  []oscrud.Handler{After},
	}

	server.RegisterEndpoint("GET", "/test2/:id/test", Test2, event, middleware)

	res := server.Endpoint("GET", "/test2/1/test", oscrud.NewRequest())
	log.Println(res.Result(), res.Error())

	res = server.Endpoint("GET", "/test2/0/test", oscrud.NewRequest())
	log.Println(res.Result(), res.Error())

	client := sql.MustConnect("mysql",
		options.Connect().
			SetHost("localhost").
			SetPort("3306").
			SetUsername("root").
			SetPassword("test"),
	)
	client.SetPrimaryKey("Key")
	service := sqlike.NewService(client).Database("test")
	server.RegisterService("test", service.ToService("user", new(User)))
	server.Start()
}
