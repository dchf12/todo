package main

import (
	"fmt"
	"log"

	table "github.com/dchf12/todo/table"
	_ "github.com/mattn/go-sqlite3"
)

const fileName = "todo.sqlite3"

func main() {
	// database connection
	db, err := table.InitDB(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	log.Println("Connected!")

	// create table if not exists
	n, err := table.Create(db)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(n)

	// read data from database
	// read all data from database
	// insert data into database
	// update data in database
	// delete data in database

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
