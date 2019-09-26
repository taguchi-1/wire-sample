package service

import (
	"context"

	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/repository"
)

// Todo service interface
type Todo interface {
	Get(ctx context.Context, id string) (*entity.Todo, error)
}

type todoImpl struct {
	todoRepo repository.Todo
}

// NewTodo todo service constructor
func NewTodo(todoRepo repository.Todo) Todo {
	return &todoImpl{todoRepo}
}

func (s *todoImpl) Get(ctx context.Context, id string) (*entity.Todo, error) {
	return s.todoRepo.Get(ctx, id)
}
