package venus

import (
	"context"
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"github.com/no-mole/venus/proto/pbclient"
	"google.golang.org/grpc/backoff"
	"google.golang.org/grpc/keepalive"
	"net"
	"net/http"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/no-mole/venus/proto/pbsysconfig"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/hashicorp/raft"
	raftBoltdbStore "github.com/hashicorp/raft-boltdb/v2"
	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/logger"
	"github.com/no-mole/venus/agent/structs"
	"github.com/no-mole/venus/agent/transport"
	"github.com/no-mole/venus/agent/venus/api"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/config"
	"github.com/no-mole/venus/agent/venus/fsm"
	"github.com/no-mole/venus/agent/venus/lessor"
	"github.com/no-mole/venus/agent/venus/metrics"
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
	"github.com/no-mole/venus/proto/pbtransport"
	"github.com/no-mole/venus/proto/pbuser"
	bolt "go.etcd.io/bbolt"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
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
	pbaccesskey.UnimplementedAccessKeyServiceServer
	pbcluster.UnimplementedClusterServiceServer
	pbtransport.UnimplementedRaftTransportServer
	pbsysconfig.UnimplementedSysConfigServiceServer

	ctx context.Context

	config *config.Config

	//r a Raft node.
	r *raft.Raft
	//fsm is the client state machine to apply commands to
	fsm *fsm.FSM
	// state machine
	state *state.State
	//stable store for server conf
	stable raft.StableStore
	//logStore store for raft log
	logStore raft.LogStore

	grpcServer   *grpc.Server
	grpcListener net.Listener
	//router http api router
	router       *gin.Engine
	httpListener net.Listener

	//transport used for intra-cluster communication
	transport *transport.Manager

	//authTokenBundle credentials manager
	authTokenBundle credentials.Bundle

	//server (local server[Leader]) or (proxy server[Follower])
	server      server.Server
	localServer server.Server
	proxyServer server.Server

	//peerToken cluster peers certification token
	peerToken string
	//baseToken server admin token for long time,for transport
	baseToken *jwt.Token
	//authenticator is an authenticator for namespace write/read
	authenticator    auth.Authenticator
	metricsCollector *metrics.PrometheusCollector

	logger *zap.Logger

	rwLock sync.RWMutex

	//client is a client only connect to raft leader
	client *clientv1.Client

	errCh             chan error
	stopLeasesWatcher chan struct{}

	lessor              *lessor.Lessor
	lessorStatusStarted bool
	leasesExpiredNotify chan int64
	sysConfig           *pbsysconfig.SysConfig

	kvWatchers    map[string]map[string]map[int64]*kvWatcherInfo
	kvWatcherLock sync.RWMutex

	peerNodeClients     map[raft.ServerAddress]*clientv1.Client
	peerNodeClientsLock sync.RWMutex
}

type kvWatcherInfo struct {
	id         int64
	ch         chan *pbkv.KVItem
	clientInfo *pbclient.ClientInfo
}

