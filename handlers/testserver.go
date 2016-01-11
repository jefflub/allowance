package handlers

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

// TestServer makes sure the server is running and connected to the db
func TestServer(w http.ResponseWriter, r *http.Request) {
	// Simply write some test data for now
	db, err := sql.Open("mysql", "allowance_user:goniff@/allowance")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	err = db.Ping()
	if err == nil {
		fmt.Fprintln(w, "DB connected!")
	} else {
		fmt.Fprintln(w, err)
	}

	fmt.Fprint(w, "Welcome!\n")

}
