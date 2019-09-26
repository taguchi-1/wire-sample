//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/taguchi-1/wire-sample/pkg/application"
	"github.com/taguchi-1/wire-sample/pkg/domain/service"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/hogedb"
	"github.com/taguchi-1/wire-sample/pkg/infrastructure/persistence"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/handler"
	"github.com/taguchi-1/wire-sample/pkg/interfaces/router"
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
