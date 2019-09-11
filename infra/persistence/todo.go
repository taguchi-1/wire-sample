package persistence

import (
	"context"

	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/repository"
)

type todoImpl struct {
}

// NewTodo constructor
func NewTodo() repository.Todo {
	return &todoImpl{}
}

func (r *todoImpl) Get(ctx context.Context, id string) (*entity.Todo, error) {
	return &entity.Todo{
		ID:    "1",
		Title: "タイトル",
	}, nil
}
