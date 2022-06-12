package table

import "database/sql"

type Todo struct {
	ID        int64  `json:"id"`
	Title     string `json:"title"`
	Completed int    `json:"completed"`
}

// initDB create a new database
func InitDB(s string) (*sql.DB, error) {
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
func Create(db *sql.DB) error {
	stmt, err := db.Prepare("CREATE TABLE IF NOT EXISTS todolist (id SERIAL PRIMARY KEY NOT NULL,title TEXT NOT NULL,completed INTEGER)")
	if err != nil {
		return err
	}
	_, err = stmt.Exec()
	if err != nil {
		return err
	}
	return nil
}

// insertTodo insert a new todo into the database and return the id
func InsertTodo(todo Todo, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO todolist (id, title, completed) VALUES (:id, :title, :completed)")
	if err != nil {
		return 0, err
	}
	r, err := stmt.Exec(sql.Named("id", todo.ID), sql.Named("title", todo.Title), sql.Named("completed", todo.Completed))
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
func UpdateTodo(todo Todo, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("UPDATE todolist SET title=:title, completed=:completed WHERE id=:id")
	if err != nil {
		return 0, err
	}
	r, err := stmt.Exec(sql.Named("id", todo.ID), sql.Named("title", todo.Title), sql.Named("completed", todo.Completed))
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
func DeleteTodo(todo Todo, db *sql.DB) (int64, error) {
	stmt, err := db.Prepare("DELETE FROM todolist WHERE id=:id")
	if err != nil {
		return 0, err
	}
	r, err := stmt.Exec(sql.Named("id", todo.ID))
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
func GetTodo(id int64, db *sql.DB) (Todo, error) {
	var todo Todo
	stmt, err := db.Prepare("SELECT * FROM todolist WHERE id=;id")
	if err != nil {
		return todo, err
	}
	row := stmt.QueryRow(sql.Named("id", id))
	if err := row.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
		return todo, err
	}
	return todo, nil
}

// getTodos get all todos from the database and return the id
func GetTodos(db *sql.DB) ([]Todo, error) {
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

		if err := rows.Scan(&todo.ID, &todo.Title, &todo.Completed); err != nil {
			return nil, err
		}
		todos = append(todos, todo)
	}
	return todos, nil
}
