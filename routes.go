package main

import (
	"net/http"

	"github.com/jefflub/allowance/handlers"
)

// Route is a route for the http handler
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes is the list of valid routes
type Routes []Route

var routes = Routes{
	Route{
		"CreateFamily",
		"POST",
		"/createfamily",
		handlers.CreateFamily,
	},
	Route{
		"TestServer",
		"GET",
		"/test",
		handlers.TestServer,
	},
	Route{
		"CreateKid",
		"POST",
		"/createkid",
		handlers.CreateKid,
	},
}
