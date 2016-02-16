package dbapi

import "errors"

// CheckFamilyMembership checks if the given kid is a member of the family
func CheckFamilyMembership(familyID int, kidID int) error {
	// Confirm kid belongs to family
	var f int
	row := db.QueryRow("SELECT familyId FROM kids WHERE kidId=?", kidID)
	err := row.Scan(&f)
	if err != nil {
		return err
	}
	if f != familyID {
		return errors.New("Invalid kid/family pair provided")
	}
	return nil
}
