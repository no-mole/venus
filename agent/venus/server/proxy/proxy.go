package proxy

import (
	"github.com/no-mole/venus/agent/venus/server"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/no-mole/venus/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbconfig"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbuser"
)

type Remote struct {
	pbkv.KVServiceServer
	pbnamespace.NamespaceServiceServer
	pblease.LeaseServiceServer
	pbmicroservice.MicroServiceServer
	pbuser.UserServiceServer
	pbcluster.ClusterServiceServer
	pbaccesskey.AccessKeyServiceServer
	pbconfig.ConfigServiceServer

	client *clientv1.Client
}

func NewRemoteServer(client *clientv1.Client) server.Server {
	return &Remote{client: client}
}
