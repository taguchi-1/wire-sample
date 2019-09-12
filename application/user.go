package application

import (
	"context"

	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/service"
)

// User user application service
type User interface {
	Get(ctx context.Context, req *entity.UserGetRequest) (*entity.UserResponse, error)
}

type userImpl struct {
	userService service.User
}

// NewUser Constructor
func NewUser(userService service.User) (User, error) {
	return &userImpl{userService}, nil
}

func (app *userImpl) Get(ctx context.Context, req *entity.UserGetRequest) (*entity.UserResponse, error) {
	user, err := app.userService.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return entity.NewUserResponse(user), nil
}
