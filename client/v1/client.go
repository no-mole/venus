package clientv1

import (
	"context"
	"fmt"
	"github.com/no-mole/venus/agent/logger"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/metrics"
	"go.uber.org/zap"
	"os"
	"strings"
	"time"

	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/no-mole/venus/agent/venus/middlewares"
	"github.com/no-mole/venus/client/v1/credentials"
	"github.com/no-mole/venus/client/v1/internal/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
)

type Client struct {
	KV
	Lease
	MicroService
	Namespace
	Cluster
	User
	AccessKey
	SysConfig

	ctx    context.Context
	cancel context.CancelFunc

	cfg *Config

	//peerToken admin token
	peerToken string

	// username is a username for authentication.
	username string
	// password is a password for authentication.
	password string

	//accessKey is a access key for authentication.
	accessKey string
	//accessKeySecret is a access key secret for authentication.
	accessKeySecret string

	resolver *resolver.ManualResolver

	authTokenBundle credentials.Bundle

	callOpts []grpc.CallOption
	conn     *grpc.ClientConn

	logger *zap.Logger
}

func NewClient(cfg Config) (_ *Client, err error) {
	if len(cfg.Endpoints) < 1 {
		return nil, fmt.Errorf("at least one Endpoint is required in client config")
	}
	if cfg.Logger == nil {
		cfg.Logger, err = logger.NewZapConfig(zap.NewAtomicLevelAt(zap.InfoLevel)).Build()
		if err != nil {
			return nil, err
		}
	}
	if cfg.Context == nil {
		cfg.Context = context.Background()
	}
	ctx, cancel := context.WithCancel(cfg.Context)
	c := &Client{
		ctx:             ctx,
		cancel:          cancel,
		cfg:             &cfg,
		callOpts:        defaultCallOpts,
		resolver:        resolver.New(cfg.Endpoints...),
		logger:          cfg.Logger.Named("client"),
		peerToken:       cfg.PeerToken,
		accessKey:       cfg.AccessKey,
		accessKeySecret: cfg.AccessKeySecret,
		username:        cfg.Username,
		password:        cfg.Password,
		authTokenBundle: credentials.NewBundle(),
	}

	if cfg.MaxCallSendMsgSize > 0 || cfg.MaxCallRecvMsgSize > 0 {
		if cfg.MaxCallRecvMsgSize > 0 && cfg.MaxCallSendMsgSize > cfg.MaxCallRecvMsgSize {
			return nil, fmt.Errorf("gRPC message recv limit (%d bytes) must be greater than send limit (%d bytes)", cfg.MaxCallRecvMsgSize, cfg.MaxCallSendMsgSize)
		}
		c.callOpts = defaultCallOpts
		if cfg.MaxCallSendMsgSize > 0 {
			c.callOpts[1] = grpc.MaxCallSendMsgSize(cfg.MaxCallSendMsgSize)
		}
		if cfg.MaxCallRecvMsgSize > 0 {
			c.callOpts[2] = grpc.MaxCallRecvMsgSize(cfg.MaxCallRecvMsgSize)
		}
	}

	c.conn, err = c.dial(c.ctx, grpc.WithResolvers(c.resolver))

	if err != nil {
		c.resolver.Close()
		return nil, err
	}

	c.KV = NewKV(c, c.logger)
	c.Lease = NewLease(c, c.logger)
	c.MicroService = NewMicroService(c, c.logger)
	c.Cluster = NewCluster(c, c.logger)
	c.Namespace = NewNamespace(c, c.logger)
	c.User = NewUser(c, c.logger)
	c.AccessKey = NewAccessKey(c, c.logger)
	c.SysConfig = NewSysConfig(c, c.logger)

	err = c.getToken()
	if err != nil {
		return nil, err
	}

	go c.checkTokenLoop()
	go c.autoSync()

	return c, nil
}

func (c *Client) buildGRPCTarget() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = fmt.Sprintf("%p", c)
	}
	return fmt.Sprintf("%s://%s/%s", resolver.Schema, hostname, strings.Join(c.cfg.Endpoints, ","))
}

