//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package usercenter

//go:generate go run github.com/google/wire/cmd/wire

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/google/wire"

	"github.com/dawn303/cc/internal/pkg/bootstrap"
	"github.com/dawn303/cc/internal/usercenter/biz"
	"github.com/dawn303/cc/internal/usercenter/server"
	"github.com/dawn303/cc/internal/usercenter/service"
	"github.com/dawn303/cc/internal/usercenter/store"
	"github.com/dawn303/cc/pkg/db"
)

func wireApp(
	bootstrap.AppInfo,
	*server.Config,
	*db.MySQLOptions,
) (*kratos.App, func(), error) {
	wire.Build(
		bootstrap.ProviderSet,
		server.ProviderSet,
		store.ProviderSet,
		db.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
	)

	return nil, nil, nil
}
