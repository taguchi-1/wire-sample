package handler

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/pkg/application"
	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
)

// User handler interface
type User interface {
	Get(c echo.Context) error
}

type userImpl struct {
	userApp application.User
}

const (
	userIDParam string = "userID"
)

//NewUser  handler constructor
func NewUser(userApp application.User) (User, error) {
	return &userImpl{userApp}, nil
}

// Get get handler
func (h *userImpl) Get(c echo.Context) error {
	ctx := c.Request().Context()
	id := c.Param(userIDParam)

	req := entity.NewUserRequest(id)
	user, err := h.userApp.Get(ctx, req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, id)
	}
	return c.JSON(http.StatusOK, user)
}
