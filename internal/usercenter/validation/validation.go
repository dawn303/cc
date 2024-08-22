package validation

import (
	"github.com/dawn303/cc/internal/usercenter/store"
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(New, wire.Bind(new(any), new(*validator)))

type validator struct {
	ds store.IStore
}

func New(ds store.IStore) (*validator, error) {
	vd := &validator{ds: ds}

	return vd, nil
}
