package namespace

import (
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/grpc"
)

type namespaceService struct {
	pbnamespace.UnimplementedNamespaceServer

	raft *raft.Raft
	fsm  *fsm.FSM
}

var (
	bucketName = []byte("namespace")
)

func New(raft *raft.Raft, fsm *fsm.FSM) (desc *grpc.ServiceDesc, impl interface{}) {
	return &pbnamespace.Namespace_ServiceDesc, &namespaceService{
		raft: raft,
		fsm:  fsm,
	}
}
