package auth

import (
	"context"
	"fmt"

	"github.com/dawn303/cc/internal/usercenter/store"
	v1 "github.com/dawn303/cc/pkg/api/usercenter/v1"
	"github.com/dawn303/cc/pkg/authn"
)

type AuthBiz interface {
	Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error)
}

type authBiz struct {
	ds    store.IStore
	authn authn.Authenticator
}

func New(ds store.IStore, authn authn.Authenticator) AuthBiz {
	return &authBiz{ds: ds, authn: authn}
}

func (b *authBiz) Login(ctx context.Context, req *v1.LoginRequest) (*v1.LoginReply, error) {

	if req.Username != "admin" {
		return nil, fmt.Errorf("user not found")
	}

	if err := authn.Compare("$2a$10$eXgOsRqZq8YPYLKt5.YFOuMpjvs4Y0pF7d83/U3r6RmNcAoz65732", req.Password); err != nil {
		return nil, fmt.Errorf("password does not match")
	}

	accessToken, err := b.authn.Sign(ctx, req.Username)
	if err != nil {
		return nil, err
	}

	return &v1.LoginReply{
		AccessToken: accessToken.GetToken(),
		Type:        accessToken.GetTokenType(),
		ExpiresAt:   accessToken.GetExpiresAt(),
	}, nil
}
