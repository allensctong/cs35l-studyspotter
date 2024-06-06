package src

import (
	"fmt"
	"database/sql"
	"studyspotter/schemas"
	"golang.org/x/crypto/bcrypt"
)

func DbInit(db *sql.DB) {	
	if _, err := db.Exec(`
		DROP TABLE IF EXISTS user; 
		CREATE TABLE user(
			username VARCHAR(255) NOT NULL UNIQUE, 
			password VARCHAR(255) NOT NULL,
			pfp TEXT DEFAULT 'http://localhost:8080/assets/default-pfp.jpg',
			bio TEXT DEFAULT ''
		);
		DROP TABLE IF EXISTS post; 
		CREATE TABLE post(
			username VARCHAR(255) NOT NULL,
			image TEXT NOT NULL UNIQUE,
			caption TEXT DEFAULT '',
			uploadtime TIMESTAMP DEFAULT CURRENT_TIMESTAMP
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
	err := db.QueryRow("SELECT username, bio, following_count, followers_count, pfp FROM user WHERE username=?", username).Scan(&user.Username, &user.Bio, &user.FollowingCount, &user.FollowersCount, &user.ProfilePicture)

	if err != nil {
		panic(err)
	}

	return user
	
}

func DBCreateUserProfile(db *sql.DB, user schemas.Login) bool {
	//hash password
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(user.Password), 12)
	if err != nil {
		panic(err)
	}
	passwordHash := string(hashBytes)

	//inser user in db
	_, err = db.Exec(`INSERT INTO user (username, password) VALUES (?, ?);`, user.Username, passwordHash)
	if err != nil {
		fmt.Println(err)
		return false
	}

	_, err = db.Exec(fmt.Sprintf("CREATE TABLE %sfollowing (username VARCHAR(255));", user.Username))
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE %sfollowers (username VARCHAR(255));", user.Username))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}
