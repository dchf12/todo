package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

type Todo struct {
	ID        int64
	Title     string
	Completed int
}

const fileName = "todo.sqlite3"

func main() {
	// database connection
	db, err := initDB(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected!")

	// create table if not exists
	n, err := createTable(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

	// insert data into database
	// update data in database
	// delete data in database
	// read data from database
	// read all data from database

	// //html template stuff
	// tmpl := template.Must(template.ParseFiles("../dist/index.html"))

	// http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {

	// 	tmpl.Execute(w, nil)
	// })

	// //static assets
	// fs := http.FileServer(http.Dir("../dist/assets/"))
	// http.Handle("/static/", http.StripPrefix("/static/", fs))

	// http.ListenAndServe(":8080", nil)
}

// initDB create a new database
func initDB(s string) (*sql.DB, error) {
	db, err := sql.Open("sqlite3", s)
	if err != nil {
		return nil, err
	}
	pingErr := db.Ping()
	if pingErr != nil {
		return nil, pingErr
	}
	return db, nil
}
