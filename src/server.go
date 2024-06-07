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

var LocalAssetsPath string = "assets/" 
var HostAddress string = "http://localhost:8080/"

//Get User Profile (GET FOR USER PAGE)
func GetUserWrapper(db *sql.DB) gin.HandlerFunc {
	GetUser := func (c *gin.Context) {
		var user schemas.UserProfile
		username := c.Param("username")
		ownName, _ := c.Cookie("Username")

		if hasUser := DBHasUser(db, username); hasUser {
			user = DBGetUserProfile(db, username)
			var following bool
			var u string
			err := db.QueryRow(fmt.Sprintf("SELECT username FROM followers%s WHERE username=?", username), ownName).Scan(&u)
			if err == sql.ErrNoRows{
				following = false
			} else {
				following = true
			}
			user.IsFriend = following
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
		post.Caption, _ = c.GetPostForm("caption")
		post.Username, _ = c.GetPostForm("username")
		//get current number of posts and save image to image directory
		var postCount int
		err = db.QueryRow("SELECT COUNT(*) FROM post").Scan(&postCount)
		imagePath := LocalAssetsPath + strconv.Itoa(postCount) + filepath.Ext(image.Filename)
		err = c.SaveUploadedFile(image, imagePath)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with upload"})
		}
		post.ImagePath = HostAddress + imagePath
		post.ID = postCount

		//create Post
		DBCreatePost(db, post)

		c.JSON(http.StatusOK, gin.H{})
	}
	return createPost
}

func ChangePfpWrapper(db *sql.DB) gin.HandlerFunc {
	changePfp := func (c *gin.Context) {
		username, _ := c.GetPostForm("username")
		newPfp, err := c.FormFile("image")
		if err != nil {
			panic(err)
		}
		pfpPath := LocalAssetsPath + "pfp" + username + filepath.Ext(newPfp.Filename)
		err = c.SaveUploadedFile(newPfp, pfpPath)
		_, err = db.Exec("UPDATE user SET pfp=? WHERE username=?;", HostAddress + pfpPath, username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with upload"})
		}
		
		c.JSON(http.StatusOK, gin.H{})
	}	

	return changePfp
}

func ChangeBioWrapper(db *sql.DB) gin.HandlerFunc {
	changeBio := func (c *gin.Context) {
		type bioBody struct {
			NewBio string `json:"bio"`
		}
		var inc bioBody
		username := c.Param("username")
		err := c.BindJSON(&inc)
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with change"})
		}
		_, err = db.Exec("UPDATE user SET bio=? WHERE username=?;", inc.NewBio, username)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with upload"})
		}
		
		c.JSON(http.StatusOK, gin.H{})
	}	

	return changeBio
}

func SearchUsersWrapper(db *sql.DB) gin.HandlerFunc {
	searchUsers := func (c *gin.Context) {
		query := c.Param("query")
		usernamesInSearch := []string{}
		query = "%" + query + "%"
		rows, err := db.Query("SELECT username FROM user WHERE username LIKE ? ORDER BY username;", query)
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with query"})
		}
		defer rows.Close()
		for rows.Next() {
			var username string
			rows.Scan(&username)
			usernamesInSearch = append(usernamesInSearch, username)
		}

		usersInSearch := []schemas.UserProfile{}
		for _, username := range usernamesInSearch {
			user := DBGetUserProfile(db, username)
			usersInSearch = append(usersInSearch, user)
		}
		c.JSON(http.StatusOK, usersInSearch)
	}

	return searchUsers
} 

func GetPostsWrapper(db *sql.DB) gin.HandlerFunc {
	getPosts := func (c *gin.Context) {
		ids := []int{}
		username, _ := c.Cookie("Username")
		rows, err := db.Query("SELECT id FROM post ORDER BY uploadtime DESC;")
		if err != nil {
			panic(err)
		}
		defer rows.Close()
		for rows.Next() {
			var id int
			rows.Scan(&id)
			ids = append(ids, id)
		}

		posts := []schemas.Post{}
		for _, id := range ids {
			post := DBGetPost(db, id)
			var liked bool
			var u string
			err := db.QueryRow(fmt.Sprintf("SELECT username FROM likes%d WHERE username=?", id), username).Scan(&u)
			if err == sql.ErrNoRows{
				liked = false
			} else {
				liked = true
			}
			post.Liked = liked
			posts = append(posts, post)
		}
		c.JSON(http.StatusOK, posts)
	}
	return getPosts
}

