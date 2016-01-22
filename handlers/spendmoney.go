package handlers

import (
	"time"

	"github.com/jefflub/allowance/dbapi"
)

// SpendMoney contains the necessary params for spending money
type SpendMoney struct {
	Token    string  `json:"token" validate:"nonzero"`
	BucketID int     `json:"bucketId" validate:"nonzero"`
	Amount   float64 `json:"amount" validate:"min=0"`
	Note     string  `json:"note" validate:"nonzero"`
}

// HandleRequest handles the money spending request
func (s SpendMoney) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(s.Token); err != nil {
		return nil, err
	}

	// TODO: Validate bucket

	var transactions = []dbapi.Transaction{
		dbapi.Transaction{0, s.BucketID, loginInfo.ParentID, s.Amount * -1, s.Note, time.Now()},
	}

	return dbapi.AddTransactions(transactions)
}
