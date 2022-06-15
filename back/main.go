package main

import (
	"log"

	"github.com/dchf12/todo/table"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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
	if err := table.Create(db); err != nil {
		log.Fatal(err)
	}

	e := echo.New()

	//html template stuff
	e.File("/", "../front/dist/index.html")
	//static assets
	e.Static("/static", "../front/dist/assets/")

	e.Use(middleware.Recover())
	e.Use(middleware.Logger())

	// routing
	e.GET("/", table.GetTodos)
	e.POST("/", table.InsertTodo)
	e.PUT("/", table.UpdateTodo)
	e.DELETE("/", table.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
