//+build wireinject

package persistence

import "github.com/google/wire"

// ProviderSet repository provider
var ProviderSet = wire.NewSet(NewTodo)
