package clientv1

import (
	"context"

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
	UserAddNamespace(ctx context.Context, uid, namespace, role string) error
	UserDelNamespace(ctx context.Context, uid, namespace string) error
	UserNamespaceList(ctx context.Context, uid string) (*pbuser.UserNamespaceListResponse, error)
}

func NewUser(c *Client) User {
	return &user{
		remote:   pbuser.NewUserServiceClient(c.conn),
		callOpts: c.callOpts,
	}
}

var _ User = &user{}

type user struct {
	remote   pbuser.UserServiceClient
	callOpts []grpc.CallOption
}

func (u *user) UserRegister(ctx context.Context, uid, name, password string) (*pbuser.UserInfo, error) {
	return u.remote.UserRegister(ctx, &pbuser.UserInfo{
		Uid:      uid,
		Name:     name,
		Password: password,
	}, u.callOpts...)
}

func (u *user) UserUnregister(ctx context.Context, uid string) (*pbuser.UserInfo, error) {
	return u.remote.UserUnregister(ctx, &pbuser.UserInfo{
		Uid: uid,
	}, u.callOpts...)
}

func (u *user) UserLogin(ctx context.Context, uid, password string) (*pbuser.LoginResponse, error) {
	return u.remote.UserLogin(ctx, &pbuser.LoginRequest{
		Uid:      uid,
		Password: password,
	}, u.callOpts...)
}

func (u *user) UserChangeStatus(ctx context.Context, uid string, status pbuser.UserStatus) error {
	_, err := u.remote.UserChangeStatus(ctx, &pbuser.ChangeUserStatusRequest{
		Uid:    uid,
		Status: status,
	}, u.callOpts...)
	return err
}

func (u *user) UserList(ctx context.Context) (*pbuser.UserListResponse, error) {
	return u.remote.UserList(ctx, &emptypb.Empty{}, u.callOpts...)
}

func (u *user) UserAddNamespace(ctx context.Context, uid, namespace, role string) error {
	_, err := u.remote.UserAddNamespace(ctx, &pbuser.UserNamespaceInfo{
		Uid:       uid,
		Namespace: namespace,
		Role:      role,
	}, u.callOpts...)
	return err
}

func (u *user) UserDelNamespace(ctx context.Context, uid, namespace string) error {
	_, err := u.remote.UserDelNamespace(ctx, &pbuser.UserNamespaceInfo{
		Uid:       uid,
		Namespace: namespace,
	}, u.callOpts...)
	return err
}

func (u *user) UserNamespaceList(ctx context.Context, uid string) (*pbuser.UserNamespaceListResponse, error) {
	return u.remote.UserNamespaceList(ctx, &pbuser.UserNamespaceListRequest{Uid: uid}, u.callOpts...)
}
