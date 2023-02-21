package clientv1

import (
	"context"
	"fmt"
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"github.com/no-mole/venus/client/v1/credentials"
	"github.com/no-mole/venus/client/v1/internal/resolver"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/keepalive"
	"os"
	"strings"
	"time"
)

type Client struct {
	KV
	Lease
	MicroService
	Namespace
	Cluster
	User
	AccessKey

	ctx context.Context

	cfg *Config
	// username is a username for authentication.
	username string
	// password is a password for authentication.
	password string

	accessKey       string
	accessKeySecret string

	resolver *resolver.ManualResolver

	authTokenBundle credentials.Bundle

	callOpts []grpc.CallOption
	conn     *grpc.ClientConn
}

func NewClient(cfg Config) (*Client, error) {
	var err error
	fmt.Printf("%+v\n", cfg)
	if len(cfg.Endpoints) < 1 {
		return nil, fmt.Errorf("at least one Endpoint is required in client config")
	}

	c := &Client{
		ctx:      context.Background(),
		cfg:      &cfg,
		callOpts: defaultCallOpts,
		resolver: resolver.New(cfg.Endpoints...),
	}

	if cfg.Context != nil {
		c.ctx = cfg.Context
	}

	if cfg.Username != "" && cfg.Password != "" {
		c.username = cfg.Username
		c.password = cfg.Password
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

	c.authTokenBundle = credentials.NewBundle()

	c.conn, err = c.dial()

	if err != nil {
		c.resolver.Close()
		return nil, err
	}

	c.KV = NewKV(c)
	c.Lease = NewLease(c)
	c.MicroService = NewMicroService(c)
	c.Cluster = NewCluster(c)
	c.Namespace = NewNamespace(c)
	//todo user access key

	err = c.getToken()
	if err != nil {
		return nil, err
	}

	go c.checkTokenLoop()

	//todo member list auto sync
	return c, nil
}

func (c *Client) buildGRPCTarget() string {
	hostname, err := os.Hostname()
	if err != nil {
		hostname = fmt.Sprintf("%p", c)
	}
	return fmt.Sprintf("%s://%s/%s", resolver.Schema, hostname, strings.Join(c.cfg.Endpoints, ","))
}

func (c *Client) dial(dailOpts ...grpc.DialOption) (*grpc.ClientConn, error) {
	dailOpts = append(dailOpts,
		grpc.WithKeepaliveParams(keepalive.ClientParameters{Time: c.cfg.DialKeepAliveTime, Timeout: c.cfg.DialKeepAliveTimeout}),
		grpc.WithResolvers(c.resolver),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
		grpc.WithPerRPCCredentials(c.authTokenBundle.PerRPCCredentials()),
		grpc.WithUnaryInterceptor(
			grpcRetry.UnaryClientInterceptor(
				grpcRetry.WithMax(c.cfg.MaxRetries),
				grpcRetry.WithPerRetryTimeout(c.cfg.PerCallTimeout),
			),
		),
		grpc.WithStreamInterceptor(
			grpcRetry.StreamClientInterceptor(
				grpcRetry.WithMax(c.cfg.MaxRetries),
				grpcRetry.WithPerRetryTimeout(c.cfg.PerCallTimeout),
			),
		),
	)

	target := c.buildGRPCTarget()
	dailCtx := c.ctx
	if c.cfg.DialTimeout > 0 {
		var cancel context.CancelFunc
		dailCtx, cancel = context.WithTimeout(dailCtx, c.cfg.DialTimeout)
		defer cancel()
	}

	conn, err := grpc.DialContext(dailCtx, target, append(dailOpts, c.cfg.DialOptions...)...)
	if err != nil {
		return nil, err
	}

	return conn, nil
}

// Dial connects to a single endpoint using the client's config.
func (c *Client) Dial(ep string) (*grpc.ClientConn, error) {
	// Using ad-hoc created resolver, to guarantee only explicitly given
	// endpoint is used.
	return c.dial(grpc.WithResolvers(resolver.New(ep)))
}

func (c *Client) getToken() error {
	if c.accessKey != "" && c.accessKeySecret != "" {
		resp, err := c.AccessKey.AccessKeyLogin(c.ctx, c.accessKey, c.accessKeySecret)
		if err != nil {
			return err
		}
		c.authTokenBundle.UpdateAuthToken(resp.TokenType, resp.AccessToken, time.Duration(resp.ExpiredIn)*time.Second)
	} else if c.username != "" && c.password != "" {
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
