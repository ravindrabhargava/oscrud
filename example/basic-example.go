package main

import (
	"net/http"

	"github.com/oscrud/oscrud"
)

// Example :
func Example(ctx *oscrud.Context) *oscrud.Context {
	return ctx.String(http.StatusOK, "This is an example")
}
