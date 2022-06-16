package model

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

var db *sql.DB

// init create a new database
func init() {
	// database connection
	var err error
	db, err = sql.Open("sqlite3", "todo.sqlite3")

	if err != nil {
		panic("failed to connect database")
	}
}
