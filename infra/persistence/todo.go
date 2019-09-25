package persistence

import (
	"context"

	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/repository"
	"github.com/taguchi-1/wire-sample/infra/hogedb"
)

type todoImpl struct {
	db *hogedb.HogeDB
}

// NewTodo constructor
func NewTodo(db *hogedb.HogeDB) repository.Todo {
	return &todoImpl{db}
}

func (r *todoImpl) Get(ctx context.Context, id string) (*entity.Todo, error) {
	return &entity.Todo{
		ID:    "1",
		Title: "タイトル",
	}, nil
}
