package handlers

import "github.com/jefflub/allowance/dbapi"

// CreateKidToken holds parameters for the create kid token call
type CreateKidToken struct {
	Token string `json:"token" validate:"nonzero"`
	KidID int    `json:"kidId"`
}

type createKidTokenResponse struct {
	KidToken string `json:"kidToken"`
}

// HandleRequest handles the create kid token request
func (c CreateKidToken) HandleRequest(vars map[string]string) (interface{}, error) {
	var err error
	if _, err = ParseLoginToken(c.Token); err != nil {
		return nil, err
	}

	var response createKidTokenResponse
	response.KidToken, err = dbapi.CreateKidToken(c.KidID)

	return response, err
}
