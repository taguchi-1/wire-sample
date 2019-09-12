//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/interface/handler"
)

// InitializeTodoHandler InitializeTodoHandler
func InitializeTodoHandler() (todoHandler handler.Todo) {
	wire.Build(ProviderSet)
	return
}

// InitializeTodoService InitializeTodoService
func InitializeTodoService() (todoService service.Todo) {
	wire.Build(ProviderSet)
	return
}
