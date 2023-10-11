package models

import (
	"database/sql"
	"fmt"

	 _"github.com/mattn/go-sqlite3"
)

var (
	db *sql.DB
)

func Setup() {
	var err error
	const file string = "todos.db"
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println("Could not connect to DB:", err)
	}

	err = db.Ping()
	if err != nil {
		panic("Could not ping DB")
	}
}