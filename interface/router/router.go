package router

import (
	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/interface/handler"
)

// NewFront front router construtor
func NewFront(
	todoHandler handler.Todo,
	userHandler handler.User,
) *echo.Echo {
	e := echo.New()
	e.GET("/todos/:todoID", todoHandler.Get)
	e.GET("/users/:userID", userHandler.Get)
	return e
}
