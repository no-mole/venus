package venus

import (
	"context"
	"fmt"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/Jille/raftadmin"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/state"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
	"os"
	"path/filepath"
)

const (
	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	raftLogCacheSize = 512
)

type Server struct {
	fsm *fsm.FSM

	Raft *raft.Raft

	state *state.State

	stable *boltdb.BoltStore

	grpcServer *grpc.Server

	sock net.Listener

	transport *transport.Manager

	config *config.Config
}

func NewServer(ctx context.Context, config *config.Config, grpcOpts []grpc.ServerOption) (*Server, error) {
	s := &Server{
		config:     config,
		grpcServer: grpc.NewServer(grpcOpts...),
	}

	baseDir := filepath.Join(config.RaftDir, config.NodeID)
	err := os.MkdirAll(baseDir, os.ModePerm)
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
		if err := f.Error(); err != nil {
			return nil, fmt.Errorf("Raft.Raft.BootstrapCluster: %v", err)
		}
	}
	return s, nil
}

type RegisterServiceFunc func(raft *raft.Raft, fsm *fsm.FSM) (desc *grpc.ServiceDesc, impl interface{})

func (s *Server) RegisterServices(services ...RegisterServiceFunc) error {
	for _, service := range services {
		desc, impl := service(s.Raft, s.fsm)
		s.grpcServer.RegisterService(desc, impl)
	}
	//把grpc server绑定到transport实现端口复用
	//s.transport.Register(s.grpcServer)
	return nil
}

func (s *Server) Start() error {
	//todo
	raftadmin.Register(s.grpcServer, s.Raft) //raft 管理 grpc
	reflection.Register(s.grpcServer)
	s.transport.Register(s.grpcServer)

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
