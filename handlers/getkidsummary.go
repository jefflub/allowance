package handlers

import "github.com/jefflub/allowance/dbapi"

// GetKidSummary holds parameters for the get kid summary call
type GetKidSummary struct {
	Token string `json:"token" validate:"nonzero"`
	KidID int    `json:"kidId" validate:"nonzero"`
}

type getKidSummaryResponse struct {
	Kid          dbapi.Kid           `json:"kid"`
	Transactions []dbapi.Transaction `json:"transactions"`
}

// HandleRequest handles the get kid summary request
func (g GetKidSummary) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(g.Token); err != nil {
		return nil, err
	}
	var response getKidSummaryResponse
	response.Kid, err = dbapi.GetKid(loginInfo.FamilyID, g.KidID)
	if err != nil {
		return nil, err
	}

	response.Transactions, err = dbapi.GetTransactions(g.KidID, 0, 10)

	return response, err
}
