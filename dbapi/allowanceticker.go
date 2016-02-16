package dbapi

import (
	"log"
	"time"

	"github.com/jinzhu/now"
)

func addWeeklyAllowance(kidID int, allowance float64) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()
	// in transaction:
	// Add the money
	log.Printf("Adding allowance %v for kid %v", allowance, kidID)
	_, err = AddMoney(0, kidID, allowance, "Allowance", nil, tx)
	if err != nil {
		return err
	}
	nextAllowance := now.New(time.Now().AddDate(0, 0, 1)).Sunday()
	log.Printf("Setting next allowance to %v", nextAllowance)
	_, err = tx.Exec("UPDATE kids SET nextallowance=? WHERE kidid=?", nextAllowance, kidID)
	if err != nil {
		return err
	}
	tx.Commit()
	return nil
}

func init() {
	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for t := range ticker.C {
			log.Println("Tick at", t)
			// Get kids who need Allowances
			if rows, err := db.Query("SELECT kidid, weeklyallowance FROM kids WHERE weeklyallowance > 0 AND nextallowance <= curdate()"); err != nil {
				log.Println("Error retrieving allowance updates", err.Error())
			} else {
				for rows.Next() {
					var kidID int
					var allowance float64
					if err = rows.Scan(&kidID, &allowance); err != nil {
						log.Println("Error scanning allowance row", err.Error())
					} else {
						err = addWeeklyAllowance(kidID, allowance)
						if err != nil {
							log.Println("Error adding allowance", err.Error())
						}
					}
				}
			}
		}
	}()
}
