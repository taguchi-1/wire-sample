package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewTodoRequest(t *testing.T) {
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
			actual := NewTodoRequest(c.input.id)
			assert.NotNil(t, actual)
		})
	}
}
