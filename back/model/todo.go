package model

import (
	"database/sql"
)

type Todo struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Completed int    `json:"completed"`
}

// CreateTodo create a new table if not exists in the database
func CreateTodo() error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS todolist (id SERIAL PRIMARY KEY NOT NULL,name TEXT NOT NULL,completed INTEGER)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// InsertTodo insert a new todo into the database
func InsertTodo(t *Todo) error {
	stmt, err := db.Prepare("INSERT INTO todolist (id, name, completed) VALUES (:id, :name, :completed)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sql.Named("id", t.ID), sql.Named("name", t.Name), sql.Named("completed", t.Completed))
	if err != nil {
		return err
	}
	return nil
}

// UpdateTodo update a todo in the database
func UpdateTodo(t *Todo) error {
	stmt, err := db.Prepare("UPDATE todolist SET name=:name, completed=:completed WHERE id=:id")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sql.Named("id", t.ID), sql.Named("name", t.Name), sql.Named("completed", t.Completed))
	if err != nil {
		return err
	}
	return nil
}

// DeleteTodo delete a todo in the database
func DeleteTodo(t *Todo) error {
	stmt, err := db.Prepare("DELETE FROM todolist WHERE id=:id")
	if err != nil {
		return err
	}
	_, err = stmt.Exec(sql.Named("id", t.ID))
	if err != nil {
		return err
	}
	return nil
}

// GetTodo get a todo from the database and return todo object
func GetTodo(t *Todo) (Todo, error) {
	var todo Todo
	stmt, err := db.Prepare("SELECT * FROM todolist WHERE id=;id")
	if err != nil {
		return todo, err
	}
	row := stmt.QueryRow(sql.Named("id", t.ID))
	if err := row.Scan(&todo.ID, &todo.Name, &todo.Completed); err != nil {
		return todo, err
	}
	return todo, nil
}

// GetTodos get all todos from the database and return todo slice
func GetTodos() ([]Todo, error) {
	var todos []Todo
	stmt, err := db.Prepare("SELECT * FROM todolist")
	if err != nil {
		return nil, err
	}
	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var todo Todo

		if err := rows.Scan(&todo.ID, &todo.Name, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
