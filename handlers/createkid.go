package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"

	"gopkg.in/validator.v2"

	"github.com/jefflub/allowance/dbapi"
)

type createKidParams struct {
	FamilyID int            `json:"familyId" validate:"nonzero"`
	KidName  string         `json:"name" validate:"nonzero"`
	Buckets  []dbapi.Bucket `json:"buckets"`
}

type createKidResponse struct {
	KidID int `json:"kidId"`
}

// CreateKid creates a kid, using default buckets or provided buckets
func CreateKid(w http.ResponseWriter, r *http.Request) {
	var params createKidParams
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		panic(err)
	}
	if err := r.Body.Close(); err != nil {
		panic(err)
	}
	if err := json.Unmarshal(body, &params); err != nil {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			panic(err)
		}
	}

	// Validate
	if err := validator.Validate(params); err != nil {
		errs := err.(validator.ErrorMap)
		// Iterate through the list of fields and respective errors
		fmt.Println("Invalid due to fields:")
		// Here we have to sort the arrays to ensure map ordering does not
		// fail our example, typically it's ok to just range through the err
		// list when order is not important.
		var errOuts []string
		for f, e := range errs {
			errOuts = append(errOuts, fmt.Sprintf("\t - %s (%v)\n", f, e))
		}

		// Again this part is extraneous and you should not need this in real
		// code.
		sort.Strings(errOuts)
		for _, str := range errOuts {
			fmt.Print(str)
		}
	}

	// Respond
	var response = new(createKidResponse)
	response.KidID = 100
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(params); err != nil {
		panic(err)
	}
}
