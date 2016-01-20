package handlers

import "github.com/jefflub/allowance/dbapi"

// Allocation is a bucket allocation
type Allocation struct {
	BucketID   int `json:"bucketId" validate:"nonzero"`
	Allocation int `json:"allocation" validate:"min=0"`
}

// AddMoney contains params for the addmoney call
type AddMoney struct {
	Token       string       `json:"token"`
	KidID       int          `json:"kidId" validate:"nonzero"`
	Amount      float64      `json:"amount" validate:"nonzero,min=0.01"`
	Note        string       `json:"note"`
	Allocations []Allocation `json:"allocations"`
}

// HandleRequest handles the addmoney request
func (a AddMoney) HandleRequest() (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(a.Token); err != nil {
		return nil, err
	}

	// Get buckets
	buckets, err := dbapi.GetBuckets(loginInfo.FamilyID, a.KidID)
	var transactions []dbapi.Transaction

	// Validate allocations
	if len(a.Allocations) > 0 {
		transactions = make([]dbapi.Transaction, len(a.Allocations))
		var allSum = 0
		for i, all := range a.Allocations {
			allSum += all.Allocation
			found := false
			for _, b := range buckets {
				if b.ID == all.BucketID {
					found = true
				}
			}
			if found == false {
				return nil, RequestError{"Invalid bucket ID specified"}
			}
			transactions[i].ParentID = loginInfo.ParentID
			transactions[i].Note = a.Note
			transactions[i].BucketID = all.BucketID
			transactions[i].Amount = (a.Amount * float64(all.Allocation)) / 100
		}
		if allSum != 100 {
			return nil, RequestError{"Invalid allocation amounts. Doesn't sum to 100"}
		}
	} else {
		transactions = make([]dbapi.Transaction, len(buckets))
		for i, b := range buckets {
			transactions[i].ParentID = loginInfo.ParentID
			transactions[i].Note = a.Note
			transactions[i].BucketID = b.ID
			transactions[i].Amount = (a.Amount * float64(b.DefaultAllocation)) / 100
		}
	}
	// Add transactions to buckets
	return dbapi.AddTransactions(transactions)
}
