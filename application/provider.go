//+build wireinject

package application

import "github.com/google/wire"

// ProviderSet application provider
var ProviderSet = wire.NewSet(NewTodo, NewUser)
