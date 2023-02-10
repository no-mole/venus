package proxy

import (
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/internal/proto/pbraftadmin"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbservice"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/grpc"
)

type Remote struct {
	pbkv.UnimplementedKVServer
	pbnamespace.UnimplementedNamespaceServiceServer
	pblease.UnimplementedLeaseServiceServer
	pbservice.UnimplementedServiceServer
	pbuser.UnimplementedUserServiceServer
	pbraftadmin.UnimplementedRaftAdminServer

	cc *grpc.ClientConn
}

func NewRemoteServer(cc *grpc.ClientConn) server.Server {
	return &Remote{cc: cc}
}

func (s *Remote) getActiveConn() *grpc.ClientConn {
	return s.cc
}
