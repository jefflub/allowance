package handlers

import "github.com/jefflub/allowance/dbapi"

// Login defines parameters for login call
type Login struct {
	Email    string `json:"email" validate:"regexp=^[0-9a-z]+@[0-9a-z]+(\\.[0-9a-z]+)+$"`
	Password string `json:"password" validate:"min=6"`
}

type loginResponse struct {
	Token  string        `json:"token"`
	Family *dbapi.Family `json:"family"`
	Parent *dbapi.Parent `json:"parent"`
}

// HandleRequest handles the login request
func (l Login) HandleRequest(vars map[string]string) (interface{}, error) {
	parent, family, err := dbapi.Login(l.Email, l.Password)
	if err != nil {
		return nil, RequestError{err.Error()}
	}

	var token string
	if token, err = CreateLoginToken(family.ID, parent.ID); err != nil {
		return nil, err
	}

	return loginResponse{token, family, parent}, nil
}
