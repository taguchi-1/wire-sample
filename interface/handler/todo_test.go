package handler

import (
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
	"github.com/taguchi-1/wire-sample/application"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/infra/persistence"
)

func TestGet(t *testing.T) {
	type input struct {
		id string
	}
	type expect struct {
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
			expect: expect{},
		},
	}
	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			todoApp, _ := application.NewTodo(service.NewTodo(persistence.NewTodo()))
			todoHandler, _ := NewTodo(todoApp)
			ec := echo.New().NewContext(httptest.NewRequest("GET", "/", nil), httptest.NewRecorder())
			ec.SetParamNames(todoIDParam)
			ec.SetParamValues(c.input.id)
			err := todoHandler.Get(ec)
			assert.Nil(t, err)
		})
	}
}
