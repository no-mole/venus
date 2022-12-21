package namespace

import (
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/proto/pbnamespace"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/grpc"
)

type namespaceService struct {
	pbnamespace.UnimplementedNamespaceServer

	raft *raft.Raft
	db   *bolt.DB
}

var (
	bucketName = []byte("namespace")
)

func New(raft *raft.Raft, db *bolt.DB) (desc *grpc.ServiceDesc, impl interface{}) {
	return &pbnamespace.Namespace_ServiceDesc, &namespaceService{
		raft: raft,
		db:   db,
	}
}
