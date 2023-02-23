package venus

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/hashicorp/raft"
	raftBoltdbStore "github.com/hashicorp/raft-boltdb/v2"
	"github.com/no-mole/venus/agent/logger"
	"github.com/no-mole/venus/agent/transport"
	"github.com/no-mole/venus/agent/venus/api"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/middlewares"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/agent/venus/server/local"
	"github.com/no-mole/venus/agent/venus/server/proxy"
	"github.com/no-mole/venus/agent/venus/state"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/no-mole/venus/client/v1/credentials"
	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/no-mole/venus/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"github.com/no-mole/venus/proto/pblease"
	"github.com/no-mole/venus/proto/pbmicroservice"
	"github.com/no-mole/venus/proto/pbnamespace"
	"github.com/no-mole/venus/proto/pbuser"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"golang.org/x/sync/errgroup"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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

var (
	stablePeerTokenKey = []byte("peer_token")
)

type Server struct {
	pbkv.UnimplementedKVServiceServer
	pbnamespace.UnimplementedNamespaceServiceServer
	pblease.UnimplementedLeaseServiceServer
	pbmicroservice.UnimplementedMicroServiceServer
	pbuser.UnimplementedUserServiceServer
	pbcluster.UnimplementedClusterServiceServer
	pbaccesskey.UnimplementedAccessKeyServiceServer

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

	authTokenBundle credentials.Bundle

	config *config.Config
	remote server.Server

	//cluster peers certification token
	peerToken string
	//server admin token
	baseToken     *jwt.Token
	authenticator auth.Authenticator

	logger *zap.Logger

	readyLock *sync.RWMutex
	once      sync.Once
}

func NewServer(ctx context.Context, conf *config.Config) (_ *Server, err error) {
	s := &Server{
		ctx:       ctx,
		config:    conf,
		readyLock: &sync.RWMutex{},
	}
	zapConf := logger.NewZapConfig(conf.ZapLoggerLevel())
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

	//fetch peer token from stable store or gen new one
	if s.config.PeerToken == "" {
		s.logger.Info("config peer token not found")
		value, err := s.stable.Get(stablePeerTokenKey)
		if err != nil && err != raftBoltdbStore.ErrKeyNotFound {
			return nil, err
		}
		if len(value) == 0 {
			s.logger.Warn("stable store peer token not found,gen new one")
			randToken := md5.Sum([]byte(time.Now().String()))
			s.peerToken = base64.RawURLEncoding.EncodeToString(randToken[:])
		} else {
			s.peerToken = string(value)
		}
	} else {
		s.peerToken = s.config.PeerToken
	}
	//save peer token stable
	err = s.stable.Set(stablePeerTokenKey, []byte(s.peerToken))
	if err != nil {
		s.logger.Error("save peer token to stable store failed", zap.Error(err))
		return nil, err
	}

	s.logger.Warn("cur peer token,must save it when you join cluster", zap.String("peer-token", s.peerToken))

	tokenProvider := auth.NewTokenProvider([]byte(s.peerToken))
	//gen long time expired token
	s.baseToken = auth.NewJwtTokenWithClaim(time.Now().Add(24*10000*time.Hour), auth.TokenTypeAdministrator, nil)
	s.authenticator = auth.NewAuthenticator(tokenProvider)
	tokenString, err := s.authenticator.Sign(s.ctx, s.baseToken)
	if err != nil {
		s.logger.Error("gen base token failed", zap.Error(err))
		return nil, err
	}
	s.authTokenBundle = credentials.NewBundle()
	s.authTokenBundle.UpdateAuthToken("bearer", tokenString, 0)
	//using grpc transport
	s.transport = transport.New(
		s.ctx,
		raft.ServerAddress(conf.GrpcEndpoint),
		[]grpc.DialOption{
			grpc.WithPerRPCCredentials(s.authTokenBundle.PerRPCCredentials()),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
		},
	)

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
	s.router = api.Router(s, s.authenticator)
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

	eg := errgroup.Group{}
	eg.Go(s.startGrpcServer)
	eg.Go(s.startHttpServer)
	go func() {
		if s.config.JoinAddr != "" {
			<-time.After(3 * time.Second)
			tokenCtx := auth.WithContext(s.ctx, s.baseToken)
			_, err := s.AddVoter(tokenCtx, &pbcluster.AddVoterRequest{
				Id:            s.config.NodeID,
				Address:       s.config.GrpcEndpoint,
				PreviousIndex: s.r.LastIndex(),
			})
			if err != nil {
				s.logger.Error("add voter failed", zap.Error(err), zap.String("endpoint", s.config.JoinAddr))
				return
			}
		}
	}()
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
	serverOptions := []grpc.ServerOption{
		grpcMiddleware.WithUnaryServerChain(
			middlewares.UnaryServerRecover(middlewares.ZapLoggerRecoverHandle(s.logger)),
			middlewares.UnaryServerAccessLog(s.logger),
			middlewares.MustLoginUnaryServerInterceptor(s.authenticator),
			//使用读写锁保护ready状态，避免remote切换时候的并发读写
			middlewares.ReadLock(s.readyLock),
		),
		grpcMiddleware.WithStreamServerChain(
			middlewares.StreamServerRecover(middlewares.ZapLoggerRecoverHandle(s.logger)),
			middlewares.StreamServerAccessLog(s.logger),
			middlewares.MustLoginStreamServerInterceptor(s.authenticator),
		),
	}
	s.grpcServer = grpc.NewServer(serverOptions...)

	for _, desc := range []*grpc.ServiceDesc{
		&pbnamespace.NamespaceService_ServiceDesc,
		&pbkv.KVService_ServiceDesc,
		&pblease.LeaseService_ServiceDesc,
		&pbmicroservice.MicroService_ServiceDesc,
		&pbuser.UserService_ServiceDesc,
		&pbaccesskey.AccessKeyService_ServiceDesc,
		&pbcluster.ClusterService_ServiceDesc,
	} {
		s.grpcServer.RegisterService(desc, s)
	}
	s.transport.Register(s.grpcServer)
}

func (s *Server) changeRemoteLoop() {
	var err error
	localServer := local.NewLocalServer(s.r, s.fsm, s.config)
	s.remote = localServer
	ch := make(chan raft.Observation, 1)
	s.r.RegisterObserver(raft.NewObserver(ch, true, func(o *raft.Observation) bool {
		_, ok := o.Data.(raft.LeaderObservation)
		return ok
	}))
	var client *clientv1.Client
	var proxyServer server.Server
	go func() {
		for range ch {
			leaderAddr, leaderID := s.r.LeaderWithID()
			s.logger.Info("raft leader changed", zap.String("leaderAddr", string(leaderAddr)), zap.String("leaderID", string(leaderID)))
			s.readyLock.Lock()
			if s.r.State() == raft.Leader {
				s.remote = localServer
			} else {
				s.once.Do(func() {
					endpoint, _ := s.r.LeaderWithID()
					cfg := clientv1.Config{
						Endpoints:   []string{string(endpoint)},
						DialOptions: []grpc.DialOption{grpc.WithChainUnaryInterceptor()},
						Logger:      s.logger,
					}
					client, err = clientv1.NewClient(cfg)
					if err != nil {
						s.logger.Error("create leader client failed", zap.Error(err), zap.String("endpoint", string(endpoint)))
						panic(err)
					}
					proxyServer = proxy.NewRemoteServer(client)
				})
				client.SetEndpoints([]string{string(leaderAddr)})
				s.remote = proxyServer
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
