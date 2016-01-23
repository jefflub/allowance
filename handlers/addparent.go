package handlers

import "github.com/jefflub/allowance/dbapi"

// AddParent contains the request parameters
type AddParent struct {
	Token          string `json:"token" validate:"nonzero"`
	ParentName     string `json:"parentName" validate:"nonzero"`
	ParentEmail    string `json:"parentEmail" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	ParentPassword string `json:"parentPassword" validate:"min=6"`
}

type addParentResponse struct {
	ID int `json:"parentId"`
}

// HandleRequest handles the add parent request
func (a AddParent) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(a.Token); err != nil {
		return nil, err
	}

	var response addParentResponse
	response.ID, err = dbapi.AddParent(loginInfo.FamilyID, a.ParentName, a.ParentEmail, a.ParentPassword, nil)
	return response, err
}
