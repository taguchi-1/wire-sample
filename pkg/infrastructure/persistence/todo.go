package persistence

import (
	"context"

	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/repository"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/hogedb"
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
