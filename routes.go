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
		"AddParent",
		"POST",
		"/addparent",
		handlers.AddParent{},
	},
	Route{
		"CreateKid",
		"POST",
		"/createkid",
		handlers.CreateKid{},
	},
	Route{
		"GetBuckets",
		"POST",
		"/getbuckets",
		handlers.GetBuckets{},
	},
	Route{
		"AddMoney",
		"POST",
		"/addmoney",
		handlers.AddMoney{},
	},
	Route{
		"SpendMoney",
		"POST",
		"/spendmoney",
		handlers.SpendMoney{},
	},
	Route{
		"GetFamilySummary",
		"POST",
		"/getfamilysummary",
		handlers.GetFamilySummary{},
	},
	Route{
		"GetKidSummary",
		"POST",
		"/getkidsummary",
		handlers.GetKidSummary{},
	},
	Route{
		"Login",
		"POST",
		"/login",
		handlers.Login{},
	},
	Route{
		"LoginCheck",
		"POST",
		"/logincheck",
		handlers.LoginCheck{},
	},
	Route{
		"GetKidFromToken",
		"GET",
		"/kid/{token}/",
		handlers.GetKidFromToken{},
	},
	Route{
		"CreateKidToken",
		"POST",
		"/createkidtoken",
		handlers.CreateKidToken{},
	},
	Route{
		"GetLinkTokenInfo",
		"POST",
		"/getlinktokeninfo",
		handlers.GetLinkTokenInfo{},
	},
	Route{
		"DeleteLinkToken",
		"POST",
		"/deletelinktoken",
		handlers.DeleteLinkToken{},
	},
	Route{
		"GetTransactions",
		"POST",
		"/gettransactions",
		handlers.GetTransactions{},
	},
	Route{
		"EditKid",
		"POST",
		"/editkid",
		handlers.EditKid{},
	},
}
