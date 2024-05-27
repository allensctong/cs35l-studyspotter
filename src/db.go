package src

import (
	"database/sql"
)

func DbInit(db *sql.DB) {	
	if _, err := db.Exec(`
		DROP TABLE IF EXISTS user; 
		CREATE TABLE user(
			id VARCHAR(255), 
			password VARCHAR(255)
		);`);
	err != nil {
		panic(err)
	}
}

