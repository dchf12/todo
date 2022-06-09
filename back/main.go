package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	// database connection
	db, err := sql.Open("sqlite3", "./todo.sqlite3")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	log.Println("Connected!")

	// create table if not exists
	s := "todolist"
	cmd := fmt.Sprintf("CREATE TABLE IF NOT EXISTS %s (id SERIAL PRIMARY KEY NOT NULL,title TEXT NOT NULL,completed INTEGER)", s)
	r, err := db.Exec(cmd)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r)
	// reade data from database
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(rows)

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
