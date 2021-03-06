package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"reflect"
	"time"

	"github.com/gorilla/mux"

	"gopkg.in/validator.v2"
)

// RequestError is an error for a request
type RequestError struct {
	Message string `json:"errorMsg"`
}

func (r RequestError) Error() string {
	return r.Message
}

// RequestHandler is a type that will get a request
type RequestHandler interface {
	HandleRequest(vars map[string]string) (interface{}, error)
}

// BaseHandler contains base functionality for all handlers
func BaseHandler(inner RequestHandler, name string) http.HandlerFunc {
	t := reflect.TypeOf(inner)
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		var err error
		// Handle request logging
		defer func() {
			status := "SUCCESS"
			var logmsg interface{}
			logmsg = ""
			e := recover()
			if e != nil {
				status = "PANIC"
				logmsg = e
			} else {
				if err != nil {
					status = "ERROR"
					logmsg = err
				}
			}
			log.Printf(
				"%s\t%s\t%s\t%s\t%s\t%v",
				status,
				r.Method,
				r.RequestURI,
				name,
				time.Since(start),
				logmsg,
			)
		}()

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")

		// Construct a new parameter object for this request
		pPtr := reflect.New(t)
		p := pPtr.Interface()
		if r.Method == "POST" {
			var body []byte
			body, err = ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
			if err == nil {
				if err = r.Body.Close(); err == nil {
					if err = json.Unmarshal(body, &p); err == nil {
						// Validate the json
						err = validator.Validate(p)
					}
				}
			}
		}

		// Run the request
		if err == nil {
			var response interface{}
			if response, err = p.(RequestHandler).HandleRequest(mux.Vars(r)); err == nil {
				w.WriteHeader(200)
				if err = json.NewEncoder(w).Encode(response); err != nil {
					panic(err)
				}
			}
		}

		// Send an error if necessary
		if err != nil {
			err = RequestError{fmt.Sprintf("Request Error: %s", err.Error())}
			w.WriteHeader(422) // unprocessable entity
			if jerr := json.NewEncoder(w).Encode(err); jerr != nil {
				panic(jerr)
			}
		}
	})
}
