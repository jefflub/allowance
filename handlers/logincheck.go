package handlers

// LoginCheck is the params for the login check
type LoginCheck struct {
	Token string `json:"token"`
}

// HandleRequest checks whether the token is currently logged in
func (l LoginCheck) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(l.Token); err != nil {
		return nil, err
	}

	return loginInfo, nil
}
