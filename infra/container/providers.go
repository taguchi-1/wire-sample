//+build wireinject

package container

import (
	"github.com/google/wire"
	"github.com/taguchi-1/wire-sample/application"
	"github.com/taguchi-1/wire-sample/domain/service"
	"github.com/taguchi-1/wire-sample/infra/persistence"
	"github.com/taguchi-1/wire-sample/interface/handler"
)

// ProviderSet container provider
var ProviderSet = wire.NewSet(
	persistence.ProviderSet,
	service.ProviderSet,
	application.ProviderSet,
	handler.ProviderSet,
)