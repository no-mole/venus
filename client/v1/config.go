package clientv1

import (
	"context"
	"crypto/tls"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type Config struct {
	// Endpoints is a list of URLs.
	Endpoints []string `json:"endpoints"`

	// DialTimeout is the timeout for failing to establish a connection.default is 200*Millisecond.
	DialTimeout time.Duration `json:"dial-timeout"`

	// DefaultCallTimeout is the timeout for client call when ctx not set deadline.default is 2*Second.
	DefaultCallTimeout time.Duration `json:"default-call-timeout"`

	// DialKeepAliveTime is the time after which client pings the server to see if
	// transport is alive.
	DialKeepAliveTime time.Duration `json:"dial-keep-alive-time"`

	// DialKeepAliveTimeout is the time that the client waits for a response for the
	// keep-alive probe. If the response is not received in this time, the connection is closed.
	DialKeepAliveTimeout time.Duration `json:"dial-keep-alive-timeout"`

	MaxRetries uint `json:"max-retries"`

	PerCallTimeout time.Duration `json:"per-call-timeout"`

	// MaxCallSendMsgSize is the client-side request send limit in bytes.
	// If 0, it defaults to 2.0 MiB (2 * 1024 * 1024).
	// Make sure that "MaxCallSendMsgSize" < server-side default send/recv limit.
	MaxCallSendMsgSize int

	// MaxCallRecvMsgSize is the client-side response receive limit.
	// If 0, it defaults to "math.MaxInt32", because range response can
	// easily exceed request send limits.
	// Make sure that "MaxCallRecvMsgSize" >= server-side default send/recv limit.
	MaxCallRecvMsgSize int

	// TLS holds the client secure credentials, if any.
	TLS *tls.Config

	//PeerToken run client with admin
	PeerToken string `json:"peer-token"`

	// Username is a username for authentication.
	Username string `json:"username"`

	// Password is a password for authentication.
	Password string `json:"password"`

	//AccessKey is a accessKey for authentication
	AccessKey string `json:"access-key"`

	//AccessKey is a secret for authentication.
	AccessKeySecret string `json:"access-key-secret"`

	// DialOptions is a list of dial options for the grpc client (e.g., for interceptors).
	// For example, pass "grpc.WithBlock()" to block until the underlying connection is up.
	// Without this, Dial returns immediately and connecting the server happens in background.
	DialOptions []grpc.DialOption

	// Context is the default client context; it can be used to cancel grpc dial out and
	// other operations that do not have an explicit context.
	Context context.Context

	// Logger sets client-side logger.
	// If nil, fallback to building LogConfig.
	Logger *zap.Logger

	AutoSyncInterval time.Duration `json:"auto-sync-interval"`
}
