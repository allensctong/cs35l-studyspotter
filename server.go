package main

import (
	"fmt"
	"log"
	"net/http"
	"database/sql"

	_ "modernc.org/sqlite"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/google/uuid"
//	"studyspotter/src"
)

type User struct {
	ID string `json:"id"`
	Password string `json:"password"`
}

func GetUsersWrapper(db *sql.DB) gin.HandlerFunc {
	GetUsers := func (c *gin.Context) {
		var users = []User{}
		var id string
		var password string

		rows, err := db.Query("SELECT id, password FROM user")
		if err != nil {
			log.Fatal(err)
		}
		defer rows.Close()

		for rows.Next() {
			if err := rows.Scan(&id, &password); err != nil {
				log.Fatal(err)
			}
			users = append(users, User{id, password})
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
		id := c.Param("id")
		var password string

		err := db.QueryRow("SELECT id, password FROM user WHERE id=?", id).Scan(&id, &password)
		
		if err == sql.ErrNoRows {
			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "user not found"})
			return
		}
		if err != nil {
			panic(err)
		}
		c.IndentedJSON(http.StatusOK, User{id, password})
	}

	return GetUser
}

func CreateUserWrapper(db *sql.DB) gin.HandlerFunc {
	 CreateUser := func (c *gin.Context) {
		var newUser User

		if err := c.BindJSON(&newUser); err != nil {
			return
		}

		err := db.QueryRow("SELECT id, password FROM user WHERE id=?", newUser.ID)

		db.Exec(fmt.Sprintf(`INSERT INTO user (id, password) 
			VALUES ("%s", "%s");`, newUser.ID, newUser.Password))
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
		return
	}
	defer db.Close()

	//src.DbInit(db)

	// Set up router.
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins: []string{"http://localhost:5173"},
	}))
	router.GET("api/user", GetUsersWrapper(db))
	router.GET("api/user/:id", GetUserWrapper(db))
	router.POST("api/user", CreateUserWrapper(db))
	router.Run("localhost:8080")
}
