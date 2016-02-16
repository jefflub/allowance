package dbapi

// GetBuckets gets all of the buckets for a particular kid
func GetBuckets(kidID int) ([]Bucket, error) {
	var buckets []Bucket

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
