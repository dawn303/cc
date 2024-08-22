package service

import (
	"github.com/google/wire"

	"github.com/dawn303/cc/internal/usercenter/biz"
	v1 "github.com/dawn303/cc/pkg/api/usercenter/v1"
)

var ProviderSet = wire.NewSet(NewUserCenterService)

type UserCenterService struct {
	v1.UnimplementedUserCenterServer
	biz biz.IBiz
}

func NewUserCenterService(biz biz.IBiz) *UserCenterService {
	return &UserCenterService{biz: biz}
}
