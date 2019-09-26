package application

import (
	"context"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/service"
)

func TestTodoGet(t *testing.T) {
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
					Todo: &entity.Todo{ID: "1", Title: "タイトル"},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			todoService := service.NewMockTodo(ctrl)
			todoService.EXPECT().Get(ctx, c.input.req.ID).MinTimes(1).Return(
				c.expect.res.Todo, nil,
			)

			todoApp, _ := NewTodo(todoService)
			actual, err := todoApp.Get(ctx, c.input.req)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.res, actual)
		})
	}
}
