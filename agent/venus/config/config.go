package config

import (
	"github.com/hashicorp/go-hclog"
	"go.uber.org/zap"
	"time"
)

type Config struct {
	NodeID           string        `json:"node_id" yaml:"node_id"`
	DaftDir          string        `json:"raft_dir"`
	GrpcEndpoint     string        `json:"grpc_endpoint"`
	HttpEndpoint     string        `json:"http_endpoint"`
	BootstrapCluster bool          `json:"bootstrap_cluster"`
	ApplyTimeout     time.Duration `json:"apply_timeout"`
	JoinAddr         string        `json:"join_addr"`
	LoggerLevel      LoggerLevel   `json:"logger_level"`
	PeerToken        string        `json:"peer_token"`
}

func GetDefaultConfig() *Config {
	return &Config{
		NodeID:           "",
		DaftDir:          "",
		GrpcEndpoint:     "127.0.0.1:3333",
		HttpEndpoint:     "127.0.0.1:2333",
		BootstrapCluster: false,
		ApplyTimeout:     1 * time.Second,
		JoinAddr:         "",
		LoggerLevel:      LoggerLevelInfo,
	}
}

type LoggerLevel string

const (
	LoggerLevelDebug = "debug"
	LoggerLevelInfo  = "info"
	LoggerLevelWarn  = "warn"
	LoggerLevelError = "err"
)

func (c *Config) ZapLoggerLevel() zap.AtomicLevel {
	switch c.LoggerLevel {
	case LoggerLevelDebug:
		return zap.NewAtomicLevelAt(zap.DebugLevel)
	case LoggerLevelInfo:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	case LoggerLevelWarn:
		return zap.NewAtomicLevelAt(zap.WarnLevel)
	case LoggerLevelError:
		return zap.NewAtomicLevelAt(zap.ErrorLevel)
	default:
		return zap.NewAtomicLevelAt(zap.InfoLevel)
	}
}

func (c *Config) HcLoggerLevel() hclog.Level {
	switch c.LoggerLevel {
	case LoggerLevelDebug:
		return hclog.Debug
	case LoggerLevelInfo:
		return hclog.Info
	case LoggerLevelWarn:
		return hclog.Warn
	case LoggerLevelError:
		return hclog.Error
	default:
		return hclog.Info

	}
}
