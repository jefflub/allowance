package handlers

import "github.com/jefflub/allowance/dbapi"

// AddMoney contains params for the addmoney call
type AddMoney struct {
	Token       string             `json:"token"`
	KidID       int                `json:"kidId" validate:"nonzero"`
	Amount      float64            `json:"amount" validate:"nonzero,min=0.01"`
	Note        string             `json:"note"`
	Allocations []dbapi.Allocation `json:"allocations"`
}

// HandleRequest handles the addmoney request
func (a AddMoney) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(a.Token); err != nil {
		return nil, err
	}

	if err = dbapi.CheckFamilyMembership(loginInfo.FamilyID, a.KidID); err != nil {
		return nil, err
	}

	// Get buckets
	return dbapi.AddMoney(loginInfo.ParentID, a.KidID, a.Amount, a.Note, a.Allocations, nil)
}
