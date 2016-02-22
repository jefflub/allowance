package dbapi

// GetKids gets all of the kids for a family
func GetKids(familyID int) ([]Kid, error) {
	var kids []Kid

	rows, err := db.Query("SELECT kidid, name, email, weeklyAllowance FROM kids WHERE familyid=?", familyID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var kid Kid
		if err = rows.Scan(&kid.ID, &kid.Name, &kid.Email, &kid.WeeklyAllowance); err != nil {
			return nil, err
		}
		kids = append(kids, kid)
	}
	return kids, err
}
