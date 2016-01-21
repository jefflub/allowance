package dbapi

import (
	"database/sql"
	"time"

	"github.com/jefflub/allowance/config"
)

var db *sql.DB

// OpenDB creates the DB object that will be used by everyone else
func OpenDB() error {
	var err error
	db, err = sql.Open("mysql", config.GetConfig().DbURL)
	if err != nil {
		return err
	}
	return nil
}

type (
	// Family represents the structure of our resource
	Family struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
		Kids []Kid  `json:"kids"`
	}

	// Parent is the parent of some kids
	Parent struct {
		ID    int    `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	// Kid is a kid with an allowance
	Kid struct {
		ID      int      `json:"id"`
		Name    string   `json:"name"`
		Email   string   `json:"email"`
		Buckets []Bucket `json:"buckets"`
	}

	// Bucket is a category of allowance savings
	Bucket struct {
		ID                int     `json:"id"`
		Name              string  `json:"name"`
		DefaultAllocation int     `json:"defaultAllocation"`
		Total             float64 `json:"total"`
	}

	// Transaction is a single save or spend within a bucket
	Transaction struct {
		ID         int       `json:"id"`
		BucketID   int       `json:"bucketId"`
		ParentID   int       `json:"parentId"`
		Amount     float64   `json:"amount"`
		Note       string    `json:"note"`
		CreateDate time.Time `json:"createDate"`
	}
)