func NewServer(ctx context.Context, conf *config.Config) (_ *Server, err error) {
	s := &Server{
		ctx:                 ctx,
		config:              conf,
		errCh:               make(chan error, 1),
		leasesExpiredNotify: make(chan int64, 16),
		stopLeasesWatcher:   make(chan struct{}, 1),
		kvWatchers:          make(map[string]map[string]map[int64]*kvWatcherInfo, 128), //todo config
		peerNodeClients:     make(map[raft.ServerAddress]*clientv1.Client),
	}
	s.lessor = lessor.NewLessor(ctx, s.leasesExpiredNotify)

	//init logger
	zapConf := logger.NewZapConfig(conf.ZapLoggerLevel())
	zapLogger, err := zapConf.Build(zap.AddCaller())
	if err != nil {
		panic(err)
	}
	s.logger = zapLogger.Named("venus").Named("server")

	//create data dir
	baseDir := filepath.Join(conf.DaftDir, conf.NodeID)
	err = os.MkdirAll(baseDir, os.ModePerm)
	if err != nil {
		s.logger.Error("make data dir", zap.Error(err), zap.String("baseDir", baseDir))
		return nil, err
	}
	//init db
	dbPath := fmt.Sprintf("%s/data.db", baseDir)
	db, err := bolt.Open(dbPath, 0666, &bolt.Options{
		Timeout:      time.Second,
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

	//init stable store
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

	// wrap the store in a LogCache to improve performance.
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
	//todo write peer token to file

	tokenProvider := auth.NewTokenProvider([]byte(s.peerToken))
	s.authenticator = auth.NewAuthenticator(tokenProvider)
	//gen long time expired token
	s.baseToken = auth.NewJwtTokenWithClaim(time.Now().Add(24*10000*time.Hour), "venus", "venus", auth.TokenTypeAdministrator, nil)
	s.baseToken.Raw, err = s.authenticator.Sign(s.ctx, s.baseToken)
	s.ctx = auth.WithContext(s.ctx, s.baseToken)
	if err != nil {
		s.logger.Error("gen base token failed", zap.Error(err))
		return nil, err
	}
	s.authTokenBundle = credentials.NewBundle()
	s.authTokenBundle.UpdateAuthToken("bearer", s.baseToken.Raw, 0)

	go s.startGrpcServer()
	go s.startHttpServer()

	//using grpc transport
	s.transport = transport.New(
		s.ctx,
		raft.ServerAddress(conf.LocalAddr),
		[]grpc.DialOption{
			grpc.WithPerRPCCredentials(s.authTokenBundle.PerRPCCredentials()),
			grpc.WithTransportCredentials(insecure.NewCredentials()),
			//todo 参数优化
			grpc.WithKeepaliveParams(keepalive.ClientParameters{
				Time:                10 * time.Second,
				Timeout:             1 * time.Second,
				PermitWithoutStream: true,
			}),
			grpc.WithConnectParams(grpc.ConnectParams{
				Backoff: backoff.Config{
					BaseDelay:  1.0 * time.Second,
					Multiplier: 1.6,
					Jitter:     0.2,
					MaxDelay:   5 * time.Second,
				},
				MinConnectTimeout: time.Second,
			}),
		},
	)

	collector := metrics.NewMetricsCollector("venus", 1*time.Second)
	s.metricsCollector = collector

	c := raft.DefaultConfig()
	c.LogLevel = conf.HcLoggerLevel().String()
	c.LocalID = raft.ServerID(conf.NodeID)
	c.LeaderLeaseTimeout = 1 * time.Second
	c.SnapshotInterval = 60 * time.Second
	c.SnapshotThreshold = 8192
	c.NoSnapshotRestoreOnStart = true //不需要从快照恢复，因为fsm/state数据是持久化的

	s.r, err = raft.NewRaft(c, s.fsm, logStore, s.stable, snap, s.transport.Transport())
	if err != nil {
		s.logger.Error("new raft failed", zap.Error(err))
		return nil, err
	}
	s.logger.Info("raft info", zap.Uint64("LastIndex", s.r.LastIndex()), zap.Uint64("AppliedIndex", s.r.AppliedIndex()))

	s.changeServeLoop()

	s.sysConfig, err = s.loadSysConf(s.ctx)
	if err != nil {
		s.logger.Error("load sys config err", zap.Error(err))
		return nil, err
	}
	s.watchSysConfig()
	s.kvWatcherDispatcher()

	//join or boot
	if s.config.JoinAddr != "" {
		s.logger.Info("join exist cluster", zap.String("joinAddr", s.config.JoinAddr))
		tokenCtx := auth.WithContext(s.ctx, s.baseToken)
		_, err = s.AddVoter(tokenCtx, &pbcluster.AddVoterRequest{
			Id:            s.config.NodeID,
			Address:       s.config.LocalAddr,
			PreviousIndex: s.r.LastIndex(),
		})
		if err != nil {
			return nil, err
		}
	} else if s.config.BootstrapCluster {
		err = s.BootstrapCluster()
		if err != nil {
			if err != raft.ErrCantBootstrap {
				s.logger.Error("bootstrap cluster failed", zap.Error(err))
				return nil, err
			}
			s.logger.Warn(raft.ErrCantBootstrap.Error())
		}
	}

	return s, nil
}

func (s *Server) listen(ep string) (net.Listener, error) {
	var lc net.ListenConfig
	listener, err := lc.Listen(s.ctx, "tcp", ep)
	if err != nil {
		s.logger.Error("start listener failed", zap.Error(err), zap.String("endpoint", ep))
		return nil, err
	}
	return listener, nil
}

func (s *Server) startGrpcServer() {
	var err error
	s.initGrpcServer()
	s.grpcListener, err = s.listen(s.config.GrpcEndpoint)
	if err != nil {
		s.logger.Error("start grpc listener failed", zap.Error(err), zap.String("endpoint", s.config.GrpcEndpoint))
		s.ReportError(err)
		return
	}
	s.logger.Info("grpc server will start!")
	s.ReportError(s.grpcServer.Serve(s.grpcListener))
}

func (s *Server) startHttpServer() {
	var err error
	if s.config.ZapLoggerLevel().Level() >= zap.InfoLevel {
		gin.SetMode(gin.ReleaseMode)
	}
	s.router = api.Router(s.config.HttpEndpoint, s, s.authenticator)
	if s.config.ZapLoggerLevel().Level() < zap.InfoLevel {
		pprof.Register(s.router)
	}
	s.httpListener, err = s.listen(s.config.HttpEndpoint)
	if err != nil {
		s.logger.Error("start http listener failed", zap.Error(err), zap.String("endpoint", s.config.HttpEndpoint))
		s.ReportError(err)
		return
	}
	s.logger.Info("http server will start!")

	if s.config.CertFile != "" && s.config.KeyFile != "" {
		err = http.ServeTLS(s.httpListener, s.router, s.config.CertFile, s.config.KeyFile)
	} else {
		err = http.Serve(s.httpListener, s.router)
	}
	s.ReportError(err)
}

func (s *Server) initGrpcServer() {
	serverOptions := []grpc.ServerOption{
		grpc.KeepaliveEnforcementPolicy(keepalive.EnforcementPolicy{
			MinTime:             200 * time.Millisecond, //最多200ms发送一次ping
			PermitWithoutStream: true,                   //没有活跃stream的情况下允许发送ping
		}),
		grpcMiddleware.WithUnaryServerChain(
			//recover server panic
			middlewares.UnaryServerRecover(middlewares.ZapLoggerRecoverHandle(s.logger)),
			//接受client传递的hostname等信息
			middlewares.UnaryServerWithCallerDetail(),
			//server access log
			middlewares.UnaryServerAccessLog(s.logger),
			//parse token from metadata
			middlewares.MustLoginUnaryServerInterceptor(s.authenticator),
			//使用读写锁保护ready状态，避免remote切换时候的并发读写
			middlewares.ReadLock(&s.rwLock),
		),
		grpcMiddleware.WithStreamServerChain(
			//recover server panic
			middlewares.StreamServerRecover(middlewares.ZapLoggerRecoverHandle(s.logger)),
			//接受client传递的hostname等信息
			middlewares.StreamServerWithCallerDetail(),
			//server access log
			middlewares.StreamServerAccessLog(s.logger),
			//parse token from metadata
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
		&pbsysconfig.SysConfigService_ServiceDesc,
	} {
		s.grpcServer.RegisterService(desc, s)
	}
	s.transport.Register(s.grpcServer)
}

func (s *Server) changeServeLoop() {
	notify := make(chan raft.Observation, 2)
	s.r.RegisterObserver(raft.NewObserver(notify, true, func(o *raft.Observation) bool {
		_, ok := o.Data.(raft.LeaderObservation)
		return ok
	}))
	var err error
	//default is local server
	s.localServer = local.NewLocalServer(s.r, s.fsm, s.config.ApplyTimeout)
	s.server = s.localServer

	cfg := clientv1.Config{
		Endpoints: []string{s.config.GrpcEndpoint}, //does not enter into force
		Logger:    s.logger,
	}
	if s.config.JoinAddr != "" {
		cfg.Endpoints = []string{s.config.JoinAddr}
	}
	s.client, err = clientv1.NewClient(cfg) //todo
	if err != nil {
		s.ReportError(err)
		return
	}
	s.proxyServer = proxy.NewRemoteServer(s.client)
	if s.config.JoinAddr != "" {
		s.server = s.proxyServer
	}
	go func() {
		for {
			select {
			case <-s.ctx.Done():
				return
			case <-notify:
				leaderAddr, leaderID := s.r.LeaderWithID()

				metrics.Collector.HasLeader(leaderAddr == "", string(leaderAddr), string(leaderID))

				s.logger.Info("raft leader changed", zap.String("leaderAddr", string(leaderAddr)), zap.String("leaderID", string(leaderID)))
				if leaderAddr == "" {
					continue
				}
				s.rwLock.Lock()
				if s.r.State() == raft.Leader {
					s.server = s.localServer
					if !s.lessorStatusStarted {
						s.lessor.StartChecker()
						s.lessorStatusStarted = true
					}
					err = s.watcherForLeases()
					if err != nil {
						s.ReportError(err)
						return
					}
				} else {
					if s.lessorStatusStarted {
						s.lessor.StopChecker()
						s.lessorStatusStarted = false
					}
					s.client.SetEndpoints([]string{string(leaderAddr)})
					s.server = s.proxyServer
				}
				s.logger.Info("set current node state", zap.String("state", s.r.State().String()))
				s.rwLock.Unlock()
			}
		}
	}()
}

func (s *Server) BootstrapCluster() error {
	cfg := raft.Configuration{
		Servers: []raft.Server{
			{
				Suffrage: raft.Voter,
				ID:       raft.ServerID(s.config.NodeID),
				Address:  raft.ServerAddress(s.config.LocalAddr),
			},
		},
	}
	err := s.r.BootstrapCluster(cfg).Error()
	if err != nil {
		return err
	}
	s.waitingForRaftCampaignLeader(500*time.Millisecond, 30*time.Second)
	//server register init default user // not login
	_, err = s.server.UserRegister(s.ctx, defaultUser)
	if err != nil {
		panic(err)
	}
	s.logger.Info("register default user",
		zap.String("defaultUserUid", defaultUser.Uid),
		zap.String("defaultUserPassword", structs.DefaultPassword),
	)
	return err
}

var (
	defaultUserUid  = "venus"
	defaultUserName = "VENUS"
	defaultUserRole = pbuser.UserRole_UserRoleAdministrator.String()
)

var defaultUser = &pbuser.UserInfo{
	Uid:      defaultUserUid,
	Name:     defaultUserName,
	Password: structs.DefaultPassword,
	Status:   pbuser.UserStatus_UserStatusEnable,
	Role:     defaultUserRole,
}

func (s *Server) ReportError(err error) {
	if err != nil {
		s.errCh <- err
	}
}

func (s *Server) Wait() error {
	s.logger.Info("server started,waiting report err")
	return <-s.errCh
}

func (s *Server) waitingForRaftCampaignLeader(backOffTime time.Duration, timeout time.Duration) {
	after := time.After(timeout)
	for {
		if s.r.State() == raft.Leader {
			return
		}
		s.logger.Info("waiting for raft campaign leader")
		select {
		case <-after:
		default:
			<-time.After(backOffTime)
		}
	}
}

func (s *Server) watcherForLeases() error {
	if s.r.State() != raft.Leader {
		return nil
	}
	chGrantId, chGrant := s.fsm.RegisterWatcher(structs.LeaseGrantRequestType)
	chRevokeId, chRevoke := s.fsm.RegisterWatcher(structs.LeaseRevokeRequestType)
	leases, err := s.Leases(s.ctx, nil)
	if err != nil {
		return err
	}
	s.logger.Info("load all leases")
	err = s.lessor.Reload(leases.Leases)
	if err != nil {
		return err
	}
	go func() {
		defer func() {
			s.fsm.UnregisterWatcher(structs.LeaseGrantRequestType, chGrantId)
			s.fsm.UnregisterWatcher(structs.LeaseRevokeRequestType, chRevokeId)
		}()
		for {
			select {
			case <-s.ctx.Done():
				s.logger.Info("stop leases watcher")
				return
			case <-s.stopLeasesWatcher:
				s.logger.Info("stop leases watcher")
				return
			case cmd, ok := <-chGrant:
				if !ok {
					s.logger.Info("stop leases watcher")
					return
				}
				_, data, _ := cmd()
				lease := &pblease.Lease{}
				err = codec.Decode(data, lease)
				if err != nil {
					s.logger.Error("decode lease grant msg", zap.Error(err))
					continue
				}
				_, err = s.lessor.Get(lease.LeaseId)
				if err != nil {
					err = s.lessor.Grant(lease)
				} else {
					err = s.lessor.Keepalive(lease.LeaseId, lease.Ddl)
				}
				if err != nil {
					s.logger.Error("lessor grant lease", zap.Error(err))
				}
			case cmd, ok := <-chRevoke:
				if !ok {
					return
				}
				_, data, _ := cmd()
				req := &pblease.RevokeRequest{}
				err = codec.Decode(data, req)
				if err != nil {
					s.logger.Error("decode revoke lease msg", zap.Error(err))
					continue
				}
				s.lessor.Revoke(req.LeaseId)
			case id, ok := <-s.leasesExpiredNotify:
				if !ok {
					s.logger.Info("stop leases watcher")
					return
				}
				if s.r.State() != raft.Leader {
					continue
				}
				_, err = s.Revoke(s.ctx, &pblease.RevokeRequest{LeaseId: id})
				if err != nil {
					s.logger.Error("leasesExpiredNotify revoke lease", zap.Error(err))
				}
			}
		}
	}()
	return nil
}

func (s *Server) watchSysConfig() {
	go func() {
		id, ch := s.fsm.RegisterWatcher(structs.SysConfigAddRequestType)
		defer s.fsm.UnregisterWatcher(structs.SysConfigAddRequestType, id)
		for {
			select {
			case <-s.ctx.Done():
				return
			case fn, ok := <-ch:
				if !ok {
					return
				}
				_, data, _ := fn()
				item := &pbsysconfig.SysConfig{}
				err := codec.Decode(data, item)
				if err != nil {
					s.logger.Error("decode sys config err", zap.Error(err))
					return
				}
				s.rwLock.Lock()
				s.sysConfig = item
				s.rwLock.Unlock()
			}
		}
	}()
}

func (s *Server) kvWatcherDispatcher() {
	go func() {
		id, ch := s.fsm.RegisterWatcher(structs.KVAddRequestType)
		defer s.fsm.UnregisterWatcher(structs.KVAddRequestType, id)
		for {
			select {
			case <-s.ctx.Done():
				return
			case cmd := <-ch:
				go func() {
					_, data, _ := cmd()
					item := &pbkv.KVItem{}
					err := codec.Decode(data, item)
					if err != nil {
						s.logger.Error("", zap.Error(err))
						return
					}
					s.kvWatcherLock.RLock()
					defer s.kvWatcherLock.RUnlock()
					if ns, ok := s.kvWatchers[item.Namespace]; ok {
						if infos, ok := ns[item.Key]; ok {
							now := time.Now().Format(time.RFC3339)
							for _, info := range infos {
								info.clientInfo.LastInteractionTime = now
								go func(cur *kvWatcherInfo) {
									cur.ch <- item
								}(info)
							}
						}
					}
				}()
			}
		}
	}()
}

func (s *Server) kvWatcherRegister(namespace, key string, clientInfo *pbclient.ClientInfo) (int64, chan *pbkv.KVItem) {
	info := &kvWatcherInfo{
		id:         time.Now().UnixNano(),
		ch:         make(chan *pbkv.KVItem, 4),
		clientInfo: clientInfo,
	}
	s.kvWatcherLock.Lock()
	defer s.kvWatcherLock.Unlock()
	if ns, ok := s.kvWatchers[namespace]; ok {
		if keys, ok := ns[key]; ok {
			keys[info.id] = info
		} else {
			ns[key] = map[int64]*kvWatcherInfo{info.id: info}
		}
	} else {
		s.kvWatchers[namespace] = map[string]map[int64]*kvWatcherInfo{key: {info.id: info}}
	}
	return info.id, info.ch
}
func (s *Server) kvWatcherUnregister(namespace, key string, id int64) {
	go func() {
		s.kvWatcherLock.Lock()
		defer s.kvWatcherLock.Unlock()
		if ns, ok := s.kvWatchers[namespace]; ok {
			if keys, ok := ns[key]; ok {
				if watcher, ok := keys[id]; ok {
					close(watcher.ch)
					delete(s.kvWatchers[namespace][key], id)
				}
			}
		}
	}()
}

func (s *Server) peerNodeClient(nodeAddr raft.ServerAddress) (*clientv1.Client, error) {
	s.peerNodeClientsLock.Lock()
	defer s.peerNodeClientsLock.Unlock()
	if client, ok := s.peerNodeClients[nodeAddr]; ok {
		return client, nil
	}
	cli, err := clientv1.NewClient(clientv1.Config{
		Endpoints:   []string{string(nodeAddr)},
		DialTimeout: time.Second,
		PeerToken:   s.peerToken,
		Context:     s.ctx,
		Logger:      s.logger.Named("cluster-nodes"),
	})
	if err != nil {
		return nil, err
	}
	s.peerNodeClients[nodeAddr] = cli
	return cli, nil
}
