package handlers

import "github.com/jefflub/allowance/dbapi"

// CreateFamily defines the json structure of CreateFamily parameters
type CreateFamily struct {
	FamilyName     string `json:"familyName" validate:"nonzero"`
	ParentName     string `json:"parentName" validate:"nonzero"`
	ParentEmail    string `json:"parentEmail" validate:"regexp=^[\\w+\\-.]+@[a-z\\d\\-.]+\\.[a-z]+$"`
	ParentPassword string `json:"parentPassword" validate:"min=6"`
}

type createFamilyResponse struct {
	Token    string `json:"token"`
	FamilyID int    `json:"familyId"`
	ParentID int    `json:"parentId"`
}

// HandleRequest handles the CreateFamily request
func (params CreateFamily) HandleRequest(vars map[string]string) (interface{}, error) {
	var response createFamilyResponse
	var err error
	response.FamilyID, response.ParentID, err = dbapi.CreateFamily(params.FamilyName, params.ParentName, params.ParentEmail, params.ParentPassword)
	if err != nil {
		return response, err
	}

	if response.Token, err = CreateLoginToken(response.FamilyID, response.ParentID); err != nil {
		return response, err
	}

	return response, nil
}
