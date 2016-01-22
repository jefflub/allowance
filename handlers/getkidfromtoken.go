package handlers

import "github.com/jefflub/allowance/dbapi"

// GetKidFromToken is an empty struct, as params are in the URL
type GetKidFromToken struct {
}

type getKidFromTokenResponse struct {
	Kid          dbapi.Kid           `json:"kid"`
	Transactions []dbapi.Transaction `json:"transactions"`
}

// HandleRequest handles the user request
func (g GetKidFromToken) HandleRequest(vars map[string]string) (interface{}, error) {
	token := vars["token"]
	var response getKidFromTokenResponse
	var err error
	response.Kid, err = dbapi.GetKidFromToken(token)
	if err != nil {
		return nil, err
	}

	response.Transactions, err = dbapi.GetTransactions(response.Kid.ID, 0, 10)

	return response, err
}
