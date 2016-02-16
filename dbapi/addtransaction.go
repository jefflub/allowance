package dbapi

import (
	"database/sql"

	"github.com/nu7hatch/gouuid"
)

// AddTransactions adds one or more transactions to the appropriate bucket
func AddTransactions(transactions []Transaction, tx *sql.Tx) ([]int, error) {
	transIDs := make([]int, len(transactions))
	var err error
	needCommit := false
	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return transIDs, err
		}
		needCommit = true
		defer tx.Rollback()
	}

	u4, err := uuid.NewV4()
	for i, t := range transactions {
		if t.Amount == 0 {
			continue
		}
		//log.Printf("Bucket: %v, Parent: %v, Amount: %v, Note: %s, UUID: %s", t.BucketID, t.ParentID, t.Amount, t.Note, u4.String())
		if t.ParentID == 0 {
			if _, err := tx.Exec("INSERT INTO transactions VALUES(NULL, ?, NULL, ?, ?, ?, NULL, NULL)", t.BucketID, t.Amount, t.Note, u4.String()); err != nil {
				return transIDs, err
			}
		} else {
			if _, err := tx.Exec("INSERT INTO transactions VALUES(NULL, ?, ?, ?, ?, ?, NULL, NULL)", t.BucketID, t.ParentID, t.Amount, t.Note, u4.String()); err != nil {
				return transIDs, err
			}
		}
		// Get the ID
		row := tx.QueryRow("SELECT LAST_INSERT_ID()")
		if err := row.Scan(&transIDs[i]); err != nil {
			return transIDs, err
		}
	}
	if needCommit {
		tx.Commit()
	}

	return transIDs, nil
}
