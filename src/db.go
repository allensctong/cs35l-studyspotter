package src

import (
	"fmt"
	"strconv"
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
			id INT NOT NULL UNIQUE,
			username VARCHAR(255) NOT NULL,
			imagepath TEXT NOT NULL UNIQUE,
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
	//get info from main usertable
	err := db.QueryRow("SELECT username, bio, pfp FROM user WHERE username=?", username).Scan(&user.Username, &user.Bio, &user.ProfilePicture)
	if err != nil {
		panic(err)
	}
	
	//get follower/following counts (TODO MAKE THIS A HELPER FUNC AND USE GOROUTINE)
	err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM following%s", user.Username)).Scan(&user.FollowingCount)
	followingRows, err := db.Query(fmt.Sprintf("SELECT username FROM following%s", user.Username))
	defer followingRows.Close()
	user.Following = []string{}
	for followingRows.Next() {
		var followee string
		followingRows.Scan(&followee)
		user.Following = append(user.Following, followee)
	}
	if followingRows.Err(); err != nil{
		panic(err)
	}

	err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM followers%s", user.Username)).Scan(&user.FollowersCount)
	rows, err := db.Query(fmt.Sprintf("SELECT username FROM followers%s", user.Username))
	defer rows.Close()
	user.Followers = []string{}
	for rows.Next() {
		var follower string
		rows.Scan(&follower)
		user.Followers = append(user.Followers, follower)
	}
	if rows.Err(); err != nil{
		panic(err)
	}

	//get images for posts
	user.Posts = []string{}
	imageRows, err := db.Query("SELECT imagepath FROM post WHERE username=? ORDER BY uploadtime DESC", user.Username)
	defer imageRows.Close()
	for imageRows.Next() {
		var post string
		imageRows.Scan(&post)
		user.Posts = append(user.Posts, post)
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

	_, err = db.Exec(fmt.Sprintf("CREATE TABLE following%s (username VARCHAR(255));", user.Username))
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE followers%s (username VARCHAR(255));", user.Username))
	if err != nil {
		fmt.Println(err)
		return false
	}

	return true
}

func DBCreatePost(db *sql.DB, post schemas.Post) bool {
	_, err := db.Exec(`INSERT INTO post (id, username, imagepath, caption) VALUES (?, ?, ?, ?);`, post.ID, post.Username, post.ImagePath, post.Caption)
	if err != nil {
		panic(err)
		return false
	}

	_, err = db.Exec(fmt.Sprintf("CREATE TABLE comment%s (username VARCHAR(255) NOT NULL, comment TEXT DEFAULT '', commenttime TIMESTAMP DEFAULT CURRENT_);", strconv.Itoa(post.ID)))
	_, err = db.Exec(fmt.Sprintf("CREATE TABLE likes%s (username VARCHAR(255));", strconv.Itoa(post.ID)))
	if err != nil {
		panic(err)
		return false
	}

	return true
}

func DBGetPost(db *sql.DB, id int) schemas.Post {
	var post schemas.Post
	post.ID = id
	//get info from main posttable
	err := db.QueryRow("SELECT username, imagepath, caption FROM post WHERE id=?;", id).Scan(&post.Username, &post.ImagePath, &post.Caption)
	if err != nil {
		panic(err)
	}

	
	err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM likes%d;", post.ID)).Scan(&post.LikeCount)
	likeRows, err := db.Query(fmt.Sprintf("SELECT username FROM likes%d", post.ID))
	if err != nil {
		panic(err)
	}
	defer likeRows.Close()
	post.LikedUserList = []string{}
	for likeRows.Next() {
		var liker string
		likeRows.Scan(&liker)
		post.LikedUserList = append(post.LikedUserList, liker)
	}
	if likeRows.Err(); err != nil{
		panic(err)
	}

	rows, err := db.Query(fmt.Sprintf("SELECT username, comment FROM comment%d ORDER BY commenttime DESC", post.ID))
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	post.Comments = []schemas.Comment{}
	for rows.Next() {
		var comment schemas.Comment
		rows.Scan(&comment.Username, &comment.Text)
		post.Comments = append(post.Comments, comment)
	}
	if rows.Err(); err != nil{
		panic(err)
	}

	return post
}
