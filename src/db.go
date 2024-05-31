package src

import (
	"database/sql"
	"studyspotter/schemas"
)

func DbInit(db *sql.DB) {	
	if _, err := db.Exec(`
		DROP TABLE IF EXISTS user; 
		CREATE TABLE user(
			username VARCHAR(255), 
			password VARCHAR(255)
		);`);
	err != nil {
		panic(err)
	}
}

func DBHasUser(db *sql.DB, username string) (schemas.User, bool) {
	var password string
	err := db.QueryRow("SELECT username, password FROM user WHERE username=?", username).Scan(&username, &password)

	if err == sql.ErrNoRows {
		return schemas.User{"", ""}, false
	}

	if err != nil {
		panic(err)
	}

	return schemas.User{username, password}, true
}
