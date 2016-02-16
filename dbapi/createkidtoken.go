package dbapi

import (
	"math/rand"
	"time"
)

const validChars = "BCDFGHJKLMNPRSTWXZ23456789"
const tokenLength = 8

func init() {
	rand.Seed(time.Now().Unix())
}

func makeTokenString() string {
	var bytes []byte
	for i := 0; i < tokenLength; i++ {
		bytes = append(bytes, validChars[rand.Intn(len(validChars))])
	}
	return string(bytes)
}

// CreateKidToken creates a token that allows retrieval of kid information
func CreateKidToken(kidID int) (string, error) {
	token := makeTokenString()

	_, err := db.Exec("INSERT INTO kidtokens VALUES(?,?,NULL)", token, kidID)

	return token, err
}
