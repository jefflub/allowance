package handlers

type LoginCheck struct {
	Token string `json:"token"`
}

func (l LoginCheck) HandleRequest(vars map[string]string) (interface{}, error) {
	var loginInfo LoginTokenInfo
	var err error
	if loginInfo, err = ParseLoginToken(l.Token); err != nil {
		return nil, err
	}

	return loginInfo, nil
}
