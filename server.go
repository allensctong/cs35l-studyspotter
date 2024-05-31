package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	_ "modernc.org/sqlite"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"studyspotter/src"
)

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

func GetUsersWrapper(db *sql.DB) gin.HandlerFunc {
	GetUsers := func (c *gin.Context) {
		var users = []User{}
		var username string
		var password string

		rows, err := db.Query("SELECT username, password FROM user")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&username, &password); err != nil {
				log.Fatal(err)
			}
			users = append(users, User{username, password})
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusOK, users)	
	}
	return GetUsers
}

func GetUserWrapper(db *sql.DB) gin.HandlerFunc {
	GetUser := func (c *gin.Context) {
		username := c.Param("username")
		var password string

		err := db.QueryRow("SELECT username, password FROM user WHERE username=?", username).Scan(&username, &password)
		
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, User{username, password})
	}

	return GetUser
}

func CreateUserWrapper(db *sql.DB) gin.HandlerFunc {
	 CreateUser := func (c *gin.Context) {
		//get username and password from request
		var user User
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		username := user.Username
		password := user.Password
		//query database for user
		err := db.QueryRow("SELECT username FROM user WHERE username=?", username).Scan(&username)
		//if user already exists
		if err != sql.ErrNoRows {
			if err != nil {
				panic(err)
			}

			c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists!"})
			return
		}
		//validate username and password length

		//hash password

		//create new user
		db.Exec("INSERT INTO user (username, password) VALUES (?, ?);", username, password)
		
		//send response back to client
		c.IndentedJSON(http.StatusCreated, User{username, password})
}

	return CreateUser
}

func main() {
	// Set up database.
	dbName := "studyspotter.db"
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		fmt.Printf("Unable to use data source: %s", err)
		return
	}
	defer db.Close()

	src.DbInit(db)
	token, _ := src.CreateToken("uwu")
	fmt.Printf("%s\n", token)

	// Set up router.
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
	}))
	router.GET("api/user", GetUsersWrapper(db))
	router.GET("api/user/:username", GetUserWrapper(db))
	router.POST("api/signup", CreateUserWrapper(db))
	router.Run("localhost:8080")

}
