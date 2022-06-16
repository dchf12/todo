package handler

import (
	"net/http"
	"strconv"

	"github.com/dchf12/todo/model"
	"github.com/labstack/echo"
)

func AddTodo(c echo.Context) error {

	return c.JSON(http.StatusCreated, todo)
}

func GetTodos(c echo.Context) error {

	return c.JSON(http.StatusOK, todos)
}

func DeleteTodo(c echo.Context) error {

	return c.NoContent(http.StatusNoContent)
}

func UpdateTodo(c echo.Context) error {

	return c.NoContent(http.StatusNoContent)
}
