package dbapi

// CreateFamily creates a family and the first parent
func CreateFamily(familyName string, parentName string, parentEmail string, parentPassword string) (int, int, error) {
	var familyID int
	var parentID int

	tx, err := db.Begin()
	if err != nil {
		return familyID, parentID, err
	}
	defer tx.Rollback()

	// Create the family
	if _, err := tx.Exec("INSERT INTO family VALUES(NULL, ?, NULL, NULL)", familyName); err != nil {
		return familyID, parentID, err
	}
	// Get the ID
	row := tx.QueryRow("SELECT LAST_INSERT_ID()")
	if err := row.Scan(&familyID); err != nil {
		return familyID, parentID, err
	}

	parentID, err = AddParent(familyID, parentName, parentEmail, parentPassword, tx)
	tx.Commit()
	return familyID, parentID, nil
}
