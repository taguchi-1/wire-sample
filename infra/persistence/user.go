package persistence

import (
	"context"

	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/repository"
)

type userImpl struct {
}

// NewUser constructor
func NewUser() repository.User {
	return &userImpl{}
}

func (r *userImpl) Get(ctx context.Context, id string) (*entity.User, error) {
	return &entity.User{
		ID:   "1",
		Name: "名前",
	}, nil
}
