package clientv1

import (
	"context"

	"go.uber.org/zap"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/no-mole/venus/proto/pbaccesskey"
	"google.golang.org/grpc"
)

type AccessKey interface {
	AccessKeyGen(ctx context.Context, alias string) (*pbaccesskey.AccessKeyInfo, error)
	AccessKeyDel(ctx context.Context, ak string) error
	AccessKeyChangeStatus(ctx context.Context, ak string, status pbaccesskey.AccessKeyStatus) error
	AccessKeyLogin(ctx context.Context, ak, secret string) (*pbaccesskey.AccessKeyLoginResponse, error)
	AccessKeyList(ctx context.Context) (*pbaccesskey.AccessKeyListResponse, error)
	AccessKeyAddNamespace(ctx context.Context, ak, namespace string) error
	AccessKeyDelNamespace(ctx context.Context, ak, namespace string) error
	AccessKeyNamespaceList(ctx context.Context, ak string) (*pbaccesskey.AccessKeyNamespaceListResponse, error)
}

func NewAccessKey(c *Client, logger *zap.Logger) AccessKey {
	return &accessKey{
		remote:   pbaccesskey.NewAccessKeyServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("access-key"),
	}
}

var _ AccessKey = &accessKey{}

type accessKey struct {
	remote pbaccesskey.AccessKeyServiceClient

	callOpts []grpc.CallOption

	logger *zap.Logger
}

func (a *accessKey) AccessKeyGen(ctx context.Context, alias string) (*pbaccesskey.AccessKeyInfo, error) {
	a.logger.Debug("AccessKeyGen", zap.String("alias", alias))
	return a.remote.AccessKeyGen(ctx, &pbaccesskey.AccessKeyInfo{
		Alias: alias,
	}, a.callOpts...)
}

func (a *accessKey) AccessKeyDel(ctx context.Context, ak string) error {
	a.logger.Debug("AccessKeyDel", zap.String("ak", ak))
	_, err := a.remote.AccessKeyDel(ctx, &pbaccesskey.AccessKeyDelRequest{Ak: ak}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyChangeStatus(ctx context.Context, ak string, status pbaccesskey.AccessKeyStatus) error {
	a.logger.Debug("AccessKeyChangeStatus", zap.String("ak", ak), zap.String("status", pbaccesskey.AccessKeyStatus_name[int32(status)]))
	_, err := a.remote.AccessKeyChangeStatus(ctx, &pbaccesskey.AccessKeyStatusChangeRequest{
		Ak:     ak,
		Status: status,
	}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyLogin(ctx context.Context, ak, secret string) (*pbaccesskey.AccessKeyLoginResponse, error) {
	a.logger.Debug("AccessKeyLogin", zap.String("ak", ak), zap.String("secret", secret))
	return a.remote.AccessKeyLogin(ctx, &pbaccesskey.AccessKeyLoginRequest{
		Ak:       ak,
		Password: secret,
	}, a.callOpts...)
}

func (a *accessKey) AccessKeyList(ctx context.Context) (*pbaccesskey.AccessKeyListResponse, error) {
	return a.remote.AccessKeyList(ctx, &emptypb.Empty{}, a.callOpts...)
}

func (a *accessKey) AccessKeyAddNamespace(ctx context.Context, ak, namespace string) error {
	a.logger.Debug("AccessKeyAddNamespace", zap.String("ak", ak), zap.String("namespace", namespace))
	_, err := a.remote.AccessKeyAddNamespace(ctx, &pbaccesskey.AccessKeyNamespaceInfo{
		Ak:        ak,
		Namespace: namespace,
	}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyDelNamespace(ctx context.Context, ak, namespace string) error {
	a.logger.Debug("AccessKeyDelNamespace", zap.String("ak", ak), zap.String("namespace", namespace))
	_, err := a.remote.AccessKeyDelNamespace(ctx, &pbaccesskey.AccessKeyNamespaceInfo{
		Ak:        ak,
		Namespace: namespace,
	}, a.callOpts...)
	return err
}

func (a *accessKey) AccessKeyNamespaceList(ctx context.Context, ak string) (*pbaccesskey.AccessKeyNamespaceListResponse, error) {
	return a.remote.AccessKeyNamespaceList(ctx, &pbaccesskey.AccessKeyNamespaceListRequest{Ak: ak}, a.callOpts...)
}
