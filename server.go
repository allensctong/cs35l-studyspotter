package main

import (
	"fmt"
	"net/http"
	"database/sql"

	_ "modernc.org/sqlite"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"studyspotter/src"
)

type user struct {
	ID string `json:"id"`
	Password string `json:"password"`
}

var users = []user{
	{ID: "RS", Password: "123"},
	{ID: "SR", Password: "456"},
}

func GetUsers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, users)	
}

func GetUser(c *gin.Context) {
	id := c.Param("id")

	for _, u := range users {
		if u.ID == id {
			c.IndentedJSON(http.StatusOK, u)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})	
}

func CreateUserWrapper(db *sql.DB) gin.HandlerFunc {
	 CreateUser := func (c *gin.Context) {
		var newUser user

		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		users = append(users, newUser)
		db.Exec(fmt.Sprintf(`INSERT INTO user (id, password) 
			VALUES ("%s", "%s");`, newUser.ID, newUser.Password))
		fmt.Printf(fmt.Sprintf("INSERT INTO user (id, password)\nVALUES (%s, %s);\n", newUser.ID, newUser.Password))
		c.IndentedJSON(http.StatusCreated, newUser)
	}

	return CreateUser
}

func main() {
	// Set up database.
	dbName := "studyspotter.db"
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		fmt.Printf("Unable to use data source: %s", err)
	}
	defer db.Close()

	src.DbInit(db)

	// Set up router.
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
	}))
	router.GET("api/user", GetUsers)
	router.GET("api/user/:id", GetUser)
	router.POST("api/user", CreateUserWrapper(db))
	router.Run("localhost:8080")
}
