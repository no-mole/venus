package clientv1

import (
	"context"

	"github.com/no-mole/venus/proto/pbnamespace"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"
)

type Namespace interface {
	NamespaceAdd(ctx context.Context, namespaceCN, namespaceEN string) (*pbnamespace.NamespaceItem, error)
	NamespaceDel(ctx context.Context, namespace string) error
	NamespacesList(ctx context.Context) (*pbnamespace.NamespacesListResponse, error)

	NamespaceAddUser(ctx context.Context, namespace, uid, role string) error
	NamespaceDelUser(ctx context.Context, namespace, uid string) error
	NamespaceUserList(ctx context.Context, namespace string) (*pbnamespace.NamespaceUserListResponse, error)

	NamespaceAddAccessKey(ctx context.Context, ak, namespace string) error
	NamespaceDelAccessKey(ctx context.Context, ak, namespace string) error
	NamespaceAccessKeyList(ctx context.Context, namespace string) (*pbnamespace.NamespaceAccessKeyListResponse, error)
}

func NewNamespace(c *Client, logger *zap.Logger) Namespace {
	return &namespace{
		remote:   pbnamespace.NewNamespaceServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("namespace"),
	}
}

var _ Namespace = &namespace{}

type namespace struct {
	remote   pbnamespace.NamespaceServiceClient
	callOpts []grpc.CallOption
	logger   *zap.Logger
}

func (n *namespace) NamespaceAddAccessKey(ctx context.Context, ak, namespace string) error {
	n.logger.Debug("NamespaceAddAccessKey", zap.String("ak", ak), zap.String("namespace", namespace))
	_, err := n.remote.NamespaceAddAccessKey(ctx, &pbnamespace.NamespaceAccessKeyInfo{
		Ak:        ak,
		Namespace: namespace,
	})
	return err
}

func (n *namespace) NamespaceDelAccessKey(ctx context.Context, ak, namespace string) error {
	n.logger.Debug("NamespaceDelAccessKey", zap.String("ak", ak), zap.String("namespace", namespace))
	_, err := n.remote.NamespaceDelAccessKey(ctx, &pbnamespace.NamespaceAccessKeyDelRequest{
		Ak:        ak,
		Namespace: namespace,
	})
	return err
}

func (n *namespace) NamespaceAccessKeyList(ctx context.Context, namespace string) (*pbnamespace.NamespaceAccessKeyListResponse, error) {
	return n.remote.NamespaceAccessKeyList(ctx, &pbnamespace.NamespaceAccessKeyListRequest{Namespace: namespace})
}

func (n *namespace) NamespaceAdd(ctx context.Context, namespaceCN, namespaceEN string) (*pbnamespace.NamespaceItem, error) {
	n.logger.Debug("NamespaceAdd", zap.String("namespaceCN", namespaceCN), zap.String("namespaceEN", namespaceEN))
	return n.remote.NamespaceAdd(ctx, &pbnamespace.NamespaceItem{
		NamespaceCn: namespaceCN,
		NamespaceEn: namespaceEN,
	}, n.callOpts...)
}

func (n *namespace) NamespaceDel(ctx context.Context, namespace string) error {
	n.logger.Debug("NamespaceDel", zap.String("namespace", namespace))
	_, err := n.remote.NamespaceDel(ctx, &pbnamespace.NamespaceDelRequest{
		Namespace: namespace,
	}, n.callOpts...)
	return err
}

func (n *namespace) NamespacesList(ctx context.Context) (*pbnamespace.NamespacesListResponse, error) {
	return n.remote.NamespacesList(ctx, &emptypb.Empty{}, n.callOpts...)
}

func (n *namespace) NamespaceAddUser(ctx context.Context, namespace, uid, role string) error {
	n.logger.Debug("NamespaceAddUser", zap.String("namespace", namespace), zap.String("uid", uid), zap.String("role", role))
	_, err := n.remote.NamespaceAddUser(ctx, &pbnamespace.NamespaceUserInfo{
		Namespace: namespace,
		Uid:       uid,
		Role:      role,
	}, n.callOpts...)
	return err
}

func (n *namespace) NamespaceDelUser(ctx context.Context, namespace, uid string) error {
	n.logger.Debug("NamespaceDelUser", zap.String("namespace", namespace), zap.String("uid", uid))
	_, err := n.remote.NamespaceDelUser(ctx, &pbnamespace.NamespaceUserDelRequest{
		Namespace: namespace,
		Uid:       uid,
	}, n.callOpts...)
	return err
}

func (n *namespace) NamespaceUserList(ctx context.Context, namespace string) (*pbnamespace.NamespaceUserListResponse, error) {
	return n.remote.NamespaceUserList(ctx, &pbnamespace.NamespaceUserListRequest{
		Namespace: namespace,
	}, n.callOpts...)
}
