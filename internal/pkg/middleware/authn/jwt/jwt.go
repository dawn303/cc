package jwt

import (
	"context"
	"strings"

	"github.com/dawn303/cc/pkg/authn"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
)

const (
	reason           string = "Unauthorized"
	bearerWord       string = "Bearer"
	bearerFormat     string = "Bearer %s"
	authorizationKey string = "Authorization"
)

var (
	ErrMissingJwtToken = errors.Unauthorized(reason, "JWT token is missing")
	ErrWrongContext    = errors.Unauthorized(reason, "Wrong context for middleware")
)

func Server(a authn.Authenticator) middleware.Middleware {
	return func(h middleware.Handler) middleware.Handler {
		return func(ctx context.Context, req any) (any, error) {
			if tr, ok := transport.FromServerContext(ctx); ok {
				auths := strings.SplitN(tr.RequestHeader().Get(authorizationKey), " ", 2)
				if len(auths) != 2 || !strings.EqualFold(auths[0], bearerWord) {
					return nil, ErrMissingJwtToken
				}

				// 验证token
				accessToken := auths[1]
				_, err := a.ParseClaims(ctx, accessToken)
				if err != nil {
					return nil, err
				}

				return h(ctx, req)
			}
			return nil, ErrWrongContext
		}
	}
}
