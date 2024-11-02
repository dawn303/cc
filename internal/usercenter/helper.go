package usercenter

import (
	"github.com/dawn303/cc/pkg/authn"
	jwtauthn "github.com/dawn303/cc/pkg/authn/jwt"
	"github.com/dawn303/cc/pkg/options"
	"github.com/golang-jwt/jwt/v4"
)

func NewAuthenticator(jwtOpts *options.JWTOptions) (authn.Authenticator, func(), error) {
	// Create options
	opts := []jwtauthn.Option{
		jwtauthn.WithIssuer("cc-usercenter"),
		jwtauthn.WithExpired(jwtOpts.Expired),
		jwtauthn.WithSigningKey([]byte(jwtOpts.Key)),
		jwtauthn.WithKeyFunc(func(t *jwt.Token) (any, error) {
			if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
				return nil, jwtauthn.ErrTokenInvalid
			}
			return []byte(jwtOpts.Key), nil
		}),
	}

	var method jwt.SigningMethod
	switch jwtOpts.SigningMethod {
	case "HS256":
		method = jwt.SigningMethodHS256
	case "HS384":
		method = jwt.SigningMethodHS384
	default:
		method = jwt.SigningMethodHS512
	}

	opts = append(opts, jwtauthn.WithSigningMethod(method))

	authn := jwtauthn.New(opts...)
	cleanFunc := func() {}
	return authn, cleanFunc, nil
}
