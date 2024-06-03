package main

import (
	"fmt"
	"database/sql"
	
	_ "modernc.org/sqlite"
	"github.com/gin-gonic/gin"
//	"github.com/gin-contrib/cors"
	"studyspotter/src"
)

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
	// Set up CORS
	router.Use(src.CORSMiddleware())

	authorized := router.Group("/") //, src.AuthRequired)
	authorized.GET("api/user", src.GetUsersWrapper(db))
	authorized.GET("api/user/:username", src.GetUserWrapper(db))
	router.POST("api/signup", src.CreateUserWrapper(db))
	router.POST("api/login", src.LoginWrapper(db))
	router.POST("api/post", src.PostWrapper(db))
	router.Run("localhost:8080")
}
