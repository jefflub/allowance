package handlers

import "github.com/jefflub/allowance/dbapi"

// CreateKid contains request parameters for creating a kid
type CreateKid struct {
	Token    string         `json:"token"`
	KidName  string         `json:"name" validate:"nonzero"`
	KidEmail string         `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Buckets  []dbapi.Bucket `json:"buckets"`
}

type createKidResponse struct {
	KidID int `json:"kidId"`
}

// HandleRequest creates a kid, using default buckets or provided buckets
func (c CreateKid) HandleRequest() (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(c.Token); err != nil {
		return nil, err
	}

	kid, err := dbapi.CreateKid(loginInfo.FamilyID, c.KidName, c.KidEmail, c.Buckets)
	if err != nil {
		return nil, err
	}

	return createKidResponse{kid.ID}, nil
}
