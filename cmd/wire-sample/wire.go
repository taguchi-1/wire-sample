//+build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/taguchi-1/wire-sample/application"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/infra/persistence"
	"github.com/taguchi-1/wire-sample/interface/handler"
)

// InitializeTodoHandler InitializeTodoHandler
func InitializeTodoHandler() (todoHandler handler.Todo) {
	wire.Build(
		handler.NewTodo,
		application.NewTodo,
		service.NewTodo,
		persistence.NewTodo,
	)
	return
}
