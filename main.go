package main

import (
	"fmt"
	"database/sql"
	
	_ "modernc.org/sqlite"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
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
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
	}))
	router.GET("api/user", src.GetUsersWrapper(db))
	router.GET("api/user/:username", src.GetUserWrapper(db))
	router.POST("api/signup", src.CreateUserWrapper(db))
	router.POST("api/login", src.LoginWrapper(db))
	router.Run("localhost:8080")
}
