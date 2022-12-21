package main

import (
	"context"
	"flag"
	"github.com/Jille/raftadmin/proto"
	"github.com/no-mole/venus/agent/venus"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/service/namespace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var nodeID = flag.String("node_id", "node1", "nodeID")
var raftDir = flag.String("raft_dir", "data/", "raft data dir")
var serverAddr = flag.String("server_addr", "127.0.0.1:3333", "listen addr")
var bootstrapCluster = flag.Bool("bootstrap_cluster", false, "bootstrap cluster")
var joinAddr = flag.String("join_addr", "", "join leader addr")

func main() {
	flag.Parse()
	ctx := context.Background()
	conf := &config.Config{
		NodeID:           *nodeID,
		RaftDir:          *raftDir,
		ServerAddr:       *serverAddr,
		BootstrapCluster: *bootstrapCluster,
	}
	s, err := venus.NewServer(ctx, conf, nil)
	if err != nil {
		log.Fatalf("venus.NewServer(%s): %v", *joinAddr, err)
	}

	if *joinAddr != "" {
		conn, err := grpc.Dial(*joinAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed grpc.Dial(%s): %v", *joinAddr, err)
		}
		client := proto.NewRaftAdminClient(conn)
		_, err = client.AddVoter(ctx, &proto.AddVoterRequest{
			Id:            *nodeID,
			Address:       *serverAddr,
			PreviousIndex: s.Raft.LastIndex(),
		})
		if err != nil {
			log.Fatalf("failed client.AddVoter(%s): %v", *joinAddr, err)
		}
	}

	err = s.RegisterServices(
		namespace.New,
	)
	if err != nil {
		log.Fatalf("venus.Server.RegisterServices(): %v", err)
	}

	err = s.Start()
	if err != nil {
		log.Fatalf("venus.Server.Start(): %v", err)
	}
}
