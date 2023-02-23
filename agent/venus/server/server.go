package server

import (
	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/no-mole/venus/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbuser"
)

type Server interface {
	pbkv.KVServiceServer
	pblease.LeaseServiceServer
	pbnamespace.NamespaceServiceServer
	pbuser.UserServiceServer
	pbmicroservice.MicroServiceServer
	pbcluster.ClusterServiceServer
	pbaccesskey.AccessKeyServiceServer
}
