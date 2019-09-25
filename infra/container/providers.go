//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/taguchi-1/wire-sample/application"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/infra/hogedb"
	"github.com/taguchi-1/wire-sample/infra/persistence"
	"github.com/taguchi-1/wire-sample/interface/handler"
	"github.com/taguchi-1/wire-sample/interface/router"
)

// ProviderSet container provider
var ProviderSet = wire.NewSet(

	hogedb.ProviderSet,
	persistence.ProviderSet,
	service.ProviderSet,
	application.ProviderSet,
	handler.ProviderSet,
	router.ProviderSet,
)
