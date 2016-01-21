package dbapi

import "errors"

// GetBuckets gets all of the buckets for a particular kid
func GetBuckets(familyID int, kidID int, checkFamily bool) ([]Bucket, error) {
	var buckets []Bucket

	if checkFamily {
		// Confirm kid belongs to family
		var f int
		row := db.QueryRow("SELECT familyId FROM kids WHERE kidId=?", kidID)
		err := row.Scan(&f)
		if err != nil {
			return nil, err
		}
		if f != familyID {
			return nil, errors.New("Invalid kid/family pair provided")
		}
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
