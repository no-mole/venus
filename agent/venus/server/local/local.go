package local

import (
	"math/rand"

	"github.com/bwmarrin/snowflake"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/no-mole/venus/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbuser"
)

type Local struct {
	r *raft.Raft

	fsm *fsm.FSM

	config *config.Config

	lessor *lessor

	snowflakeNode *snowflake.Node

	pbkv.KVServiceServer
	pbnamespace.NamespaceServiceServer
	pblease.LeaseServiceServer
	pbmicroservice.MicroServiceServer
	pbuser.UserServiceServer
	pbcluster.ClusterServiceServer
	pbaccesskey.AccessKeyServiceServer
}

func NewLocalServer(r *raft.Raft, fsm *fsm.FSM, conf *config.Config) server.Server {
	snowflakeNode, _ := snowflake.NewNode(int64(rand.Intn(1023)))
	return &Local{
		r:      r,
		fsm:    fsm,
		config: conf,
		lessor: &lessor{ //todo new lessor
			leases: map[int64]*Lease{},
		},
		snowflakeNode: snowflakeNode,
	}
}
