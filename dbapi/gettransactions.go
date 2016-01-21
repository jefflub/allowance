package dbapi

var query = `SELECT transactionId, transactions.bucketId, createparentid, amount, note, transactions.createdate
             FROM transactions, buckets
             WHERE buckets.kidid=? AND transactions.bucketid = buckets.bucketid
             ORDER BY transactions.createdate desc
             LIMIT ?, ?`

// GetTransactions returns count transactions for a kid in descending date order,
// starting with offset
func GetTransactions(kidID int, offset int, count int) ([]Transaction, error) {
	var trans []Transaction

	rows, err := db.Query(query, kidID, offset, count)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Transaction
		err = rows.Scan(&t.ID, &t.BucketID, &t.ParentID, &t.Amount, &t.Note, &t.CreateDate)
		if err != nil {
			return nil, err
		}
		trans = append(trans, t)
	}
	return trans, nil
}
