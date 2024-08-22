package bootstrap

import (
	krtlog "github.com/go-kratos/kratos/v2/log"

	"github.com/dawn303/cc/pkg/log"
)

func NewLogger(info AppInfo) krtlog.Logger {
	return krtlog.With(log.Default(),
		"ts", krtlog.DefaultTimestamp,
		"caller", krtlog.DefaultCaller,
		"service.id", info.ID,
		"service.name", info.Name,
		"service.version", info.Version,
	)
}
