package dbapi

// EditKid edits the data for a kid
func EditKid(familyID int, kidID int, name string, email string, weeklyAllowance float64) (Kid, error) {
	kid := Kid{0, name, email, weeklyAllowance, nil}
	_, err := db.Exec("UPDATE kids SET name=?, email=?, weeklyallowance=? WHERE familyid=? AND kidid=?", name, email, weeklyAllowance, familyID, kidID)
	return kid, err
}
