package dbapi

var linkTokenQuery = `select token, kids.kidid, name
                      from kidtokens, kids
                      where kids.familyid=? and kidtokens.KidID = kids.KidID`

// GetLinkTokens returns the list of link tokens for a family
func GetLinkTokens(familyID int) ([]LinkTokenInfo, error) {
	var tokenInfos []LinkTokenInfo

	rows, err := db.Query(linkTokenQuery, familyID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var ti LinkTokenInfo
		err = rows.Scan(&ti.LinkToken, &ti.KidID, &ti.KidName)
		if err != nil {
			return nil, err
		}
		tokenInfos = append(tokenInfos, ti)
	}
	return tokenInfos, nil
}
