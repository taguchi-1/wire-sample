package application

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/domain/entity"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/infra/persistence"
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
					User: entity.User{ID: "1", Name: "名前"},
				},
			},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			ctx := context.Background()
			userRepo := persistence.NewUser()
			userService := service.NewUser(userRepo)
			userApp, _ := NewUser(userService)

			actual, err := userApp.Get(ctx, c.input.req)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.res, actual)
		})
	}
}
