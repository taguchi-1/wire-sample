//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/interface/handler"
	"github.com/taguchi-1/wire-sample/interface/router"
)

// InitializeTodoHandler InitializeTodoHandler
func InitializeTodoHandler() (todoHandler handler.Todo, err error) {
	wire.Build(ProviderSet)
	return
}

// InitializeFrontRouter InitializeFrontRouter
func InitializeFrontRouter(e *echo.Echo) (frontRouter router.Front, err error) {
	wire.Build(ProviderSet)
	return
}

// InitializeBackgroundRouter InitializeBackgroundRouter
func InitializeBackgroundRouter(e *echo.Echo) (backgroundRouter router.Background, err error) {
	wire.Build(ProviderSet)
	return
}
