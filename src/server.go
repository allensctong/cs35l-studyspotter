package src

import (
//	"fmt"
	"log"
	"net/http"
	"database/sql"

	"studyspotter/schemas"
	"github.com/gin-gonic/gin"
)

func GetUsersWrapper(db *sql.DB) gin.HandlerFunc {
	GetUsers := func (c *gin.Context) {
		var users = []schemas.User{}
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
			users = append(users, schemas.User{username, password})
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

		if user, hasUser := DBHasUser(db, username); hasUser {
			c.IndentedJSON(http.StatusOK, user)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		
	}

	return GetUser
}

func CreateUserWrapper(db *sql.DB) gin.HandlerFunc {
	 CreateUser := func (c *gin.Context) {
		//get username and password from request
		var user schemas.User
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		username := user.Username
		password := user.Password
		//query database for user
		if _, hasUser := DBHasUser(db, username); hasUser {
			//if user already exists
			c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists!"})
			return
		}
		//validate username and password length

		//hash password

		//create new user
		db.Exec("INSERT INTO user (username, password) VALUES (?, ?);", username, password)
		
		//send response back to client
		c.IndentedJSON(http.StatusCreated, gin.H{})
	}

	return CreateUser
}

func LoginWrapper(db *sql.DB) gin.HandlerFunc {
	Login := func (c *gin.Context) {
		var incomingUser schemas.User
		if err := c.BindJSON(&incomingUser); err != nil {
			panic(err)
		}
		username := incomingUser.Username
		password := incomingUser.Password

		if dbUser, hasUser := DBHasUser(db, username); hasUser {
			//login credentials are valid
			if dbUser.Password == password {
				tokenString, err := CreateToken(username)
				if err != nil {
					panic(err)
				}
				c.SetSameSite(http.SameSiteDefaultMode)
				c.SetCookie("Authorization", tokenString, 3600 * 24, "", "", false, true)
				c.JSON(http.StatusOK, gin.H{})
				return
			}
			//wrong password
			c.JSON(http.StatusBadRequest, gin.H{"message": "invalid password"})
		}

		//user not found
		c.JSON(http.StatusBadRequest, gin.H{"message": "user not found!"})

	}

	return Login
}

func AuthRequired(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err == http.ErrNoCookie {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	err = VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	c.Next()
}
