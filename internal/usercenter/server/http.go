package server

import (
	"context"

	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/middleware/selector"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/gorilla/handlers"

	"github.com/dawn303/cc/internal/usercenter/service"
	v1 "github.com/dawn303/cc/pkg/api/usercenter/v1"
)

func NewWhiteListMatcher() selector.MatchFunc {
	whitelist := make(map[string]struct{})
	whitelist[v1.OperationUserCenterLogin] = struct{}{}
	return func(ctx context.Context, operation string) bool {
		if _, ok := whitelist[operation]; ok {
			return false
		}
		return true
	}
}

func NewHTTPServer(c *Config, gw *service.UserCenterService, middlewares []middleware.Middleware) *http.Server {
	opts := []http.ServerOption{
		http.Middleware(middlewares...),
		http.Filter(handlers.CORS(
			handlers.AllowedHeaders([]string{
				"X-Requested-With",
				"Content-Type",
				"Authorization",
				"X-Idempotent-ID",
			}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}),
			handlers.AllowedOrigins([]string{"*"}),
			handlers.AllowCredentials(),
		)),
	}
	if c.HTTP.Network != "" {
		opts = append(opts, http.Network(c.HTTP.Network))
	}
	if c.HTTP.Timeout != 0 {
		opts = append(opts, http.Timeout(c.HTTP.Timeout))
	}
	if c.HTTP.Addr != "" {
		opts = append(opts, http.Address(c.HTTP.Addr))
	}
	if c.TLS.UseTLS {
		opts = append(opts, http.TLSConfig(c.TLS.MustTLSConfig()))
	}

	srv := http.NewServer(opts...)
	v1.RegisterUserCenterHTTPServer(srv, gw)
	return srv
}
