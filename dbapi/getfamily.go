package dbapi

// GetFamily returns a family object from the DB
func GetFamily(familyID int) (Family, error) {
	var family Family
	// Get the family by ID
	row := db.QueryRow("SELECT name FROM family WHERE familyid=?", familyID)
	err := row.Scan(&family.Name)
	return family, err
}
