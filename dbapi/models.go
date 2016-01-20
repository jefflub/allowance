package dbapi

import (
	"database/sql"
	"time"

	"github.com/jefflub/allowance/config"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("mysql", config.GetConfig().DbURL)
	if err != nil {
		panic(err)
	}
}

type (
	// Family represents the structure of our resource
	Family struct {
		ID   int    `json:"id"`
		Name string `json:"name"`
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
