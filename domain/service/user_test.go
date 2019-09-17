package service

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/repository"
)

func TestUserGet(t *testing.T) {
	type input struct {
		id string
	}
	type expect struct {
		user *entity.User
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
				user: &entity.User{
					ID:   "1",
					Name: "名前",
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			userRepo := repository.NewMockUser(ctrl)
			userRepo.EXPECT().Get(ctx, c.input.id).MinTimes(1).Return(
				c.expect.user, nil,
			)
			userService := NewUser(userRepo)
			actual, err := userService.Get(ctx, c.input.id)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.user, actual)
		})
	}
}
