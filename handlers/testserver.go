package handlers

import (
	"database/sql"
	"log"
)

// TestServer is a request handler that provides a quick server and DB check
type TestServer struct {
	TestData  string `json:"testData" validate:"nonzero"`
	CheckDB   bool   `json:"checkDb"`
	FakeError bool   `json:"fakeError"`
	FakePanic bool   `json:"fakePanic"`
}

type testServerResponse struct {
	TestData string `json:"testData"`
	DBOK     bool   `json:"dbOk"`
}

// HandleRequest handles the request
func (t TestServer) HandleRequest() (interface{}, error) {
	var response testServerResponse
	response.TestData = t.TestData

	if t.FakePanic {
		panic("Fake panic!")
	}

	if t.FakeError {
		return nil, RequestError{"Fake Error!"}
	}

	// Simply write some test data for now
	db, err := sql.Open("mysql", "allowance_user:goniff@/allowance")
	if err != nil {
		response.DBOK = false
	}
	defer db.Close()

	if t.CheckDB {
		if err := db.Ping(); err != nil {
			log.Println(err)
			response.DBOK = false
		} else {
			response.DBOK = true
		}
	}

	return response, nil
}
