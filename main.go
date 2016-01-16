package main

import (
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jefflub/allowance/handlers"
)

func main() {
	fmt.Println("Hello")

	r := newRouter()

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		var handler http.Handler
		handler = handlers.BaseHandler(route.Template, route.Name)

		router.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
