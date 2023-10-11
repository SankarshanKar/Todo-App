package models

import (
	"database/sql"
)

var (
	db *sql.DB
)

func Setup() {
	var err error
	const file string = "todos.db"
	db, err = sql.Open("sqlite3", file)
	if err != nil {
		panic("Could not connect to DB")
	}

	err = db.Ping()
	if err != nil {
		panic("Could not ping DB")
	}
}