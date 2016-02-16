package dbapi

import "database/sql"

// GetTransactions returns count transactions for a kid in descending date order,
// starting with offset
func GetTransactions(kidID int, offset int, count int) ([]Transaction, error) {
	var query = `SELECT transactionId, transactions.bucketId, buckets.Name, createparentid, amount, note, transactions.createdate
               FROM transactions, buckets
               WHERE buckets.kidid=? AND transactions.bucketid = buckets.bucketid
               ORDER BY transactions.createdate desc
               LIMIT ?, ?`

	var trans []Transaction

	rows, err := db.Query(query, kidID, offset, count)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var t Transaction
		var pid sql.NullInt64
		err = rows.Scan(&t.ID, &t.BucketID, &t.BucketName, &pid, &t.Amount, &t.Note, &t.CreateDate)
		if err != nil {
			return nil, err
		}
		if pid.Valid {
			t.ParentID = int(pid.Int64)
		}
		trans = append(trans, t)
	}
	return trans, nil
}
