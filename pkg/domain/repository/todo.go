package repository

import (
	"context"

	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
)

// Todo repository interface
type Todo interface {
	Get(ctx context.Context, id string) (*entity.Todo, error)
}
