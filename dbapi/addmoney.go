package dbapi

import (
	"database/sql"
	"errors"
)

// AddMoney adds appropriate transactions for a given amount
func AddMoney(parentID int, kidID int, amount float64, note string, allocations []Allocation, tx *sql.Tx) ([]int, error) {
	var buckets []Bucket
	var err error
	if buckets, err = GetBuckets(kidID); err != nil {
		return nil, err
	}

	var transactions []Transaction

	// Validate allocations
	if allocations != nil && len(allocations) > 0 {
		transactions = make([]Transaction, len(allocations))
		var allSum = 0
		for i, all := range allocations {
			allSum += all.Allocation
			found := false
			for _, b := range buckets {
				if b.ID == all.BucketID {
					found = true
				}
			}
			if found == false {
				return nil, errors.New("Invalid bucket ID specified")
			}
			transactions[i].ParentID = parentID
			transactions[i].Note = note
			transactions[i].BucketID = all.BucketID
			transactions[i].Amount = (amount * float64(all.Allocation)) / 100
		}
		if allSum != 100 {
			return nil, errors.New("Invalid allocation amounts. Doesn't sum to 100")
		}
	} else {
		transactions = make([]Transaction, len(buckets))
		for i, b := range buckets {
			transactions[i].ParentID = parentID
			transactions[i].Note = note
			transactions[i].BucketID = b.ID
			transactions[i].Amount = (amount * float64(b.DefaultAllocation)) / 100
		}
	}
	// Add transactions to buckets
	return AddTransactions(transactions, tx)
}
