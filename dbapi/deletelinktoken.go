package dbapi

// DeleteLinkToken deletes a link token if the provided family is the owner
func DeleteLinkToken(familyID int, linkToken string) error {
	var token string
	// Confirm family ownership
	row := db.QueryRow("SELECT token FROM kidtokens, kids WHERE token=? AND kidtokens.kidid=kids.kidid AND kids.FamilyID=?", linkToken, familyID)
	if err := row.Scan(&token); err != nil {
		return err
	}

	// Delete the token
	_, err := db.Exec("DELETE FROM kidtokens WHERE token=?", linkToken)
	return err
}
