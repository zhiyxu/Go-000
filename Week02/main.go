package Week02

import (
	"database/sql"
	"log"

	"github.com/pkg/errors"
)

// DAO
var db *sql.DB

func QueryRowById(id int) (string, error) {
	var username string
	// if id doesn't exist, the err would be sql.ErrNoRows
	if err := db.QueryRow("SELECT username FROM users WHERE id=?", id).Scan(&username);
		err != nil {
		return "", errors.Wrap(err, "failed to retrieve data")
	}

	return username, nil
}

//　business
func biz() {
	username, err := QueryRowById(5)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			// data not found logic
			return
		}
		log.Printf("%+v", err)
		return
	}
	log.Printf("username: %v", username)
}