package venus

import (
	"context"
	"fmt"
	"github.com/no-mole/venus/proto/pbservice"
	"log"
	"math/rand"
	"net"
	"os"
	"path/filepath"
	"time"

	transport "github.com/Jille/raft-grpc-transport"
	"github.com/bwmarrin/snowflake"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/state"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	raftLogCacheSize = 512
)

type Server struct {
	pbkv.UnimplementedKVServer
	pbnamespace.UnimplementedNamespaceServer
	pblease.UnimplementedLeaseServiceServer
	pbservice.UnimplementedServiceServer

	fsm *fsm.FSM

	Raft *raft.Raft

	state *state.State

	stable *boltdb.BoltStore

	grpcServer *grpc.Server

	sock net.Listener

	transport *transport.Manager

	config *config.Config

	lessor *lessor

	snowflakeNode *snowflake.Node
}

func NewServer(ctx context.Context, config *config.Config, grpcOpts []grpc.ServerOption) (_ *Server, err error) {
	if config.ApplyTimeout == 0 {
		config.ApplyTimeout = 5 * time.Second
	}
	s := &Server{
		config:     config,
		grpcServer: grpc.NewServer(grpcOpts...),
		lessor: &lessor{ //todo new lessor
			leases: map[int64]*Lease{},
		},
	}

	s.snowflakeNode, err = snowflake.NewNode(int64(rand.Intn(1023)))
	if err != nil {
		return nil, fmt.Errorf(`snowflake.NewNode: %v`, err.Error())
	}

	baseDir := filepath.Join(config.RaftDir, config.NodeID)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf(`os.MkdirAll(%q): %v`, baseDir, err)
	}

	dbPath := fmt.Sprintf("%s/data.dat", baseDir)
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		return nil, fmt.Errorf(`Raft.bolt.Open(%q, ...): %v`, dbPath, err)
	}

	s.state = state.New(ctx, db)

	boltFSM, err := fsm.NewBoltFSM(ctx, s.state)
	if err != nil {
		return nil, fmt.Errorf(`fsm.NewBoltFSM(%q, ...): %v`, baseDir, err)
	}

	s.fsm = boltFSM

	//@todo consul 用法
	//raftLayer := consul.NewRaftLayer()
	//
	//// Create a transport layer.
	//transConfig := &Raft.NetworkTransportConfig{
	//	Stream:                raftLayer,
	//	MaxPool:               3,
	//	Timeout:               10 * time.Second,
	//	ServerAddressProvider: serverAddressProvider,
	//}
	//
	//trans := Raft.NewNetworkTransportWithConfig(transConfig)

	tm := transport.New(raft.ServerAddress(config.ServerAddr), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})
	s.transport = tm

	stable, err := boltdb.NewBoltStore(filepath.Join(baseDir, "stable.dat"))
	if err != nil {
		return nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "stable.dat"), err)
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raftLogCacheSize, stable)
	if err != nil {
		return nil, fmt.Errorf("Raft.NewLogCache():%v", err)
	}
	logStore := cacheStore

	s.stable = stable

	snap, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf(`Raft.NewFileSnapshotStore(%q, ...): %v`, baseDir, err)
	}

	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(config.NodeID)
	r, err := raft.NewRaft(c, boltFSM, logStore, stable, snap, tm.Transport())
	if err != nil {
		return nil, fmt.Errorf("Raft.NewRaft: %v", err)
	}
	s.Raft = r

	if config.BootstrapCluster {
		cfg := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       raft.ServerID(config.NodeID),
					Address:  raft.ServerAddress(config.ServerAddr),
				},
			},
		}
		f := r.BootstrapCluster(cfg)
		err := f.Error()
		if err != nil {
			if err == raft.ErrCantBootstrap {
				log.Printf("Raft.BootstrapCluster: %s", err.Error())
			} else {
				return nil, err
			}
		}
	}
	return s, nil
}

func (s *Server) Start() error {
	for _, desc := range []*grpc.ServiceDesc{
		&pbnamespace.Namespace_ServiceDesc,
		&pbkv.KV_ServiceDesc,
		&pblease.LeaseService_ServiceDesc,
		&pbservice.Service_ServiceDesc,
	} {
		s.grpcServer.RegisterService(desc, s)
	}
	Register(s.grpcServer, s.Raft) //Raft 管理 grpc
	s.transport.Register(s.grpcServer)
	//reflection.Register(s.grpcServer)

	_, port, err := net.SplitHostPort(s.config.ServerAddr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	sock, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.sock = sock
	err = s.grpcServer.Serve(s.sock)
	return err
}

func (s *Server) runLoop() {
	for {
		switch s.Raft.State() {
		case raft.Leader:
			s.runLeader()
		default:
			<-s.Raft.LeaderCh()
		}
	}
}

func (s *Server) runLeader() {
	//run leader loop
	//TODO load leases
	//TODO check leases ddl
	<-s.Raft.LeaderCh()
	//TODO clear
}
