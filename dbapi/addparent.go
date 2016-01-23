package dbapi

import (
	"database/sql"

	"golang.org/x/crypto/bcrypt"
)

// AddParent adds a parent. Uses passed transaction or a new one if one is not provided
func AddParent(familyID int, name string, email string, password string, tx *sql.Tx) (int, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	needCommit := false
	if tx == nil {
		tx, err = db.Begin()
		if err != nil {
			return 0, err
		}
		needCommit = true
		defer tx.Rollback()
	}

	// Create the parent
	if _, err := tx.Exec("INSERT INTO parents VALUES(NULL, ?, ?, ?, ?, NULL, NULL)", familyID, name, email, hashedPassword); err != nil {
		return 0, err
	}
	// Get the ID
	var ID int
	row := tx.QueryRow("SELECT LAST_INSERT_ID()")
	err = row.Scan(&ID)
	if needCommit {
		tx.Commit()
	}

	return ID, err
}