func CommentWrapper(db *sql.DB) gin.HandlerFunc {
	commentF := func (c *gin.Context) {
		id := c.Param("id")
		var comment schemas.Comment
		err := c.BindJSON(&comment)
		if err != nil {
			panic(err)
		}

		_, err = db.Exec(fmt.Sprintf("INSERT INTO comment%s (username, comment) VALUES ('%s', '%s');", id, comment.Username, comment.Text))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with comment"})
			panic(err)
		}

		c.JSON(http.StatusOK, comment)

	}
	return commentF
}

func GetLikeWrapper(db *sql.DB) gin.HandlerFunc {
	getLike := func (c *gin.Context) {
		id := c.Param("id")
		var likes int
		err := db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM likes%s;", id)).Scan(&likes)
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with getting likes"})
		}

		c.JSON(http.StatusOK, gin.H{"likes": likes})

	}
	return getLike
}

func LikeWrapper(db *sql.DB) gin.HandlerFunc {
	likePost := func (c *gin.Context) {
		id := c.Param("id")
		type user struct{
			Username string `json:"username"`
		}
		var inc user
		c.BindJSON(&inc)
		username := inc.Username

		//check if username is in likes
		exists := false
		var u string
		err := db.QueryRow(fmt.Sprintf("SELECT username FROM likes%s WHERE username=?", id), username).Scan(&u)
		if err == sql.ErrNoRows{
			exists = false
		} else {
			exists = true
		}
		
		//insert into table if exists, else delete
		if exists {
			fmt.Println("delete")
			tx, _ := db.Begin()
			_, err = tx.Exec(fmt.Sprintf("DELETE FROM likes%s WHERE username=?;", id), username)
			if err != nil{
				tx.Rollback()
			} else {
				tx.Commit()
			}
		} else {
			tx, _ := db.Begin()
			_, err = tx.Exec(fmt.Sprintf("INSERT INTO likes%s (username) VALUES (?);", id), username)
			if err != nil{
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}

		//get updated like count
		var likes int
		err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM likes%s;", id)).Scan(&likes)
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with getting likes"})
		}

		err = db.QueryRow(fmt.Sprintf("SELECT username FROM likes%s WHERE username=?", id), username).Scan(&u)
		if err == sql.ErrNoRows{
			exists = false
		} else {
			exists = true
		}

		c.JSON(http.StatusOK, gin.H{"likes": likes, "liked": exists})
	}

	return likePost
}

func AddFriendWrapper(db *sql.DB) gin.HandlerFunc {
	addFriend := func (c *gin.Context) {
		followed := c.Param("username")
		type user struct{
			Follower string `json:"username"`
		}
		var inc user
		c.BindJSON(&inc)
		follower := inc.Follower

		if follower == followed {
			c.JSON(http.StatusBadRequest, gin.H{"message": "error: cannot follow self!"})
		}
		//check if username is in likes
		following := false
		var u string
		err := db.QueryRow(fmt.Sprintf("SELECT username FROM followers%s WHERE username=?", followed), follower).Scan(&u)
		if err == sql.ErrNoRows{
			following = false
		} else {
			following = true
		}
		
		//insert into table if not exists, else delete
		if following {
			//unfollow
			tx, _ := db.Begin()
			_, err = tx.Exec(fmt.Sprintf("DELETE FROM followers%s WHERE username=?;", followed), follower)
			_, err = tx.Exec(fmt.Sprintf("DELETE FROM following%s WHERE username=?;", follower), followed)
			if err != nil{
				tx.Rollback()
			} else {
				tx.Commit()
			}
		} else {
			//follow
			tx, _ := db.Begin()
			_, err = tx.Exec(fmt.Sprintf("INSERT INTO followers%s (username) VALUES (?);", followed), follower)
			_, err = tx.Exec(fmt.Sprintf("INSERT INTO following%s (username) VALUES (?);", follower), followed)
			if err != nil{
				tx.Rollback()
			} else {
				tx.Commit()
			}
		}

		//get updated like count
		var followers int
		err = db.QueryRow(fmt.Sprintf("SELECT COUNT(*) FROM followers%s;", followed)).Scan(&followers)
		if err != nil {
			panic(err)
			c.JSON(http.StatusBadRequest, gin.H{"message": "error with friending"})
		}

		err = db.QueryRow(fmt.Sprintf("SELECT username FROM followers%s WHERE username=?", followed), follower).Scan(&u)
		if err == sql.ErrNoRows{
			following = false
		} else {
			following = true
		}

		c.JSON(http.StatusOK, gin.H{"followers": followers, "isFriend": following})
	}
	return addFriend
}


/* ---------------------------MIDDLWARE FUNCTIONS--------------------------- */

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

	c.Next()
}

/* ---------------------------HELPER FUNCTIONS--------------------------- */
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
