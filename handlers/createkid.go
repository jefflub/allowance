package handlers

import "github.com/jefflub/allowance/dbapi"

// CreateKid contains request parameters for creating a kid
type CreateKid struct {
	Token           string         `json:"token"`
	KidName         string         `json:"name" validate:"nonzero"`
	KidEmail        string         `json:"email" validate:"regexp=^[\\w+\\-.]+@[a-z\\d\\-.]+\\.[a-z]+$"`
	WeeklyAllowance float64        `json:"weeklyAllowance"`
	Buckets         []dbapi.Bucket `json:"buckets"`
}

type createKidResponse struct {
	KidID int `json:"kidId"`
}

// HandleRequest creates a kid, using default buckets or provided buckets
func (c CreateKid) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(c.Token); err != nil {
		return nil, err
	}

	kid, err := dbapi.CreateKid(loginInfo.FamilyID, c.KidName, c.KidEmail, c.WeeklyAllowance, c.Buckets)
	if err != nil {
		return nil, err
	}

	return createKidResponse{kid.ID}, nil
}
