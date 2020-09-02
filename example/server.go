package main

import (
	"github.com/oscrud/oscrud"
)

func main() {
	server := oscrud.NewOscrud()
	server.RegisterEndpoint("GET", "/example", Example)
	server.RegisterEndpoint("GET", "/goroutine", GoRoutine)
	server.Start()
}
