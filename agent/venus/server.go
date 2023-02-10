package venus

import (
	"context"
	"fmt"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hashicorp/go-hclog"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/agent/venus/server/local"
	"github.com/no-mole/venus/agent/venus/server/proxy"
	"github.com/no-mole/venus/agent/venus/state"
	"github.com/no-mole/venus/internal/proto/pbraftadmin"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbservice"
	"github.com/no-mole/venus/proto/pbuser"
	bolt "go.etcd.io/bbolt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
	"log"
	"net"
	"os"
	"path/filepath"
	"sync"
	"time"
)

const (
	// raftLogCacheSize is the maximum number of logs to cache in-memory.
	// This is used to reduce disk I/O for the recently committed entries.
	raftLogCacheSize = 512
)

type Server struct {
	pbkv.UnimplementedKVServer
	pbnamespace.UnimplementedNamespaceServiceServer
	pblease.UnimplementedLeaseServiceServer
	pbservice.UnimplementedServiceServer
	pbuser.UnimplementedUserServiceServer
	pbraftadmin.UnimplementedRaftAdminServer

	ctx context.Context

	r      *raft.Raft
	fsm    *fsm.FSM
	state  *state.State
	stable *boltdb.BoltStore

	grpcServer *grpc.Server
	sock       net.Listener
	transport  *transport.Manager
	config     *config.Config
	remote     server.Server

	readyLock sync.RWMutex
}

func NewServer(ctx context.Context, config *config.Config) (_ *Server, err error) {
	if config.ApplyTimeout == 0 {
		config.ApplyTimeout = 5 * time.Second
	}

	s := &Server{
		ctx:    ctx,
		config: config,
	}

	baseDir := filepath.Join(config.RaftDir, config.NodeID)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		return nil, fmt.Errorf(`os.MkdirAll(%q): %v`, baseDir, err)
	}
	dbPath := fmt.Sprintf("%s/data.dat", baseDir)
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		return nil, fmt.Errorf(`r.bolt.Open(%q, ...): %v`, dbPath, err)
	}
	s.state = state.New(ctx, db)

	s.fsm, err = fsm.NewBoltFSM(ctx, s.state)
	if err != nil {
		return nil, fmt.Errorf(`fsm.NewBoltFSM(%q, ...): %v`, baseDir, err)
	}

	s.stable, err = boltdb.NewBoltStore(filepath.Join(baseDir, "stable.dat"))
	if err != nil {
		return nil, fmt.Errorf(`boltdb.NewBoltStore(%q): %v`, filepath.Join(baseDir, "stable.dat"), err)
	}
	// Wrap the store in a LogCache to improve performance.
	logStore, err := raft.NewLogCache(raftLogCacheSize, s.stable)
	if err != nil {
		return nil, fmt.Errorf("r.NewLogCache():%v", err)
	}

	snap, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		return nil, fmt.Errorf(`r.NewFileSnapshotStore(%q, ...): %v`, baseDir, err)
	}

	//using grpc transport
	s.transport = transport.New(raft.ServerAddress(config.GrpcEndpoint), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	c := raft.DefaultConfig()
	c.HeartbeatTimeout = 3 * time.Second
	c.ElectionTimeout = 3 * time.Second
	c.LogLevel = hclog.Warn.String()
	c.LocalID = raft.ServerID(config.NodeID)
	s.r, err = raft.NewRaft(c, s.fsm, logStore, s.stable, snap, s.transport.Transport())
	if err != nil {
		return nil, fmt.Errorf("r.NewRaft: %v", err)
	}

	if config.BootstrapCluster {
		cfg := raft.Configuration{
			Servers: []raft.Server{
				{
					Suffrage: raft.Voter,
					ID:       raft.ServerID(config.NodeID),
					Address:  raft.ServerAddress(config.GrpcEndpoint),
				},
			},
		}
		err = s.r.BootstrapCluster(cfg).Error()
		if err != nil {
			if err == raft.ErrCantBootstrap {
				log.Printf("r.BootstrapCluster: %s", err.Error())
			} else {
				return nil, err
			}
		}
	}

	s.initGrpcServer()
	s.changeRemoteLoop()

	return s, nil
}