func (c *Client) dial(ctx context.Context, dailOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	dailOpts = append(dailOpts,
		middlewares.ClientWaitForReady(),
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: c.cfg.DialKeepAliveTime, Timeout: c.cfg.DialKeepAliveTimeout}),
		grpc.WithResolvers(c.resolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(c.authTokenBundle.PerRPCCredentials()),
		grpc.WithChainUnaryInterceptor(
			middlewares.MustLoginUnaryClientInterceptor(),
			metrics.Collector.RpcRequestTotal(),
			metrics.Collector.RpcRequestDurationTime(),
			grpcRetry.UnaryClientInterceptor(
				grpcRetry.WithMax(c.cfg.MaxRetries),
				grpcRetry.WithPerRetryTimeout(c.cfg.PerCallTimeout),
			),
		),
		grpc.WithChainStreamInterceptor(
			middlewares.MustLoginStreamClientInterceptor(),
			metrics.Collector.RpcStreamRequestTotal(),
			metrics.Collector.RpcStreamRequestDurationTime(),
			grpcRetry.StreamClientInterceptor(
				grpcRetry.WithMax(c.cfg.MaxRetries),
				grpcRetry.WithPerRetryTimeout(c.cfg.PerCallTimeout),
			),
		),
	)

	target := c.buildGRPCTarget()
	if c.cfg.DialTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, c.cfg.DialTimeout)
		defer cancel()
	}
	conn, err := grpc.DialContext(ctx, target, append(dailOpts, c.cfg.DialOptions...)...)
	if err != nil {
		return nil, err
	}
	return conn, nil
}

// DialContext connects to a single endpoint using the client's config.
func (c *Client) DialContext(ctx context.Context, ep string) (*grpc.ClientConn, error) {
	// Using ad-hoc created resolver, to guarantee only explicitly given
	// endpoint is used.
	return c.dial(ctx, grpc.WithResolvers(resolver.New(ep)))
}

func (c *Client) getToken() error {
	if c.peerToken != "" {
		c.logger.Info("gen token with peer token")
		token := auth.NewJwtTokenWithClaim(time.Now().Add(24*10000*time.Hour), "venus", "venus", auth.TokenTypeAdministrator, nil)
		tokenProvider := auth.NewTokenProvider([]byte(c.peerToken))
		tokenString, err := tokenProvider.Sign(c.ctx, token)
		if err != nil {
			return err
		}
		c.authTokenBundle.UpdateAuthToken("bearer", tokenString, time.Duration(0))
	} else if c.accessKey != "" && c.accessKeySecret != "" {
		c.logger.Info("gen token with access key/secret")
		resp, err := c.AccessKey.AccessKeyLogin(c.ctx, c.accessKey, c.accessKeySecret)
		if err != nil {
			return err
		}
		c.authTokenBundle.UpdateAuthToken(resp.TokenType, resp.AccessToken, time.Duration(resp.ExpiredIn)*time.Second)
	} else if c.username != "" && c.password != "" {
		c.logger.Info("gen token with user/password")
		resp, err := c.User.UserLogin(c.ctx, c.username, c.password)
		if err != nil {
			return err
		}
		c.authTokenBundle.UpdateAuthToken(resp.TokenType, resp.AccessToken, time.Duration(resp.ExpiredIn))
	}
	return nil
}

func (c *Client) checkTokenLoop() {
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-ticker.C:
			if c.authTokenBundle.ShouldUpdateToken() {
				err := c.getToken()
				if err != nil {
					//todo logger err
				}
			}
		}
	}
}

func (c *Client) autoSync() {
	if c.cfg.AutoSyncInterval == 0 {
		return
	}
	for {
		select {
		case <-c.ctx.Done():
			return
		case <-time.After(c.cfg.AutoSyncInterval):
			ctx, cancel := context.WithTimeout(c.ctx, 5*time.Second)
			err := c.Sync(ctx)
			cancel()
			if err != nil && err != ctx.Err() {
				//todo logger err
			}
		}

	}
}

func (c *Client) Sync(ctx context.Context) error {
	resp, err := c.Nodes(ctx)
	if err != nil {
		return err
	}
	var eps []string
	for _, node := range resp.Nodes {
		if node.Id != "" && node.Address != "" {
			eps = append(eps, node.Address)
		}
	}
	c.resolver.SetEndpoints(eps)
	return nil
}

func (c *Client) Close() error {
	c.resolver.Close()
	err := c.conn.Close()
	if err != nil {
		return err
	}
	c.cancel()
	return nil
}

func (c *Client) SetEndpoints(eps []string) {
	c.resolver.SetEndpoints(eps)
}
