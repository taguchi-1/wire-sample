package application

import (
	"context"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/service"
)

func TestUserGet(t *testing.T) {
	type input struct {
		req *entity.UserGetRequest
	}
	type expect struct {
		res *entity.UserResponse
	}
	cases := []struct {
		name   string
		input  input
		expect expect
	}{
		{
			name: "Test Case A.",
			input: input{
				req: &entity.UserGetRequest{ID: "1"},
			},
			expect: expect{
				res: &entity.UserResponse{
					User: &entity.User{ID: "1", Name: "名前"},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			ctx := context.Background()
			userService := service.NewMockUser(ctrl)
			userService.EXPECT().Get(ctx, c.input.req.ID).MinTimes(1).Return(
				c.expect.res.User, nil,
			)

			userApp, _ := NewUser(userService)
			actual, err := userApp.Get(ctx, c.input.req)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.res, actual)
		})
	}
}