func (s *Server) Start() error {
	if s.config.JoinAddr != "" {
		conn, err := grpc.Dial(s.config.JoinAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			log.Fatalf("failed grpc.Dial(%s): %s", s.config.JoinAddr, err)
		}
		defer conn.Close()
		println("join", s.config.JoinAddr)
		client := pbraftadmin.NewRaftAdminClient(conn)
		_, err = client.AddVoter(s.ctx, &pbraftadmin.AddVoterRequest{
			Id:            s.config.NodeID,
			Address:       s.config.GrpcEndpoint,
			PreviousIndex: s.r.LastIndex(),
		})
		if err != nil {
			log.Fatalf("failed server.AddVoter(%s): %s", s.config.JoinAddr, err)
		}
	}

	_, port, err := net.SplitHostPort(s.config.GrpcEndpoint)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s.sock, err = net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	println("start serve")
	go func() {
		ticker := time.NewTicker(time.Second * 2)
		for {
			<-ticker.C
			fmt.Printf(",cur:%s\n", s.r.String())
		}
	}()
	return s.grpcServer.Serve(s.sock)
}

func (s *Server) initGrpcServer() {
	grpc.WithDefaultCallOptions(
		grpc.WaitForReady(true),
	)
	// Give up after 5 retries.
	retryOpts := []grpc_retry.CallOption{
		grpc_retry.WithBackoff(grpc_retry.BackoffExponential(100 * time.Millisecond)),
		grpc_retry.WithMax(5),
	}
	grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(retryOpts...))

	recoverHandle := func(p interface{}) (err error) {
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc.ServerOption{
		grpc_middleware.WithUnaryServerChain(
			//使用读写锁保护ready状态，避免remote切换时候的并发读写
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				s.readyLock.RLock()
				defer s.readyLock.RUnlock()
				return handler(ctx, req)
			},
			grpc_recovery.UnaryServerInterceptor(grpc_recovery.WithRecoveryHandler(recoverHandle)),
		),
		grpc_middleware.WithStreamServerChain(
			grpc_recovery.StreamServerInterceptor(grpc_recovery.WithRecoveryHandler(recoverHandle)),
		),
	}

	s.grpcServer = grpc.NewServer(opts...)

	for _, desc := range []*grpc.ServiceDesc{
		&pbnamespace.NamespaceService_ServiceDesc,
		&pbkv.KV_ServiceDesc,
		&pblease.LeaseService_ServiceDesc,
		&pbservice.Service_ServiceDesc,
		&pbuser.UserService_ServiceDesc,
		&pbraftadmin.RaftAdmin_ServiceDesc,
	} {
		s.grpcServer.RegisterService(desc, s)
	}
	s.transport.Register(s.grpcServer)
}

func (s *Server) changeRemoteLoop() {
	s.remote = local.NewLocalServer(s.r, s.fsm, s.config)
	ch := make(chan raft.Observation, 1)
	s.r.RegisterObserver(raft.NewObserver(ch, true, func(o *raft.Observation) bool {
		_, ok := o.Data.(raft.LeaderObservation)
		return ok
	}))
	rs := &grpcResolver{
		r: s.r,
	}
	cc, err := grpc.Dial(fmt.Sprintf("%s://%s", scheme, s.config.GrpcEndpoint), grpc.WithResolvers(rs), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	go func() {
		for range ch {
			fmt.Printf("leader changed,cur:%s\n", s.r.String())
			s.readyLock.Lock()
			if s.r.State() == raft.Leader {
				s.remote = local.NewLocalServer(s.r, s.fsm, s.config)
			} else {
				rs.ResolveNow(resolver.ResolveNowOptions{})
				s.remote = proxy.NewRemoteServer(cc)
			}
			s.readyLock.Unlock()
		}
	}()
}
