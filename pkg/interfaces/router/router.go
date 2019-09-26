package router

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/handler"
)

// Front Distinguishing type
type Front *echo.Echo

// Background Distinguishing type
type Background *echo.Echo

// NewFront front router construtor
func NewFront(
	e *echo.Echo,
	todoHandler handler.Todo,
	userHandler handler.User,
) Front {
	e.GET("/todos/:todoID", todoHandler.Get)
	e.GET("/users/:userID", userHandler.Get)
	return e
}

// NewBackground backgrond router construtor
func NewBackground(
	e *echo.Echo,
) Background {
	e.GET("/health", func(c echo.Context) error { return c.String(http.StatusOK, "OK") })
	return e
}
