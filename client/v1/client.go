package clientv1

import (
	"context"
	"fmt"
	"github.com/no-mole/venus/client/v1/internal/resolver"
	"google.golang.org/grpc"
	"os"
	"strings"
)

type Client struct {
	KV
	Lease
	MicroService
	Namespace
	Cluster

	ctx context.Context

	cfg *Config
	// Username is a username for authentication.
	Username string
	// Password is a password for authentication.
	Password string

	resolver *resolver.ManualResolver

	callOpts []grpc.CallOption
	conn     *grpc.ClientConn
}

func NewClient(cfg Config) (*Client, error) {
	var err error
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
		c.Username = cfg.Username
		c.Password = cfg.Password
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

	c.conn, err = c.dial(grpc.WithResolvers(c.resolver))
	if err != nil {
		c.resolver.Close()
		return nil, err
	}

	c.KV = NewKV(c)
	c.Lease = NewLease(c)
	c.MicroService = NewMicroService(c)
	c.Cluster = NewCluster(c)
	c.Namespace = NewNamespace(c)

	//todo member list auto sync
	//todo gentoken

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
