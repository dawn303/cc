package authn

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type IToken interface {
	GetToken() string
	GetTokenType() string
	GetExpiresAt() int64
	EncodeToJSON() ([]byte, error)
}

type Authenticator interface {
	Sign(ctx context.Context, userId string) (IToken, error)
	Destroy(ctx context.Context, accessToken string) error
	ParseClaims(ctx context.Context, accessToken string) (*jwt.RegisteredClaims, error)
	Release() error
}

func Encrypt(source string) (string, error) {
	hashBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashBytes), err
}

func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
