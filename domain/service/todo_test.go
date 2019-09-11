package service

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/infra/persistence"
)

func TestGet(t *testing.T) {
	type input struct {
		id string
	}
	type expect struct {
		todo *entity.Todo
	}
	cases := []struct {
		name   string
		input  input
		expect expect
	}{
		{
			name: "Test Case A.",
			input: input{
				id: "1",
			},
			expect: expect{
				todo: &entity.Todo{
					ID:    "1",
					Title: "タイトル",
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			todoService := NewTodo(persistence.NewTodo())
			actual, err := todoService.Get(ctx, c.input.id)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.todo, actual)
		})
	}
}
