package handlers

import "github.com/jefflub/allowance/dbapi"

// GetBuckets contains parameters for the getbuckets call
type GetBuckets struct {
	Token string `json:"token"`
	KidID int    `json:"kidId"`
}

type getBucketsResponse struct {
	Buckets []dbapi.Bucket `json:"buckets"`
}

// HandleRequest handles the getbuckets request
func (g GetBuckets) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(g.Token); err != nil {
		return nil, err
	}

	if err = dbapi.CheckFamilyMembership(loginInfo.FamilyID, g.KidID); err != nil {
		return nil, err
	}

	return dbapi.GetBuckets(g.KidID)
}
