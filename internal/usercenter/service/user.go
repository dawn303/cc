package service

import (
	"context"

	v1 "github.com/dawn303/cc/pkg/api/usercenter/v1"
	"github.com/dawn303/cc/pkg/log"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

func (s *UserCenterService) Login(ctx context.Context, rq *v1.LoginRequest) (*v1.LoginReply, error) {
	log.C(ctx).Infow("Login function called", "username", rq.Username)
	return s.biz.Auths().Login(ctx, rq)
}

func (s *UserCenterService) CreateUser(ctx context.Context, rq *v1.CreateUserRequest) (*v1.UserReply, error) {
	log.C(ctx).Infow("CreateUser function called", "username", rq.Username)
	return nil, nil
}

func (s *UserCenterService) ListUser(ctx context.Context, rq *v1.ListUserRequest) (*v1.ListUserResponse, error) {
	log.C(ctx).Infow("ListUser function called", "offset", rq.Offset, "limit", rq.Limit)
	return nil, nil
}

func (s *UserCenterService) DeleteUser(ctx context.Context, rq *v1.DeleteUserRequest) (*emptypb.Empty, error) {
	log.C(ctx).Infow("DeleteUser function called", "username", rq.Username)
	return &emptypb.Empty{}, nil
}
