//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/taguchi-1/wire-sample/interface/handler"
)

// InitializeTodoHandler InitializeTodoHandler
func InitializeTodoHandler() (todoHandler handler.Todo, err error) {
	wire.Build(ProviderSet)
	return
}
