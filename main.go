package main

import (
	"context"
	"flag"
	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/config"
	"log"
)

var nodeID = flag.String("node_id", "node1", "nodeID")
var raftDir = flag.String("raft_dir", "data/", "raft data dir")
var serverAddr = flag.String("server_addr", "127.0.0.1:3333", "listen addr")
var bootstrapCluster = flag.Bool("bootstrap_cluster", false, "bootstrap cluster")
var joinAddr = flag.String("join_addr", "", "join leader addr")
var prometheusAddr = flag.String("prometheus_address", ":9090", "prometheus address")

func main() {
	flag.Parse()
	ctx := context.Background()
	conf := &config.Config{
		NodeID:           *nodeID,
		RaftDir:          *raftDir,
		GrpcEndpoint:     *serverAddr,
		BootstrapCluster: *bootstrapCluster,
		JoinAddr:         *joinAddr,
		PrometheusAddr:   *prometheusAddr,
	}
	s, err := venus.NewServer(ctx, conf)
	if err != nil {
		log.Fatalf("venus.NewServer(%s): %v", *joinAddr, err)
	}

	err = s.Start()
	if err != nil {
		log.Fatalf("venus.VenusServer.Start(): %v", err)
	}
}
