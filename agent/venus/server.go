package venus

import (
	"context"
	"fmt"
	transport "github.com/Jille/raft-grpc-transport"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/hashicorp/raft"
	raftBoltdbStore "github.com/hashicorp/raft-boltdb/v2"
	"github.com/no-mole/venus/agent/venus/api"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/agent/venus/server/local"
	"github.com/no-mole/venus/agent/venus/server/proxy"
	"github.com/no-mole/venus/agent/venus/state"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbservice"
	"github.com/no-mole/venus/proto/pbuser"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
	"google.golang.org/grpc/status"
	"net"
	"net/http"
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
	pbcluster.UnimplementedClusterServer

	ctx context.Context

	r        *raft.Raft
	fsm      *fsm.FSM
	state    *state.State
	stable   raft.StableStore
	logStore raft.LogStore

	grpcServer   *grpc.Server
	grpcListener net.Listener

	router       *gin.Engine
	httpListener net.Listener

	transport *transport.Manager
	config    *config.Config
	remote    server.Server

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

	baseDir := filepath.Join(conf.DaftDir, conf.NodeID)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		s.logger.Error("make data dir", zap.Error(err), zap.String("baseDir", baseDir))
		return nil, err
	}
	dbPath := fmt.Sprintf("%s/data.db", baseDir)
	db, err := bolt.Open(dbPath, 0666, &bolt.Options{
		Timeout:      10 * time.Millisecond,
		FreelistType: bolt.FreelistMapType,
		NoSync:       true,
	})
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

	stableStoreFilePath := filepath.Join(baseDir, "stable.db")
	stable, err := raftBoltdbStore.New(raftBoltdbStore.Options{
		Path:   stableStoreFilePath,
		NoSync: true,
	})
	if err != nil {
		s.logger.Error("new stable store failed", zap.Error(err), zap.String("stableStoreFilePath", stableStoreFilePath))
		return nil, err
	}
	s.stable = stable
	// Wrap the store in a LogCache to improve performance.
	logStore, err := raft.NewLogCache(raftLogCacheSize, stable)
	if err != nil {
		s.logger.Error("wrap log cache failed", zap.Error(err))
		return nil, err
	}
	s.logStore = logStore

	snap, err := raft.NewFileSnapshotStore(baseDir, 3, os.Stderr)
	if err != nil {
		s.logger.Error("raft new file snapshot store failed", zap.Error(err))
		return nil, err
	}

	//using grpc transport
	s.transport = transport.New(raft.ServerAddress(conf.GrpcEndpoint), []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())})

	c := raft.DefaultConfig()
	c.LogLevel = conf.HcLoggerLevel().String()
	c.LocalID = raft.ServerID(conf.NodeID)
	c.SnapshotInterval = 60 * time.Second
	c.SnapshotThreshold = 8192

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

	if s.config.ZapLoggerLevel().Level() >= zap.InfoLevel {
		gin.SetMode(gin.ReleaseMode)
	}

	s.router = api.Router(s)
	if s.config.ZapLoggerLevel().Level() < zap.InfoLevel {
		pprof.Register(s.router)
	}

	return s, nil
}

func (s *Server) Start() error {
	if s.config.BootstrapCluster {
		err := s.BootstrapCluster()
		if err != nil {
			if err != raft.ErrCantBootstrap {
				s.logger.Error("bootstrap pbcluster failed", zap.Error(err))
				return err
			}
			s.logger.Warn("bootstrap pbcluster failed", zap.Error(err))
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
		client := pbcluster.NewClusterClient(conn)
		_, err = client.AddVoter(s.ctx, &pbcluster.AddVoterRequest{
			Id:            s.config.NodeID,
			Address:       s.config.GrpcEndpoint,
			PreviousIndex: s.r.LastIndex(),
		})
		if err != nil {
			s.logger.Error("add voter failed", zap.Error(err), zap.String("endpoint", s.config.JoinAddr))
			return err
		}
	}

	eg := errgroup.Group{}
	eg.Go(s.startGrpcServer)
	eg.Go(s.startHttpServer)
	return eg.Wait()
}

func (s *Server) startGrpcServer() (err error) {
	s.grpcListener, err = s.listen(s.config.GrpcEndpoint)
	if err != nil {
		s.logger.Error("start grpc listener failed", zap.Error(err), zap.String("endpoint", s.config.GrpcEndpoint))
		return err
	}
	s.logger.Info("grpc server started!")
	return s.grpcServer.Serve(s.grpcListener)
}

func (s *Server) startHttpServer() (err error) {
	s.httpListener, err = s.listen(s.config.HttpEndpoint)
	if err != nil {
		s.logger.Error("start http listener failed", zap.Error(err), zap.String("endpoint", s.config.HttpEndpoint))
		return err
	}
	s.logger.Info("http server started!")
	return http.Serve(s.httpListener, s.router)
}

func (s *Server) listen(ep string) (net.Listener, error) {
	listener, err := net.Listen("tcp", ep)
	if err != nil {
		s.logger.Error("start listener failed", zap.Error(err), zap.String("endpoint", ep))
		return nil, err
	}
	return listener, nil
}

func (s *Server) initGrpcServer() {
	grpc.WithDefaultCallOptions(
		grpc.WaitForReady(true),
	)
	// Give up after 5 retries.
	retryOpts := []grpcRetry.CallOption{
		grpcRetry.WithBackoff(grpcRetry.BackoffExponential(100 * time.Millisecond)),
		grpcRetry.WithMax(5),
	}
	grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor(retryOpts...))

	recoverHandle := func(ctx context.Context, fullMethodName string, p interface{}) (err error) {
		s.logger.Error("grpc server panic")
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}

	opts := []grpc.ServerOption{
		grpcMiddleware.WithUnaryServerChain(
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
				defer s.logger.Debug("grpc service caller", zap.String("serviceName", info.FullMethod), zap.String("duration", time.Now().Sub(start).String()))
				return handler(ctx, req)
			},
		),
		grpcMiddleware.WithStreamServerChain(
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
				defer s.logger.Info("grpc stream service caller", zap.String("serviceName", info.FullMethod), zap.String("duration", time.Now().Sub(start).String()))
				return handler(srv, ss)
			},
		),
	}
	s.grpcServer = grpc.NewServer(opts...)

	for _, desc := range []*grpc.ServiceDesc{&pbnamespace.NamespaceService_ServiceDesc, &pbkv.KV_ServiceDesc, &pblease.LeaseService_ServiceDesc, &pbservice.Service_ServiceDesc, &pbuser.UserService_ServiceDesc, &pbcluster.Cluster_ServiceDesc} {
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
