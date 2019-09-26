package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/pkg/application"
	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
)

// Todo handler interface
type Todo interface {
	Get(c echo.Context) error
}

type todoImpl struct {
	todoApp application.Todo
}

const (
	todoIDParam string = "todoID"
)

//NewTodo  handler constructor
func NewTodo(todoApp application.Todo) (Todo, error) {
	return &todoImpl{todoApp}, nil
}

// Get get handler
func (h *todoImpl) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param(todoIDParam)

	req := entity.NewTodoRequest(id)
	todo, err := h.todoApp.Get(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, id)
	}
	return c.JSON(http.StatusOK, todo)
}
