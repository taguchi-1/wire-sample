package persistence

import (
	"context"

	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/repository"
	"github.com/taguchi-1/wire-sample/infra/hogedb"
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
