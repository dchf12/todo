package main

import (
	"github.com/dchf12/todo/handler"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {

	e := echo.New()

	//html template stuff
	e.File("/", "../front/dist/index.html")
	//static assets
	e.Static("/static", "../front/dist/assets/")

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// routing
	e.GET("/", handler.GetTodos)
	e.POST("/", handler.AddTodo)
	e.PUT("/", handler.UpdateTodo)
	e.DELETE("/", handler.DeleteTodo)

	e.Logger.Fatal(e.Start(":8080"))
}
