package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/Jille/raftadmin"
	"github.com/Jille/raftadmin/proto"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/service"
	"github.com/no-mole/venus/service/namespace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"time"
)

var nodeID = flag.String("node_id", "", "nodeID")
var raftDir = flag.String("raft_dir", "data/", "raft data dir")
var serverAddr = flag.String("server_addr", "127.0.0.1:3333", "listen addr")
var bootstrapCluster = flag.Bool("bootstrap_cluster", false, "bootstrap cluster")
var joinAddr = flag.String("join_addr", "", "join leader addr")

func main() {
	flag.Parse()
	ctx := context.Background()

	_, port, err := net.SplitHostPort(*serverAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sock, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	err = service.InitServer(ctx, *nodeID, *raftDir, *serverAddr, *bootstrapCluster)
	if err != nil {
		log.Fatalf("failed to InitServer: %v", err)
	}
	server := service.Server()

	if *joinAddr != "" {
		conn, err := grpc.Dial(*joinAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed grpc.Dial(%s): %v", *joinAddr, err)
		}
		client := proto.NewRaftAdminClient(conn)
		_, err = client.AddVoter(ctx, &proto.AddVoterRequest{
			Id:            *nodeID,
			Address:       *serverAddr,
			PreviousIndex: server.Raft.LastIndex(),
		})
		if err != nil {
			log.Fatalf("failed client.AddVoter(%s): %v", *joinAddr, err)
		}
	}

	s := grpc.NewServer()
	pbnamespace.RegisterNamespaceServer(s, namespace.New())
	raftadmin.Register(s, server.Raft) //raft 管理 grpc
	reflection.Register(s)
	server.Transport.Register(s)
	//leaderhealth.Setup(r, s, []string{"Example"})
	go func() {
		c := time.NewTicker(time.Second * 5)
		for {
			<-c.C
			fmt.Printf("%+v\n\n", server.Raft.GetConfiguration().Configuration().Servers)
		}
	}()
	if err := s.Serve(sock); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
