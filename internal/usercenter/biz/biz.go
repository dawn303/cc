package biz

import (
	"github.com/google/wire"

	"github.com/dawn303/cc/internal/usercenter/store"
)

var ProviderSet = wire.NewSet(NewBiz, wire.Bind(new(IBiz), new(*biz)))

type IBiz interface {
}

type biz struct {
	ds store.IStore
}

func NewBiz(ds store.IStore) *biz {
	return &biz{ds: ds}
}
