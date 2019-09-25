package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/application"
	"github.com/taguchi-1/wire-sample/domain/entity"
)

func TestUserGet(t *testing.T) {
	type input struct {
		id string
	}
	type expect struct {
		code      int
		headerMap http.Header
		res       *entity.UserResponse
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
				code: http.StatusOK,
				headerMap: http.Header{
					"Content-Type": {"application/json; charset=UTF-8"},
				},
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

			// create resquest context
			res := httptest.NewRecorder()
			ec := echo.New().NewContext(httptest.NewRequest("GET", "/", nil), res)
			ec.SetParamNames(userIDParam)
			ec.SetParamValues(c.input.id)

			// create mock
			userApp := application.NewMockUser(ctrl)
			userApp.EXPECT().Get(ec.Request().Context(), &entity.UserGetRequest{
				ID: c.input.id,
			}).MinTimes(1).Return(
				c.expect.res, nil,
			)
			userHandler, _ := NewUser(userApp)
			actual := &entity.UserResponse{}

			// execute handler
			err := userHandler.Get(ec)

			// check response code & header & body
			assert.Nil(t, err)
			assert.Equal(t, c.expect.code, res.Code)
			assert.Equal(t, c.expect.headerMap, res.HeaderMap)
			err = json.Unmarshal(res.Body.Bytes(), actual)
			assert.Nil(t, err)
			assert.Equal(t, c.expect.res, actual)
		})
	}
}
