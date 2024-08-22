package server

import (
	"context"
	"encoding/json"

	"github.com/dawn303/cc/pkg/log"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewServers, NewHTTPServer, NewMiddlewares)

func NewServers(hs *http.Server) []transport.Server {
	return []transport.Server{hs}
}

func NewMiddlewares() []middleware.Middleware {
	return []middleware.Middleware{
		recovery.Recovery(
			recovery.WithHandler(func(ctx context.Context, rq, err any) error {
				data, _ := json.Marshal(rq)
				log.C(ctx).Errorw(err.(error), "Catching a panic", "rq", string(data))
				return nil
			}),
		),
		// selector.Server(jwt.Server(a)).Match(NewWhiteListMatcher()).Build(),
		// validate.Validator(v),
		// logging.Server(logger),
	}
}
