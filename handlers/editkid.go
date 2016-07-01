package handlers

import "github.com/jefflub/allowance/dbapi"

// EditKid models the kid editing request
type EditKid struct {
	Token           string  `json:"token"`
	KidID           int     `json:"kidId"`
	KidName         string  `json:"name" validate:"nonzero"`
	KidEmail        string  `json:"email" validate:"regexp=^[\\w+\\-.]+@[a-z\\d\\-.]+\\.[a-z]+$"`
	WeeklyAllowance float64 `json:"weeklyAllowance"`
}

// HandleRequest handles the edit kid request
func (e EditKid) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(e.Token); err != nil {
		return nil, err
	}

	kid, err := dbapi.EditKid(loginInfo.FamilyID, e.KidID, e.KidName, e.KidEmail, e.WeeklyAllowance)
	if err != nil {
		return nil, err
	}

	return kid, nil
}
