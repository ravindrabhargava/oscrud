package main

import (
	"net/http"

	"github.com/oscrud/oscrud"
)

// GoRoutine :
func GoRoutine(ctx *oscrud.Context) *oscrud.Context {
	return ctx.String(http.StatusOK, "This is an example")
}
