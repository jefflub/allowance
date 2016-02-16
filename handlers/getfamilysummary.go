package handlers

import "github.com/jefflub/allowance/dbapi"

// GetFamilySummary holds parameters for the call
type GetFamilySummary struct {
	Token string `json:"token" validate:"nonzero"`
}

// HandleRequest handles the get family summary request
func (g GetFamilySummary) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	var family dbapi.Family
	if loginInfo, err = ParseLoginToken(g.Token); err != nil {
		return family, err
	}

	family, err = dbapi.GetFamily(loginInfo.FamilyID)
	if err != nil {
		return family, err
	}

	family.Kids, err = dbapi.GetKids(loginInfo.FamilyID)
	if err != nil {
		return family, err
	}

	for idx := range family.Kids {
		family.Kids[idx].Buckets, err = dbapi.GetBuckets(family.Kids[idx].ID)
		if err != nil {
			return family, err
		}
	}

	return family, err
}
