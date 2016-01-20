package dbapi

import (
	"database/sql"
	"log"

	"github.com/jefflub/allowance/config"
	"github.com/nu7hatch/gouuid"
)

// AddTransactions adds one or more transactions to the appropriate bucket
func AddTransactions(transactions []Transaction) ([]int, error) {
	transIDs := make([]int, len(transactions))
	db, err := sql.Open("mysql", config.GetConfig().DbURL)
	if err != nil {
		return transIDs, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return transIDs, err
	}
	defer tx.Rollback()
	u4, err := uuid.NewV4()
	for i, t := range transactions {
		if t.Amount == 0 {
			continue
		}
		log.Printf("Bucket: %v, Parent: %v, Amount: %v, Note: %s, UUID: %s", t.BucketID, t.ParentID, t.Amount, t.Note, u4.String())
		if _, err := tx.Exec("INSERT INTO transactions VALUES(NULL, ?, ?, ?, ?, ?, NULL, NULL)", t.BucketID, t.ParentID, t.Amount, t.Note, u4.String()); err != nil {
			return transIDs, err
		}
		// Get the ID
		row := tx.QueryRow("SELECT LAST_INSERT_ID()")
		if err := row.Scan(&transIDs[i]); err != nil {
			return transIDs, err
		}
	}
	tx.Commit()
	return transIDs, nil
}
