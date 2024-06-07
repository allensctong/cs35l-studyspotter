package main

import (
	"fmt"
	"os"
	"errors"
	"database/sql"
	
	_ "modernc.org/sqlite"
	"github.com/gin-gonic/gin"
//	"github.com/gin-contrib/cors"
	"studyspotter/src"
)

func main() {
	//this is not ideal but i didnt have time to implement a proper secret manager
	src.GenerateKey()
	
	// Set up database.
	dbName := "studyspotter.db"

	//check if init is needed
	init := false
	_, err := os.OpenFile(dbName, os.O_WRONLY, 0444)
	if errors.Is(err, os.ErrNotExist) {
		init = true
	}
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		fmt.Printf("Unable to use data source: %s", err)
		return
	}
	defer db.Close()

	if init {
		fmt.Println("DB set up!")
		src.DbInit(db)
	}

	// Set up router.
	router := gin.Default()
	// Set up CORS
	router.Use(src.CORSMiddleware())

	//authorized endpoints (MUST BE SIGNED IN TO ACCESS)
	authorized := router.Group("/", src.AuthRequired)
	authorized.GET("api/user/:username", src.GetUserWrapper(db))
	authorized.GET("api/user/search/:query", src.SearchUsersWrapper(db))
	authorized.PUT("api/user/:username/bio", src.ChangeBioWrapper(db))
	authorized.PUT("api/user/:username/pfp", src.ChangePfpWrapper(db))
	authorized.PUT("api/user/:username/friend", src.AddFriendWrapper(db))
	authorized.POST("api/post", src.CreatePostWrapper(db))
	authorized.GET("api/post", src.GetPostsWrapper(db))
	authorized.POST("api/post/:id/comment", src.CommentWrapper(db))
	authorized.PUT("api/post/:id/like", src.LikeWrapper(db))
	authorized.GET("api/post/:id/like", src.GetLikeWrapper(db))

	//unauthorized endpoints (anyone can call)
	router.POST("api/signup", src.CreateUserWrapper(db))
	router.POST("api/login", src.LoginWrapper(db))
	router.Static("/assets", "./assets")

	//start router
	router.Run("localhost:8080")
}
