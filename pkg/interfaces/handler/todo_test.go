package handler

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	gomock "github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/pkg/application"
	"github.com/taguchi-1/wire-sample/pkg/domain/entity"
)

func TestGet(t *testing.T) {
	type input struct {
		id string
	}
	type expect struct {
		code      int
		headerMap http.Header
		res       *entity.TodoResponse
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

			// create resquest context
			res := httptest.NewRecorder()
			ec := echo.New().NewContext(httptest.NewRequest("GET", "/", nil), res)
			ec.SetParamNames(todoIDParam)
			ec.SetParamValues(c.input.id)

			// create mock
			todoApp := application.NewMockTodo(ctrl)
			todoApp.EXPECT().Get(ec.Request().Context(), &entity.TodoGetRequest{
				ID: c.input.id,
			}).MinTimes(1).Return(
				c.expect.res, nil,
			)
			todoHandler, _ := NewTodo(todoApp)
			actual := &entity.TodoResponse{}

			// execute handler
			err := todoHandler.Get(ec)

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
