package dbapi

import (
	"database/sql"
	"errors"

	"github.com/jefflub/allowance/config"
)

var defaultBuckets = []Bucket{
	Bucket{0, "Spending", 80, 0.0},
	Bucket{0, "Saving", 10, 0.0},
	Bucket{0, "Giving", 10, 0.0},
}

// CreateKid creates a kid and their buckets in the DB
func CreateKid(familyID int, name string, email string, buckets []Bucket) (Kid, error) {
	kid := Kid{0, name, email, buckets}
	if len(buckets) == 0 {
		kid.Buckets = defaultBuckets
	} else {
		var allocationSum = 0
		for _, b := range buckets {
			if b.DefaultAllocation < 0 {
				return kid, errors.New("Invalid default allocation < 0")
			}
			allocationSum += b.DefaultAllocation
		}
		if allocationSum != 100 {
			return kid, errors.New("Invalid bucket default allocations. Must sum to 100.")
		}
	}

	db, err := sql.Open("mysql", config.GetConfig().DbURL)
	if err != nil {
		return kid, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return kid, err
	}
	defer tx.Rollback()

	// Create the family
	if _, err := tx.Exec("INSERT INTO kids VALUES(NULL, ?, ?, ?, NULL, NULL)", familyID, name, email); err != nil {
		return kid, err
	}
	// Get the ID
	row := tx.QueryRow("SELECT LAST_INSERT_ID()")
	if err := row.Scan(&kid.ID); err != nil {
		return kid, err
	}

	// Add buckets
	for _, b := range kid.Buckets {
		if _, err := tx.Exec("INSERT INTO buckets VALUES(NULL, ?, ?, ?, NULL, NULL)", kid.ID, b.Name, b.DefaultAllocation); err != nil {
			return kid, err
		}
	}
	tx.Commit()

	return kid, nil
}
