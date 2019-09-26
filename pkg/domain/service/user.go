package service

import (
	"context"

	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/repository"
)

// User service interface
type User interface {
	Get(ctx context.Context, id string) (*entity.User, error)
}

type userImpl struct {
	userRepo repository.User
}

// NewUser user service constructor
func NewUser(userRepo repository.User) User {
	return &userImpl{userRepo}
}

func (s *userImpl) Get(ctx context.Context, id string) (*entity.User, error) {
	return s.userRepo.Get(ctx, id)
}
