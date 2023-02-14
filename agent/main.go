package main

import (
	"context"
	"flag"
	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/config"
)

var nodeID = flag.String("node_id", "node1", "nodeID")
var raftDir = flag.String("raft_dir", "data/", "raft data dir")
var serverAddr = flag.String("server_addr", "127.0.0.1:3333", "listen addr")
var bootstrapCluster = flag.Bool("bootstrap_cluster", false, "bootstrap cluster")
var joinAddr = flag.String("join_addr", "", "join leader addr")
var logLevel = flag.String("log_level", "info", "logger level[debug|info|warn|err]")

func main() {
	flag.Parse()
	ctx := context.Background()
	conf := config.GetDefaultConfig()
	conf.NodeID = *nodeID
	conf.RaftDir = *raftDir
	conf.GrpcEndpoint = *serverAddr
	conf.BootstrapCluster = *bootstrapCluster
	conf.JoinAddr = *joinAddr
	conf.LoggerLevel = config.LoggerLevel(*logLevel)
	s, err := venus.NewServer(ctx, conf)
	if err != nil {
		panic(err)
	}

	err = s.Start()
	if err != nil {
		panic(err)
	}
}
