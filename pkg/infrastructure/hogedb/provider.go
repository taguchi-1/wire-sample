//+build wireinject

package hogedb

import "github.com/google/wire"

// ProviderSet hogedb provider
var ProviderSet = wire.NewSet(
	NewConfigA,
	NewConfigB,
	wire.Struct(new(HogeDB), "AConfig", "BConfig"),
)
