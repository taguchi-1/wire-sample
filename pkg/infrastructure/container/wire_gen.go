// Code generated by Wire. DO NOT EDIT.

//go:generate wire
//+build !wireinject

package container

import (
	"github.com/labstack/echo"
	"github.com/taguchi-1/wire-sample/pkg/application"
	"github.com/taguchi-1/wire-sample/pkg/domain/service"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/hogedb"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/persistence"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/handler"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/router"
)

// Injectors from wire.go:

func InitializeTodoHandler(e *echo.Echo, env hogedb.Env) (handler.Todo, error) {
	aConfig := hogedb.NewConfigA(env)
	bConfig := hogedb.NewConfigB(env)
	hogeDB := &hogedb.HogeDB{
		AConfig: aConfig,
		BConfig: bConfig,
	}
	todo := persistence.NewTodo(hogeDB)
	serviceTodo := service.NewTodo(todo)
	applicationTodo, err := application.NewTodo(serviceTodo)
	if err != nil {
		return nil, err
	}
	handlerTodo, err := handler.NewTodo(applicationTodo)
	if err != nil {
		return nil, err
	}
	return handlerTodo, nil
}

func InitializeFrontRouter(e *echo.Echo, env hogedb.Env) (router.Front, error) {
	aConfig := hogedb.NewConfigA(env)
	bConfig := hogedb.NewConfigB(env)
	hogeDB := &hogedb.HogeDB{
		AConfig: aConfig,
		BConfig: bConfig,
	}
	todo := persistence.NewTodo(hogeDB)
	serviceTodo := service.NewTodo(todo)
	applicationTodo, err := application.NewTodo(serviceTodo)
	if err != nil {
		return nil, err
	}
	handlerTodo, err := handler.NewTodo(applicationTodo)
	if err != nil {
		return nil, err
	}
	user := persistence.NewUser(hogeDB)
	serviceUser := service.NewUser(user)
	applicationUser, err := application.NewUser(serviceUser)
	if err != nil {
		return nil, err
	}
	handlerUser, err := handler.NewUser(applicationUser)
	if err != nil {
		return nil, err
	}
	front := router.NewFront(e, handlerTodo, handlerUser)
	return front, nil
}

func InitializeBackgroundRouter(e *echo.Echo) (router.Background, error) {
	background := router.NewBackground(e)
	return background, nil
}