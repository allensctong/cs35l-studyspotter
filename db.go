package main

import (
	"fmt"
	"database/sql"
	_ "modernc.org/sqlite"
)

func DbInit(dbName string) {
	db, err := sql.Open("sqlite", dbName)
	if err != nil {
		fmt.Printf("Unable to use data source: %s", err)
	}
	defer db.Close()

	if _, err := db.Exec("create table users(UID INT, Username VARCHAR(255), Password VARCHAR(255));");
	err != nil {
		panic(err)
	}
}

/*func main() {
	DbInit("studyspotter.db")
}*/
