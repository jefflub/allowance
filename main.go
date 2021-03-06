package main

import (
	"log"
	"net/http"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/mux"
	"github.com/jefflub/allowance/config"
	"github.com/jefflub/allowance/dbapi"
	"github.com/jefflub/allowance/handlers"
)

func main() {
	configFile := "config.toml"
	if len(os.Args) > 1 {
		configFile = os.Args[1]
	}
	log.Printf("Loading configuration file: %s", configFile)
	err := config.LoadConfig(configFile)
	if err != nil {
		log.Printf("Error loading config: %v", err)
		return
	}
	if err = dbapi.OpenDB(); err != nil {
		log.Printf("Error opening DB: %v", err)
		return
	}
	r := newRouter()

	// Fire up the server
	http.ListenAndServe("localhost:3000", r)
}

func newRouter() *mux.Router {
	router := mux.NewRouter().StrictSlash(true)
	sub := router.PathPrefix("/api").Subrouter()
	for _, route := range routes {
		var handler http.Handler
		handler = handlers.BaseHandler(route.Template, route.Name)

		sub.
			Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(handler)

	}
	return router
}
