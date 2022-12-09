package venus

import (
	"context"
	"fmt"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/proto"
	"os"
	"path/filepath"
)

const (
	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	raftLogCacheSize = 512
)

type Server struct {
	Fsm *fsm.BoltFSM

	Raft *raft.Raft

	Transport *transport.Manager

	config *config.Config
}

func NewServer(ctx context.Context, config *config.Config) (*Server, error) {
	s := &Server{
		config: config,
	}
	c := raft.DefaultConfig()
	c.LocalID = raft.ServerID(config.NodeID)

	baseDir := filepath.Join(config.RaftDir, config.NodeID)
	err := os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf(`os.MkdirAll(%q): %v`, baseDir, err)
	}

	stable, err := boltdb.NewBoltStore(filepath.Join(baseDir, "stable.dat"))
	if err != nil {
		return nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "stable.dat"), err)
	}

	// Wrap the store in a LogCache to improve performance.
	cacheStore, err := raft.NewLogCache(raftLogCacheSize, stable)
	if err != nil {
		return nil, fmt.Errorf("Raft.NewLogCache():%v", err)
	}
	log := cacheStore

	snap, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf(`Raft.NewFileSnapshotStore(%q, ...): %v`, baseDir, err)
	}

	fsmInstance, err := fsm.New(ctx, proto.UnmarshalOptions{}, &fsm.BoltFSMConfig{
		DBPath:      fmt.Sprintf("%s/data.dat", baseDir),
		OpenMode:    0666,
		BoltOptions: nil,
	})
	if err != nil {
		return nil, fmt.Errorf(`Fsm.New(%q, ...): %v`, baseDir, err)
	}

	s.Fsm = fsmInstance

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
	s.Transport = tm

	r, err := raft.NewRaft(c, fsmInstance, log, stable, snap, tm.Transport())
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
