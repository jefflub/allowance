package dbapi

import (
	"database/sql"

	"github.com/jefflub/allowance/config"

	"golang.org/x/crypto/bcrypt"
)

// Login checks the password for a user and returns a login object if correct
func Login(email string, password string) (*Parent, *Family, error) {
	db, err := sql.Open("mysql", config.GetConfig().DbURL)
	if err != nil {
		return nil, nil, err
	}
	defer db.Close()

	var parent = Parent{}
	var family = Family{}
	passwordHash := make([]byte, 60)
	row := db.QueryRow("SELECT parents.familyid, ParentId, parents.name, family.name, PasswordHash FROM parents, family WHERE Email=? and parents.familyid=family.familyid", email)
	if err := row.Scan(&family.ID, &parent.ID, &parent.Name, &family.Name, &passwordHash); err != nil {
		return nil, nil, err
	}
	if err := bcrypt.CompareHashAndPassword(passwordHash, []byte(password)); err != nil {
		return nil, nil, err
	}
	parent.Email = email
	return &parent, &family, nil
}
