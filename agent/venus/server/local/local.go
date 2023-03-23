package local

import (
	"math/rand"
	"time"

	"github.com/no-mole/venus/proto/pbsysconfig"

	"github.com/bwmarrin/snowflake"
	"github.com/hashicorp/raft"
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
	pbkv.KVServiceServer
	pbnamespace.NamespaceServiceServer
	pblease.LeaseServiceServer
	pbmicroservice.MicroServiceServer
	pbuser.UserServiceServer
	pbcluster.ClusterServiceServer
	pbaccesskey.AccessKeyServiceServer
	pbsysconfig.SysConfigServiceServer

	r *raft.Raft

	fsm *fsm.FSM

	snowflakeNode *snowflake.Node

	applyTimeout time.Duration
}

func NewLocalServer(r *raft.Raft, fsm *fsm.FSM, applyTimeout time.Duration) server.Server {
	if applyTimeout == 0 {
		applyTimeout = time.Second
	}
	snowflakeNode, _ := snowflake.NewNode(int64(rand.Intn(1023)))
	return &Local{
		r:             r,
		fsm:           fsm,
		snowflakeNode: snowflakeNode,
		applyTimeout:  applyTimeout,
	}
}
