package handlers

import (
	"log"

	"github.com/jefflub/allowance/dbapi"
)

// GetTransactions holds the parameters for the get transactions call
type GetTransactions struct {
	Token  string `json:"token"`
	KidID  int    `json:"kidId"`
	Offset int    `json:"offset"`
	Count  int    `json:"count"`
}

type getTransactionsResponse struct {
	HasNextPage  bool                `json:"hasNextPage"`
	Transactions []dbapi.Transaction `json:"transactions"`
}

// HandleRequest handles the get transactions request
func (g GetTransactions) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(g.Token); err != nil {
		return nil, err
	}

	if err = dbapi.CheckFamilyMembership(loginInfo.FamilyID, g.KidID); err != nil {
		return nil, err
	}
	log.Println("Offset, Count: ", g.Offset, g.Count)
	var response getTransactionsResponse
	var trans []dbapi.Transaction
	if trans, err = dbapi.GetTransactions(g.KidID, g.Offset, g.Count+1); err != nil {
		return nil, err
	}

	log.Println("Transactions: ", len(trans))
	if len(trans) == g.Count+1 {
		response.HasNextPage = true
		response.Transactions = trans[0:g.Count]
	} else {
		response.HasNextPage = false
		response.Transactions = trans
	}
	if response.Transactions == nil {
		log.Println("Transactions nil")
	} else {
		log.Println("Transactions length: ", len(response.Transactions))
	}

	return response, nil
}
