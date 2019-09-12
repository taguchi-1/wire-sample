//+build wireinject

package router

import "github.com/google/wire"

// ProviderSet handler provider
var ProviderSet = wire.NewSet(NewFront)
