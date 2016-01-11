package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"sort"

	"github.com/jefflub/allowance/dbapi"
	"gopkg.in/validator.v2"
)

type createFamilyParams struct {
	FamilyName     string `json:"familyName" validate:"nonzero"`
	ParentName     string `json:"parentName" validate:"nonzero"`
	ParentEmail    string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	ParentPassword string `json:"parentPassword" validate:"min=6"`
}

type createFamilyResponse struct {
	FamilyID int `json:"familyId"`
	ParentID int `json:"parentId"`
}

// CreateFamily is a handler that creates a family and the first parent
func CreateFamily(w http.ResponseWriter, r *http.Request) {
	var params createFamilyParams
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

	family, parent, err := dbapi.CreateFamily(params.FamilyName, params.ParentName, params.ParentEmail, params.ParentPassword)
	if err != nil {
		panic(err)
	}

	// Respond
	var response = new(createFamilyResponse)
	response.FamilyID = family.ID
	response.ParentID = parent.ID
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		panic(err)
	}
}
