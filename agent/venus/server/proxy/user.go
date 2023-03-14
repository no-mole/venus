package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (s *Remote) UserRegister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	return s.client.UserRegister(ctx, info.Uid, info.Name, info.Password)
}

func (s *Remote) UserUnregister(ctx context.Context, info *pbuser.UserInfo) (*pbuser.UserInfo, error) {
	return s.client.UserUnregister(ctx, info.Uid)
}

func (s *Remote) UserChangeStatus(ctx context.Context, req *pbuser.ChangeUserStatusRequest) (*emptypb.Empty, error) {
	err := s.client.UserChangeStatus(ctx, req.Uid, req.Status)
	return &emptypb.Empty{}, err
}
