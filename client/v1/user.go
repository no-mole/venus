package clientv1

import (
	"context"

	"github.com/no-mole/venus/proto/pbnamespace"

	"go.uber.org/zap"

	"google.golang.org/protobuf/types/known/emptypb"

	"google.golang.org/grpc"

	"github.com/no-mole/venus/proto/pbuser"
)

type User interface {
	UserRegister(ctx context.Context, uid, name, password string) (*pbuser.UserInfo, error)
	UserUnregister(ctx context.Context, uid string) (*pbuser.UserInfo, error)
	UserLogin(ctx context.Context, uid, password string) (*pbuser.LoginResponse, error)
	UserChangeStatus(ctx context.Context, uid string, status pbuser.UserStatus) error
	UserList(ctx context.Context) (*pbuser.UserListResponse, error)
	UserNamespaceList(ctx context.Context, uid string) (*pbnamespace.NamespaceUserListResponse, error)
}

func NewUser(c *Client, logger *zap.Logger) User {
	return &user{
		remote:   pbuser.NewUserServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("user"),
	}
}

var _ User = &user{}

type user struct {
	remote   pbuser.UserServiceClient
	callOpts []grpc.CallOption
	logger   *zap.Logger
}

func (u *user) UserRegister(ctx context.Context, uid, name, password string) (*pbuser.UserInfo, error) {
	u.logger.Debug("UserRegister", zap.String("uid", uid), zap.String("name", name), zap.String("password", password))
	return u.remote.UserRegister(ctx, &pbuser.UserInfo{
		Uid:      uid,
		Name:     name,
		Password: password,
	}, u.callOpts...)
}

func (u *user) UserUnregister(ctx context.Context, uid string) (*pbuser.UserInfo, error) {
	u.logger.Debug("UserUnregister", zap.String("uid", uid))
	return u.remote.UserUnregister(ctx, &pbuser.UserInfo{
		Uid: uid,
	}, u.callOpts...)
}

func (u *user) UserLogin(ctx context.Context, uid, password string) (*pbuser.LoginResponse, error) {
	u.logger.Debug("UserLogin", zap.String("uid", uid), zap.String("password", password))
	return u.remote.UserLogin(ctx, &pbuser.LoginRequest{
		Uid:      uid,
		Password: password,
	}, u.callOpts...)
}

func (u *user) UserChangeStatus(ctx context.Context, uid string, status pbuser.UserStatus) error {
	u.logger.Debug("UserChangeStatus", zap.String("uid", uid), zap.String("status", pbuser.UserStatus_name[int32(status)]))
	_, err := u.remote.UserChangeStatus(ctx, &pbuser.ChangeUserStatusRequest{
		Uid:    uid,
		Status: status,
	}, u.callOpts...)
	return err
}

func (u *user) UserList(ctx context.Context) (*pbuser.UserListResponse, error) {
	return u.remote.UserList(ctx, &emptypb.Empty{}, u.callOpts...)
}

func (u *user) UserNamespaceList(ctx context.Context, uid string) (*pbnamespace.NamespaceUserListResponse, error) {
	return u.remote.UserNamespaceList(ctx, &pbuser.UserNamespaceListRequest{Uid: uid}, u.callOpts...)
}
