package database

import (
	"database/sql"
	"fmt"
	"log"
)

var db *sql.DB

func InitDB(datasource string) *sql.DB{
	var err error
	db, err = sql.Open("sqlite3", datasource)
	if err != nil {
		log.Fatal(err)
		fmt.Println("Failed to create the handle")
	}
	fmt.Println("database access established")
	//defer db.Close()
	//}
	return db
}