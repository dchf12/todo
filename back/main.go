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

// createTable create a new table if not exists in the database
func createTable(db *sql.DB) (int64, error) {
	r, err := db.Exec("CREATE TABLE IF NOT EXISTS todolist (id SERIAL PRIMARY KEY NOT NULL,title TEXT NOT NULL,completed INTEGER)")
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// insertTodo insert a new todo into the database and return the id
func insertTodo(todo Todo, db *sql.DB) (int64, error) {
	r, err := db.Exec("INSERT INTO %s (title, completed) VALUES (?, ?)", todo.Title, todo.Completed)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// updateTodo update a todo in the database and return the id
func updateTodo(todo Todo, db *sql.DB) (int64, error) {
	r, err := db.Exec("UPDATE todolist SET title=?, completed=? WHERE id=?", todo.Title, todo.Completed, todo.ID)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// deleteTodo delete a todo in the database and return the id
func deleteTodo(todo Todo, db *sql.DB) (int64, error) {
	r, err := db.Exec("DELETE FROM todolist WHERE id=?", todo.ID)
	if err != nil {
		return 0, err
	}
	id, err := r.LastInsertId()
	if err != nil {
		return 0, err
	}
	return id, nil
}

// getTodo get a todo from the database and return the id
func getTodo(id int64, db *sql.DB) (Todo, error) {
	var todo Todo
	row := db.QueryRow("SELECT * FROM todolist WHERE id=?", id)
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
		return todo, err
	}
	return todo, nil
}

// getTodos get all todos from the database and return the id
func getTodos(db *sql.DB) ([]Todo, error) {
	var todos []Todo
	rows, err := db.Query("SELECT * FROM todolist")
	if err != nil {
		return todos, err
	}
	for rows.Next() {
		var todo Todo

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
