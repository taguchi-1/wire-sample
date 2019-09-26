package service

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
	"github.com/taguchi-1/wire-sample/pkg/domain/repository"
)

func TestTodoGet(t *testing.T) {
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
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			todoRepo := repository.NewMockTodo(ctrl)
			todoRepo.EXPECT().Get(ctx, c.input.id).MinTimes(1).Return(
				c.expect.todo, nil,
			)
			todoService := NewTodo(todoRepo)
			actual, err := todoService.Get(ctx, c.input.id)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.todo, actual)
		})
	}
}
