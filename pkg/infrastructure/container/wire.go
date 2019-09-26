//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/hogedb"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/handler"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/router"
)

// InitializeTodoHandler InitializeTodoHandler
func InitializeTodoHandler(e *echo.Echo, env hogedb.Env) (todoHandler handler.Todo, err error) {
	wire.Build(ProviderSet)
	return
}

// InitializeFrontRouter InitializeFrontRouter
func InitializeFrontRouter(e *echo.Echo, env hogedb.Env) (frontRouter router.Front, err error) {
	wire.Build(ProviderSet)
	return
}

// InitializeBackgroundRouter InitializeBackgroundRouter
func InitializeBackgroundRouter(e *echo.Echo) (backgroundRouter router.Background, err error) {
	wire.Build(ProviderSet)
	return
}
