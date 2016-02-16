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

	// Allocation is a bucket allocation
	Allocation struct {
		BucketID   int `json:"bucketId" validate:"nonzero"`
		Allocation int `json:"allocation" validate:"min=0"`
	}

	// Transaction is a single save or spend within a bucket
	Transaction struct {
		ID         int       `json:"id"`
		BucketID   int       `json:"bucketId"`
		BucketName string    `json:"bucketName"`
		ParentID   int       `json:"parentId"`
		ParentName string    `json:"parentName"`
		Amount     float64   `json:"amount"`
		Note       string    `json:"note"`
		CreateDate time.Time `json:"createDate"`
	}

	// LinkTokenInfo is information about a link token
	LinkTokenInfo struct {
		LinkToken string `json:"linkToken"`
		KidID     int    `json:"kidId"`
		KidName   string `json:"kidName"`
	}
)
