package proxy

import (
	"github.com/no-mole/venus/agent/venus/server"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbaccesskey"
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
	pbaccesskey.AccessKeyServiceServer

	cc     *grpc.ClientConn
	client *clientv1.Client
}

func NewRemoteServer(cc *grpc.ClientConn, client *clientv1.Client) server.Server {
	return &Remote{cc: cc, client: client}
}

func (s *Remote) getActiveConn() *grpc.ClientConn {
	return s.cc
}
