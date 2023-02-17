package proxy

import (
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/grpc"
)

type Remote struct {
	pbkv.KVServiceServer
	pbnamespace.NamespaceServiceServer
	pblease.LeaseServiceServer
	pbmicroservice.MicroServiceServer
	pbuser.UserServiceServer
	pbcluster.ClusterServer

	cc *grpc.ClientConn
}

func NewRemoteServer(cc *grpc.ClientConn) server.Server {
	return &Remote{cc: cc}
}

func (s *Remote) getActiveConn() *grpc.ClientConn {
	return s.cc
}
