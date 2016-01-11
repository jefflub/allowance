package dbapi

import "database/sql"

// CreateFamily creates a family and the first parent
func CreateFamily(familyName string, parentName string, parentEmail string, parentPassword string) (Family, Parent, error) {
	var family Family
	var parent Parent

	db, err := sql.Open("mysql", "allowance_user:goniff@/allowance")
	if err != nil {
		return family, parent, err
	}
	defer db.Close()

	tx, err := db.Begin()
	if err != nil {
		return family, parent, err
	}
	defer tx.Rollback()

	// Create the family
	if _, err := tx.Exec("INSERT INTO family VALUES(NULL, ?, NULL, NULL)", familyName); err != nil {
		return family, parent, err
	}
	// Get the ID
	row := tx.QueryRow("SELECT LAST_INSERT_ID()")
	if err := row.Scan(&family.ID); err != nil {
		return family, parent, err
	}
	// Create the parent
	if _, err := tx.Exec("INSERT INTO parents VALUES(NULL, ?, ?, ?, NULL, NULL)", family.ID, parentName, parentEmail); err != nil {
		return family, parent, err
	}
	// Get the ID
	row = tx.QueryRow("SELECT LAST_INSERT_ID()")
	if err := row.Scan(&parent.ID); err != nil {
		return family, parent, err
	}
	tx.Commit()
	return family, parent, nil
}
