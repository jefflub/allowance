package dbapi

// GetKid gets a single kid and their buckets
func GetKid(familyID int, kidID int) (Kid, error) {
	var kid Kid

	row := db.QueryRow("SELECT kidid, name, email FROM kids WHERE familyid=? AND kidid=?", familyID, kidID)
	if err := row.Scan(&kid.ID, &kid.Name, &kid.Email); err != nil {
		return kid, err
	}

	buckets, err := GetBuckets(kidID)
	if err != nil {
		return kid, err
	}
	kid.Buckets = buckets
	return kid, nil
}
