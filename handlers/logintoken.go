package handlers

import (
	"encoding/json"
	"time"

	"github.com/dvsekhvalnov/jose2go"
)

// LoginTokenInfo contains current login information
type LoginTokenInfo struct {
	Expiration int64 `json:"exp"`
	FamilyID   int   `json:"fId"`
	ParentID   int   `json:"pId"`
}

var key = []byte{97, 48, 97, 50, 97, 98, 100, 56, 45, 54, 49, 54, 50, 45, 52, 49, 99, 51, 45, 56, 51, 100, 54, 45, 49, 99, 102, 53, 53, 57, 98, 52, 54, 97, 102, 99}

// CreateLoginToken creates a token to be used for login
func CreateLoginToken(familyID int, parentID int) (string, error) {
	exp := time.Now().Unix() + (3600 * 24)
	loginInfo := LoginTokenInfo{exp, familyID, parentID}

	payload, err := json.Marshal(loginInfo)
	if err != nil {
		panic(err)
	}
	token, err := jose.Sign(string(payload), jose.HS256, key)
	return string(token), err
}

// ParseLoginToken takes a login token and confirms that it's valid
func ParseLoginToken(token string) (LoginTokenInfo, error) {
	payload, _, err := jose.Decode(token, key)
	response := new(LoginTokenInfo)

	if err == nil {
		err = json.Unmarshal([]byte(payload), response)
		if err == nil {
			if response.Expiration < time.Now().Unix() {
				return *response, RequestError{"Login token is expired"}
			}
			return *response, nil
		}
	}

	return *response, err
}
