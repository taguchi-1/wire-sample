package application

import (
	"context"

	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/service"
)

// Todo todo application service
type Todo interface {
	Get(ctx context.Context, req *entity.TodoGetRequest) (*entity.TodoResponse, error)
}

type todoImpl struct {
	todoService service.Todo
}

// NewTodo Constructor
func NewTodo(todoService service.Todo) (Todo, error) {
	return &todoImpl{todoService}, nil
}

func (app *todoImpl) Get(ctx context.Context, req *entity.TodoGetRequest) (*entity.TodoResponse, error) {
	todo, err := app.todoService.Get(ctx, req.ID)
	if err != nil {
		return nil, err
	}
	return entity.NewTodoResponse(todo), nil
}
