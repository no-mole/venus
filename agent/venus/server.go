package venus

import (
	"context"
	"fmt"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hashicorp/raft"
	boltdb "github.com/hashicorp/raft-boltdb"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/logger"
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
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
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

	logger *zap.Logger
}

func NewServer(ctx context.Context, conf *config.Config) (_ *Server, err error) {
	s := &Server{
		ctx:    ctx,
		config: conf,
	}
	zapConf := newZapConfig(conf)
	zapLogger, err := zapConf.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	s.logger = zapLogger.Named("venus").Named("server")

	baseDir := filepath.Join(conf.RaftDir, conf.NodeID)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		s.logger.Error("make data dir", zap.Error(err), zap.String("baseDir", baseDir))
		return nil, err
	}
	dbPath := fmt.Sprintf("%s/data.dat", baseDir)
	db, err := bolt.Open(dbPath, 0666, nil)
	if err != nil {
		s.logger.Error("bolt db open failed", zap.Error(err), zap.String("dbPath", dbPath))
		return nil, err
	}
	s.state = state.New(ctx, db, s.logger)

	s.fsm, err = fsm.NewBoltFSM(ctx, s.state, s.logger)
	if err != nil {
		s.logger.Error("new bolt fsm failed", zap.Error(err))
		return nil, err
	}

	stableStoreFilePath := filepath.Join(baseDir, "stable.dat")
	s.stable, err = boltdb.NewBoltStore(stableStoreFilePath)
	if err != nil {
		s.logger.Error("new stable store failed", zap.Error(err), zap.String("stableStoreFilePath", stableStoreFilePath))
		return nil, err
	}
	// Wrap the store in a LogCache to improve performance.
	logStore, err := raft.NewLogCache(raftLogCacheSize, s.stable)
	if err != nil {
		s.logger.Error("wrap log cache failed", zap.Error(err))
		return nil, err
	}

	snap, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		s.logger.Error("raft new file snapshot store failed", zap.Error(err))
		return nil, err
	}

	//using grpc transport
	s.transport = transport.New(raft.ServerAddress(conf.GrpcEndpoint), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	c := raft.DefaultConfig()
	c.HeartbeatTimeout = 3 * time.Second
	c.ElectionTimeout = 3 * time.Second
	c.SnapshotInterval = 10 * time.Second
	c.LogLevel = conf.HcLoggerLevel().String()
	c.Logger = logger.New("raft", conf.HcLoggerLevel(), s.logger)
	c.LocalID = raft.ServerID(conf.NodeID)

	s.r, err = raft.NewRaft(c, s.fsm, logStore, s.stable, snap, s.transport.Transport())
	if err != nil {
		s.logger.Error("new raft failed", zap.Error(err))
		return nil, err
	}
	s.logger.Info("raft info",
		zap.Uint64("LastIndex", s.r.LastIndex()),
		zap.Uint64("AppliedIndex", s.r.AppliedIndex()),
	)

	s.initGrpcServer()
	s.changeRemoteLoop()

	return s, nil
}

func (s *Server) Start() error {
	if s.config.BootstrapCluster {
		err := s.BootstrapCluster()
		if err != nil {
			s.logger.Warn("bootstrap cluster failed", zap.Error(err))
			return err
		}
	}
	if s.config.JoinAddr != "" {
		s.logger.Info("join node", zap.String("endpoint", s.config.JoinAddr))
		conn, err := grpc.Dial(s.config.JoinAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
		if err != nil {
			s.logger.Error("dial node failed", zap.Error(err), zap.String("endpoint", s.config.JoinAddr))
			return err
		}
		defer conn.Close()
		client := pbraftadmin.NewRaftAdminClient(conn)
		_, err = client.AddVoter(s.ctx, &pbraftadmin.AddVoterRequest{
			Id:            s.config.NodeID,
			Address:       s.config.GrpcEndpoint,
			PreviousIndex: s.r.LastIndex(),
		})
		if err != nil {
			s.logger.Error("add voter failed", zap.Error(err), zap.String("endpoint", s.config.JoinAddr))
			return err
		}
	}

	s.logger.Info("start listener", zap.String("endpoint", s.config.GrpcEndpoint))
	_, port, err := net.SplitHostPort(s.config.GrpcEndpoint)
	if err != nil {
		s.logger.Error("split host port failed", zap.Error(err), zap.String("endpoint", s.config.GrpcEndpoint))
		return err
	}
	s.sock, err = net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		s.logger.Error("start listener failed", zap.Error(err), zap.String("endpoint", s.config.GrpcEndpoint))
		return err
	}
	s.logger.Info("server started!")
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

	recoverHandle := func(ctx context.Context, fullMethodName string, p interface{}) (err error) {
		s.logger.Error("grpc server panic")
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
			//recover stream panicked
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				panicked := true
				defer func() {
					if r := recover(); r != nil || panicked {
						err = recoverHandle(ctx, info.FullMethod, r)
					}
				}()
				resp, err = handler(ctx, req)
				panicked = false
				return resp, err
			},
			func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
				start := time.Now()
				defer s.logger.Info("grpc service caller", zap.String("serviceName", info.FullMethod), zap.Int64("durationNano", time.Now().Sub(start).Nanoseconds()))
				return handler(ctx, req)
			},
		),
		grpc_middleware.WithStreamServerChain(
			//recover stream panicked
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
				panicked := true
				defer func() {
					if r := recover(); r != nil || panicked {
						err = recoverHandle(ss.Context(), info.FullMethod, r)
					}
				}()
				err = handler(srv, ss)
				panicked = false
				return err
			},
			func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
				start := time.Now()
				defer s.logger.Info("grpc stream service caller", zap.String("serviceName", info.FullMethod), zap.Int64("durationNano", time.Now().Sub(start).Nanoseconds()))
				return handler(srv, ss)
			},
		),
	}

	s.grpcServer = grpc.NewServer(opts...)
	for _, desc := range []*grpc.ServiceDesc{&pbnamespace.NamespaceService_ServiceDesc, &pbkv.KV_ServiceDesc, &pblease.LeaseService_ServiceDesc, &pbservice.Service_ServiceDesc, &pbuser.UserService_ServiceDesc, &pbraftadmin.RaftAdmin_ServiceDesc} {
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
	s.logger.Debug("register raft observer")
	endpoint := fmt.Sprintf("%s://venus-servers", scheme)
	cc, err := grpc.Dial(endpoint, grpc.WithResolvers(rs), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		s.logger.Error("create default grpc client conn failed", zap.Error(err), zap.String("endpoint", endpoint))
		panic(err)
	}
	go func() {
		for range ch {
			leaderAddr, leaderID := s.r.LeaderWithID()
			s.logger.Info("raft leader changed", zap.String("leaderAddr", string(leaderAddr)), zap.String("leaderID", string(leaderID)))
			s.readyLock.Lock()
			if s.r.State() == raft.Leader {
				s.remote = local.NewLocalServer(s.r, s.fsm, s.config)
			} else {
				rs.ResolveNow(resolver.ResolveNowOptions{})
				s.remote = proxy.NewRemoteServer(cc)
			}
			s.logger.Info("set current node state", zap.String("state", s.r.State().String()))
			s.readyLock.Unlock()
		}
	}()
}

func (s *Server) BootstrapCluster() error {
	cfg := raft.Configuration{
		Servers: []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       raft.ServerID(s.config.NodeID),
				Address:  raft.ServerAddress(s.config.GrpcEndpoint),
			},
		},
	}
	return s.r.BootstrapCluster(cfg).Error()
}
