package handler

import (
	"net/http"
	"strconv"

	"github.com/dchf12/todo/model"
	"github.com/labstack/echo"
)

func AddTodo(c echo.Context) error {
	todo := new(model.Todo)
	model.CreateTodo()
	if err := c.Bind(todo); err != nil {
		return err
	}
	model.InsertTodo(todo)

	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {
	todos, err := model.GetTodos()
	if err != nil {
		return err
	}
	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {
	todoID, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		return echo.ErrNotFound
	}
	if err := model.DeleteTodo(&model.Todo{ID: int64(todoID)}); err != nil {
		return echo.ErrNotFound
	}
	return c.NoContent(http.StatusNoContent)
}

func UpdateTodo(c echo.Context) error {

	return c.NoContent(http.StatusNoContent)
}
