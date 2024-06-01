package src

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	"studyspotter/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func GetUsersWrapper(db *sql.DB) gin.HandlerFunc {
	GetUsers := func (c *gin.Context) {
		var users = []schemas.Login{}
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
			users = append(users, schemas.Login{username, password})
		}

		if err := rows.Err(); err != nil {
			log.Fatal(err)
		}
		c.IndentedJSON(http.StatusOK, users)	
	}
	return GetUsers
}

//Get User Profile (GET FOR USER PAGE)
func GetUserWrapper(db *sql.DB) gin.HandlerFunc {
	GetUser := func (c *gin.Context) {
		var user schemas.UserProfile
		username := c.Param("username")

		if hasUser := DBHasUser(db, username); hasUser {
			user = DBGetUserProfile(db, username)
			c.IndentedJSON(http.StatusOK, user)
			return
		}

		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
		
	}

	return GetUser
}

//Create User Profile
func CreateUserWrapper(db *sql.DB) gin.HandlerFunc {
	 CreateUser := func (c *gin.Context) {
		//get username and password from request
		var user schemas.Login
		if err := c.BindJSON(&user); err != nil {
			panic(err)
		}
		username := user.Username
		password := user.Password
		//query database for user
		if hasUser := DBHasUser(db, username); hasUser {
			//if user already exists
			c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists!"})
			return
		}

		//hash password
		hashBytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
		if err != nil {
			return
		}

		passwordHash := string(hashBytes)

		//create new user
		db.Exec(`INSERT INTO user (username, password, following, followers, bio) VALUES (?, ?, 0, 0, "");`, username, passwordHash)
		
		//send response back to client
		c.IndentedJSON(http.StatusCreated, gin.H{})
	}

	return CreateUser
}

func LoginWrapper(db *sql.DB) gin.HandlerFunc {
	Login := func (c *gin.Context) {
		var incomingUser schemas.Login
		if err := c.BindJSON(&incomingUser); err != nil {
			panic(err)
		}
		username := incomingUser.Username
		password := incomingUser.Password

		if hasUser := DBHasUser(db, username); hasUser {
			//check that login credentials are valid
			
			if CheckPasswordHash(password, DBGetPasswordHash(db, username)) {
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

/* ---------------------------HELPER FUNCTIONS--------------------------- */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func AuthRequired(c *gin.Context) {
	tokenString, err := c.Cookie("Authorization")

	if err == http.ErrNoCookie {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}

	tokenString, err = VerifyToken(tokenString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{})
		return
	}
	fmt.Printf("%s\n", tokenString)

	c.Next()
}
