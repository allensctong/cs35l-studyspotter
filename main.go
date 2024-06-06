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

	// Set up router.
	router := gin.Default()
	// Set up CORS
	router.Use(src.CORSMiddleware())

	//authorized endpoints (MUST BE SIGNED IN TO ACCESS)
	authorized := router.Group("/", src.AuthRequired)
	authorized.GET("api/user/:username", src.GetUserWrapper(db))
//	authorized.GET("api/user/:username/search", src.TODO(db))
//	authorized.PUT("api/user/:username/bio", TODO)
//	authorized.PUT("api/user/:username/pfp", TODO)
	authorized.POST("api/post", src.CreatePostWrapper(db))
//	authorized.GET("api/post", src.GetPostWrapper(db))

	//unauthorized endpoints (anyone can call)
	router.POST("api/signup", src.CreateUserWrapper(db))
	router.POST("api/login", src.LoginWrapper(db))
	router.Static("/assets", "./assets")

	//start router
	router.Run("localhost:8080")
}
