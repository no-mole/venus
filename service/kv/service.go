package kv

import (
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/grpc"
)

type kvService struct {
	pbkv.UnimplementedKVServer

	raft *raft.Raft
	fsm  *fsm.FSM
}

var (
	bucketNamePrefix = "kv_"
)

func genBucketName(namespace string) []byte {
	return []byte(bucketNamePrefix + namespace)
}

func New(raft *raft.Raft, fsm *fsm.FSM) (desc *grpc.ServiceDesc, impl interface{}) {
	return &pbkv.KV_ServiceDesc, &kvService{
		raft: raft,
		fsm:  fsm,
	}
}
