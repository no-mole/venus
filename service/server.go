package service

import (
	"context"
	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/config"
)

var server *venus.Server

func InitServer(ctx context.Context, nodeID, raftDir, serverAddr string, bootstrapCluster bool) (err error) {
	server, err = venus.NewServer(ctx, &config.Config{
		NodeID:           nodeID,
		RaftDir:          raftDir,
		ServerAddr:       serverAddr,
		BootstrapCluster: bootstrapCluster,
	})
	return err
}

func Server() *venus.Server {
	return server
}
