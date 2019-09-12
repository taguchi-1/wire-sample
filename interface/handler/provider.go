//+build wireinject

package handler

import "github.com/google/wire"

// ProviderSet handler provider
var ProviderSet = wire.NewSet(NewTodo)
