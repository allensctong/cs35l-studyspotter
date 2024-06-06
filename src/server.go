package src

import (
	"fmt"
	"strconv"
	"net/http"
	"database/sql"
	"path/filepath"

	"studyspotter/schemas"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

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

		//try to add user to database
		if madeUser := DBCreateUserProfile(db, user); madeUser {
			//if user successfully created
			c.JSON(http.StatusCreated, gin.H{})
			return
		}
		
		//send response back to client
		c.JSON(http.StatusBadRequest, gin.H{"message": "user already exists!"})
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
				c.SetCookie("Authorization", tokenString, 3600 * 24, "", "", false, false)
				c.SetCookie("Username", username, 3600 * 24, "", "", false, false)
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

func CreatePostWrapper(db *sql.DB) gin.HandlerFunc {
	createPost := func (c *gin.Context) {
		var post schemas.Post
		image, err := c.FormFile("image")
		if err != nil {
			panic(err)
		}
		post.Caption, _ := c.GetPostForm("caption")
		post.Username, _ := c.GetPostForm("username")
		//get current number of posts and save image to image directory
		var postCount int
		err = db.QueryRow("SELECT COUNT(*) FROM post").Scan(&postCount)
		imagePath := "assets/" + strconv.Itoa(postCount) + filepath.Ext(file.Filename)
		err = c.SaveUploadedFile(file, imagePath)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with upload"})
		}

		c.JSON(http.StatusOK, gin.H{})
	}
	return createPost
}

func CORSMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:5173")
	c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
        c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
        c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")
        if c.Request.Method == "OPTIONS" {
            c.AbortWithStatus(204)
            return
        }

        c.Next()
    }
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
