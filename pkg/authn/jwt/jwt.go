package jwt

import (
	"context"
	"time"

	"github.com/dawn303/cc/pkg/authn"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v4"
	"github.com/nicksnyder/go-i18n/v2/i18n"
)

const (
	reason     = "Unauthorized"
	defaultKey = "cc(!#)666"
)

var (
	ErrTokenInvalid     = errors.Unauthorized(reason, "Token is invalid")
	ErrTokenExpired     = errors.Unauthorized(reason, "Token has expired")
	ErrTokenParseFailed = errors.Unauthorized(reason, "Failed to parse token")
	ErrSigningMethod    = errors.Unauthorized(reason, "Wrong signing method")
	ErrSignTokenFailed  = errors.Unauthorized(reason, "Failed to sign token")
)

var (
	MessageTokenInvalid     = &i18n.Message{ID: "jwt.token.invalid", Other: ErrTokenInvalid.Error()}
	MessageTokenExpired     = &i18n.Message{ID: "jwt.token.expired", Other: ErrTokenExpired.Error()}
	MessageTokenParseFailed = &i18n.Message{ID: "jwt.token.parse.failed", Other: ErrTokenParseFailed.Error()}
	MessageSigningMethod    = &i18n.Message{ID: "jwt.wrong.signing.method", Other: ErrSigningMethod.Error()}
	MessageSignTokenFailed  = &i18n.Message{ID: "jwt.token.sign.failed", Other: ErrSignTokenFailed.Error()}
)

var defaultOptions = options{
	tokenType:     "Bearer",
	expired:       2 * time.Hour,
	signingMethod: jwt.SigningMethodHS256,
	signingKey:    []byte(defaultKey),
	keyFunc: func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrTokenInvalid
		}
		return []byte(defaultKey), nil
	},
}

type options struct {
	signingMethod jwt.SigningMethod
	signingKey    any
	keyFunc       jwt.Keyfunc
	issuer        string
	expired       time.Duration
	tokenType     string
	tokenHeader   map[string]any
}

type Option func(*options)

func WithSigningMethod(method jwt.SigningMethod) Option {
	return func(o *options) {
		o.signingMethod = method
	}
}

func WithSigningKey(key any) Option {
	return func(o *options) {
		o.signingKey = key
	}
}

func WithKeyFunc(keyFunc jwt.Keyfunc) Option {
	return func(o *options) {
		o.keyFunc = keyFunc
	}
}

func WithIssuer(issuer string) Option {
	return func(o *options) {
		o.issuer = issuer
	}
}

func WithExpired(expired time.Duration) Option {
	return func(o *options) {
		o.expired = expired
	}
}

func WithTokenHeader(header map[string]any) Option {
	return func(o *options) {
		o.tokenHeader = header
	}
}

type JWTAuth struct {
	opts *options
}

func New(opts ...Option) *JWTAuth {
	o := defaultOptions
	for _, opt := range opts {
		opt(&o)
	}
	return &JWTAuth{opts: &o}
}

func (a *JWTAuth) Sign(ctx context.Context, userId string) (authn.IToken, error) {
	now := time.Now()
	expiresAt := now.Add(a.opts.expired)

	token := jwt.NewWithClaims(a.opts.signingMethod, &jwt.RegisteredClaims{
		Issuer:    a.opts.issuer,                 // 令牌颁发者，表示该令牌是由谁创建的
		IssuedAt:  jwt.NewNumericDate(now),       // 令牌颁发时间
		ExpiresAt: jwt.NewNumericDate(expiresAt), // 令牌过期时间
		NotBefore: jwt.NewNumericDate(now),       // 令牌生效时间
		Subject:   userId,                        // 令牌主体，表示该令牌是关于谁的
	})

	if a.opts.tokenHeader != nil {
		for k, v := range a.opts.tokenHeader {
			token.Header[k] = v
		}
	}

	refreshToken, err := token.SignedString(a.opts.signingKey)
	if err != nil {
		return nil, ErrSignTokenFailed
	}

	tokenInfo := &tokenInfo{
		Token:     refreshToken,
		Type:      a.opts.tokenType,
		ExpiresAt: expiresAt.Unix(),
	}

	return tokenInfo, nil
}

func (a *JWTAuth) parseToken(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	token, err := jwt.ParseWithClaims(accessToken, &jwt.RegisteredClaims{}, a.opts.keyFunc)
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if !ok {
			return nil, errors.Unauthorized(reason, err.Error())
		}
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			return nil, ErrTokenInvalid
		}
		if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			return nil, ErrTokenExpired
		}
		return nil, ErrTokenParseFailed
	}

	if !token.Valid {
		return nil, ErrTokenInvalid
	}

	if token.Method != a.opts.signingMethod {
		return nil, ErrSigningMethod
	}

	return token.Claims.(*jwt.RegisteredClaims), nil
}

func (a *JWTAuth) Destroy(ctx context.Context, accessToken string) error {
	_, err := a.parseToken(ctx, accessToken)
	if err != nil {
		return err
	}

	// todo
	// 当存储了token时，将存储的token设置过期

	return nil
}

func (a *JWTAuth) ParseClaims(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error) {
	if accessToken == "" {
		return nil, ErrTokenInvalid
	}

	claims, err := a.parseToken(ctx, accessToken)
	if err != nil {
		return nil, err
	}

	return claims, nil
}

func (a *JWTAuth) Release() error {
	return nil
}
