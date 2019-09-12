package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/infra/persistence"
)

func TestGet(t *testing.T) {
	type input struct {
		req *entity.TodoGetRequest
	}
	type expect struct {
		res *entity.TodoResponse
	}
	cases := []struct {
		name   string
		input  input
		expect expect
	}{
		{
			name: "Test Case A.",
			input: input{
				req: &entity.TodoGetRequest{ID: "1"},
			},
			expect: expect{
				res: &entity.TodoResponse{
					Todo: entity.Todo{ID: "1", Title: "タイトル"},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			todoRepo := persistence.NewTodo()
			todoService := service.NewTodo(todoRepo)
			todoApp, _ := NewTodo(todoService)

			actual, err := todoApp.Get(ctx, c.input.req)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.res, actual)
		})
	}
}
