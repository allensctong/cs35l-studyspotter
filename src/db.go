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
			password VARCHAR(255),
			bio TEXT,
			following INT,
			followers INT
		);
		DROP TABLE IF EXISTS post; 
		CREATE TABLE post(
			username VARCHAR(255),
			caption TEXT,
			uploadtime DATETIME
		);`);
	err != nil {
		panic(err)
	}
}

func DBHasUser(db *sql.DB, username string) bool {
	err := db.QueryRow("SELECT username FROM user WHERE username=?", username).Scan(&username)
	if err == sql.ErrNoRows {
		return false
	}

	if err != nil {
		panic(err)
	}

	return true
}

func DBGetPasswordHash(db *sql.DB, username string) string {
	var password string
	err := db.QueryRow("SELECT password FROM user WHERE username=?", username).Scan(&password)

	if err != nil {
		panic(err)
	}

	return password
}

func DBGetUserProfile(db *sql.DB, username string) schemas.UserProfile {
	var user schemas.UserProfile
	err := db.QueryRow("SELECT username, bio, following, followers FROM user WHERE username=?", username).Scan(&user.Username, &user.Bio, &user.Following, &user.Followers)

	if err != nil {
		panic(err)
	}

	return user
	
}
