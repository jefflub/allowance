package main

import "github.com/jefflub/allowance/handlers"

// Route is a route for the http handler
type Route struct {
	Name     string
	Method   string
	Pattern  string
	Template handlers.RequestHandler
}

// Routes is the list of valid routes
type Routes []Route

var routes = Routes{
	/*
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
		},*/
	Route{
		"TestServer",
		"POST",
		"/test",
		handlers.TestServer{},
	},
	Route{
		"CreateFamily",
		"POST",
		"/createfamily",
		handlers.CreateFamily{},
	},
	Route{
		"Login",
		"POST",
		"/login",
		handlers.Login{},
	},
}
