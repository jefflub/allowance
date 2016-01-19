package handlers

import "github.com/jefflub/allowance/dbapi"

// CreateFamily defines the json structure of CreateFamily parameters
type CreateFamily struct {
	FamilyName     string `json:"familyName" validate:"nonzero"`
	ParentName     string `json:"parentName" validate:"nonzero"`
	ParentEmail    string `json:"parentEmail" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	ParentPassword string `json:"parentPassword" validate:"min=6"`
}

type createFamilyResponse struct {
	Token    string `json:"token"`
	FamilyID int    `json:"familyId"`
	ParentID int    `json:"parentId"`
}

// HandleRequest handles the CreateFamily request
func (params CreateFamily) HandleRequest() (interface{}, error) {
	family, parent, err := dbapi.CreateFamily(params.FamilyName, params.ParentName, params.ParentEmail, params.ParentPassword)
	if err != nil {
		return nil, err
	}

	var token string
	if token, err = CreateLoginToken(family.ID, parent.ID); err != nil {
		return nil, err
	}

	// Respond
	response := createFamilyResponse{token, family.ID, parent.ID}
	return response, nil
}
