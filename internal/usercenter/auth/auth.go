package auth

import (
	"github.com/google/wire"
)

var ProviderSet = wire.NewSet(NewAuth, wire.Bind(new(AuthProvider), new(*auth)))

type AuthProvider interface {
}

type auth struct {
}

func NewAuth() *auth {
	return &auth{}
}
