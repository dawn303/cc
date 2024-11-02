package biz

import (
	"github.com/google/wire"

	"github.com/dawn303/cc/internal/usercenter/biz/auth"
	"github.com/dawn303/cc/internal/usercenter/store"
	"github.com/dawn303/cc/pkg/authn"
)

var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface {
	Auths() auth.AuthBiz
}

type biz struct {
	ds    store.IStore
	authn authn.Authenticator
}

func NewBiz(ds store.IStore, authn authn.Authenticator) *biz {
	return &biz{ds: ds, authn: authn}
}

func (b *biz) Auths() auth.AuthBiz {
	return auth.New(b.ds, b.authn)
}
