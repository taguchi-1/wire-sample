package repository

import (
	"context"

	"github.com/taguchi-1/wire-sample/domain/entity"
)

// User repository interface
type User interface {
	Get(ctx context.Context, id string) (*entity.User, error)
}
