package handlers

import "github.com/jefflub/allowance/dbapi"

type DeleteLinkToken struct {
	Token     string `json:"token"`
	LinkToken string `json:"linkToken"`
}

type deleteLinkTokenResponse struct {
	LinkTokens []dbapi.LinkTokenInfo `json:"linkTokens"`
}

func (d DeleteLinkToken) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(d.Token); err != nil {
		return nil, err
	}

	if err = dbapi.DeleteLinkToken(loginInfo.FamilyID, d.LinkToken); err != nil {
		return nil, err
	}

	var response deleteLinkTokenResponse
	response.LinkTokens, err = dbapi.GetLinkTokens(loginInfo.FamilyID)
	return response, err
}
