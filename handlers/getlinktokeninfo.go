package handlers

import "github.com/jefflub/allowance/dbapi"

// GetLinkTokenInfo gets the link tokens for a family
type GetLinkTokenInfo struct {
	Token string `json:"token"`
}

type getLinkTokenInfoResponse struct {
	LinkTokens []dbapi.LinkTokenInfo `json:"linkTokens"`
}

// HandleRequest handles the link token get request
func (g GetLinkTokenInfo) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(g.Token); err != nil {
		return nil, err
	}
	var response getLinkTokenInfoResponse
	response.LinkTokens, err = dbapi.GetLinkTokens(loginInfo.FamilyID)
	return response, err
}
