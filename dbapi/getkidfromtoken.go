package dbapi

// GetKidFromToken gets a kid from a valid token
func GetKidFromToken(token string) (Kid, error) {
	var kid Kid

	row := db.QueryRow("SELECT kids.kidid, name, email FROM kids, kidtokens WHERE token=? AND kidtokens.kidid = kids.kidid", token)
	if err := row.Scan(&kid.ID, &kid.Name, &kid.Email); err != nil {
		return kid, err
	}

	buckets, err := GetBuckets(0, kid.ID, false)
	if err != nil {
		return kid, err
	}
	kid.Buckets = buckets
	return kid, nil
}
