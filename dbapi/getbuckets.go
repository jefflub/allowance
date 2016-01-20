package dbapi

import (
	"database/sql"
	"errors"

	"github.com/jefflub/allowance/config"
)

// GetBuckets gets all of the buckets for a particular kid
func GetBuckets(familyID int, kidID int) ([]Bucket, error) {
	buckets := make([]Bucket, 0, 5)
	db, err := sql.Open("mysql", config.GetConfig().DbURL)
	if err != nil {
		return nil, err
	}
	defer db.Close()

	// Confirm kid belongs to family
	var f int
	row := db.QueryRow("SELECT familyId FROM kids WHERE kidId=?", kidID)
	err = row.Scan(&f)
	if err != nil {
		return nil, err
	}
	if f != familyID {
		return nil, errors.New("Invalid kid/family pair provided")
	}

	rows, err := db.Query("SELECT BucketID, Name, DefaultAllocation, CurrentTotal FROM buckets WHERE kidid=?", kidID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var bucket Bucket
		if err := rows.Scan(&bucket.ID, &bucket.Name, &bucket.DefaultAllocation, &bucket.Total); err != nil {
			return nil, err
		}
		buckets = append(buckets, bucket)
	}

	return buckets, nil
}
