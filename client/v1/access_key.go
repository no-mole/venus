package clientv1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/no-mole/venus/proto/pbaccesskey"
	"google.golang.org/grpc"
)

type AccessKey interface {
	AccessKeyGen(ctx context.Context, ak, alias string) (*pbaccesskey.AccessKeyInfo, error)
	AccessKeyDel(ctx context.Context, ak string) error
	AccessKeyChangeStatus(ctx context.Context, ak string, status pbaccesskey.AccessKeyStatus) error
	AccessKeyLogin(ctx context.Context, ak, secret string) (*pbaccesskey.AccessKeyLoginResponse, error)
	AccessKeyList(ctx context.Context) (*pbaccesskey.AccessKeyListResponse, error)
	AccessKeyAddNamespace(ctx context.Context, ak, namespace string) error
	AccessKeyDelNamespace(ctx context.Context, ak, namespace string) error
	AccessKeyNamespaceList(ctx context.Context, ak string) (*pbaccesskey.AccessKeyNamespaceListResponse, error)
}

func NewAccessKey(c *Client) AccessKey {
	return &accessKey{
		remote:   pbaccesskey.NewAccessKeyServiceClient(c.conn),
		callOpts: c.callOpts,
	}
}

var _ AccessKey = &accessKey{}

type accessKey struct {
	remote   pbaccesskey.AccessKeyServiceClient
	callOpts []grpc.CallOption
}

func (a *accessKey) AccessKeyGen(ctx context.Context, ak, alias string) (*pbaccesskey.AccessKeyInfo, error) {
	return a.remote.AccessKeyGen(ctx, &pbaccesskey.AccessKeyInfo{
		Ak:    ak,
		Alias: alias,
	}, a.callOpts...)
}

func (a *accessKey) AccessKeyDel(ctx context.Context, ak string) error {
	_, err := a.remote.AccessKeyDel(ctx, &pbaccesskey.AccessKeyInfo{Ak: ak}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyChangeStatus(ctx context.Context, ak string, status pbaccesskey.AccessKeyStatus) error {
	_, err := a.remote.AccessKeyChangeStatus(ctx, &pbaccesskey.AccessKeyStatusChangeRequest{
		Ak:     ak,
		Status: status,
	}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyLogin(ctx context.Context, ak, secret string) (*pbaccesskey.AccessKeyLoginResponse, error) {
	return a.remote.AccessKeyLogin(ctx, &pbaccesskey.AccessKeyLoginRequest{
		Ak:       ak,
		Password: secret,
	}, a.callOpts...)
}

func (a *accessKey) AccessKeyList(ctx context.Context) (*pbaccesskey.AccessKeyListResponse, error) {
	return a.remote.AccessKeyList(ctx, &emptypb.Empty{}, a.callOpts...)
}

func (a *accessKey) AccessKeyAddNamespace(ctx context.Context, ak, namespace string) error {
	_, err := a.remote.AccessKeyAddNamespace(ctx, &pbaccesskey.AccessKeyNamespaceInfo{
		Ak:        ak,
		Namespace: namespace,
	}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyDelNamespace(ctx context.Context, ak, namespace string) error {
	_, err := a.remote.AccessKeyDelNamespace(ctx, &pbaccesskey.AccessKeyNamespaceInfo{
		Ak:        ak,
		Namespace: namespace,
	}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyNamespaceList(ctx context.Context, ak string) (*pbaccesskey.AccessKeyNamespaceListResponse, error) {
	return a.remote.AccessKeyNamespaceList(ctx, &pbaccesskey.AccessKeyNamespaceListRequest{Ak: ak}, a.callOpts...)
}
