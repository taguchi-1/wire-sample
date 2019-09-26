package persistence

import (
	"context"

	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/repository"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/hogedb"
)

type userImpl struct {
	db *hogedb.HogeDB
}

// NewUser constructor
func NewUser(db *hogedb.HogeDB) repository.User {
	return &userImpl{db}
}

func (r *userImpl) Get(ctx context.Context, id string) (*entity.User, error) {
	return &entity.User{
		ID:   "1",
		Name: "名前",
	}, nil
}
