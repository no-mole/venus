package namespace

import (
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/state"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/grpc"
)

type namespaceService struct {
	pbnamespace.UnimplementedNamespaceServer

	raft  *raft.Raft
	state *state.State
}

var (
	bucketName = []byte("namespace")
)

func New(raft *raft.Raft, state *state.State) (desc *grpc.ServiceDesc, impl interface{}) {
	return &pbnamespace.Namespace_ServiceDesc, &namespaceService{
		raft:  raft,
		state: state,
	}
}
