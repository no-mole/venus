package clientv1

import (
	"context"
	"github.com/no-mole/venus/proto/pbnamespace"
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

func NewNamespace(c *Client) Namespace {
	return &namespace{
		remote:   pbnamespace.NewNamespaceServiceClient(c.conn),
		callOpts: c.callOpts,
	}
}

var _ Namespace = &namespace{}

type namespace struct {
	remote   pbnamespace.NamespaceServiceClient
	callOpts []grpc.CallOption
}

func (n *namespace) NamespaceAddAccessKey(ctx context.Context, ak, namespace string) error {
	//TODO implement me
	panic("implement me")
}

func (n *namespace) NamespaceDelAccessKey(ctx context.Context, ak, namespace string) error {
	//TODO implement me
	panic("implement me")
}

func (n *namespace) NamespaceAccessKeyList(ctx context.Context, namespace string) (*pbnamespace.NamespaceAccessKeyListResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (n *namespace) NamespaceAdd(ctx context.Context, namespaceCN, namespaceEN string) (*pbnamespace.NamespaceItem, error) {
	return n.remote.NamespaceAdd(ctx, &pbnamespace.NamespaceItem{
		NamespaceCn: namespaceCN,
		NamespaceEn: namespaceEN,
	}, n.callOpts...)
}

func (n *namespace) NamespaceDel(ctx context.Context, namespace string) error {
	_, err := n.remote.NamespaceDel(ctx, &pbnamespace.NamespaceDelRequest{
		Namespace: namespace,
	}, n.callOpts...)
	return err
}

func (n *namespace) NamespacesList(ctx context.Context) (*pbnamespace.NamespacesListResponse, error) {
	return n.remote.NamespacesList(ctx, &emptypb.Empty{}, n.callOpts...)
}

func (n *namespace) NamespaceAddUser(ctx context.Context, namespace, uid, role string) error {
	_, err := n.remote.NamespaceAddUser(ctx, &pbnamespace.NamespaceUserInfo{
		Namespace: namespace,
		Uid:       uid,
		Role:      role,
	}, n.callOpts...)
	return err
}

func (n *namespace) NamespaceDelUser(ctx context.Context, namespace, uid string) error {
	_, err := n.remote.NamespaceDelUser(ctx, &pbnamespace.NamespaceUserInfo{
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
