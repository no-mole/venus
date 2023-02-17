package local

import (
	"github.com/bwmarrin/snowflake"
	"github.com/hashicorp/raft"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbservice"
	"github.com/no-mole/venus/proto/pbuser"
	"math/rand"
)

type Local struct {
	r *raft.Raft

	fsm *fsm.FSM

	config *config.Config

	lessor *lessor

	snowflakeNode *snowflake.Node

	pbkv.UnimplementedKVServer
	pbnamespace.UnimplementedNamespaceServiceServer
	pblease.UnimplementedLeaseServiceServer
	pbservice.UnimplementedServiceServer
	pbuser.UnimplementedUserServiceServer
	pbcluster.UnimplementedClusterServer
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
