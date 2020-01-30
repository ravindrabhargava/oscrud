package main

import (
	"log"
	"oscrud"
)

// Before :
func Before(ctx oscrud.Context) oscrud.Context {
	log.Println("I'm Before")
	return ctx
}

// After :
func After(ctx oscrud.Context) oscrud.Context {
	log.Println("I'm After")
	return ctx.End()
}
