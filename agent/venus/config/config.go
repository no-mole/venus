package config

import (
	"errors"
	"fmt"
	"net"
	"time"

	"github.com/hashicorp/go-hclog"
	"go.uber.org/zap"
)

type Config struct {
	NodeID string `json:"node_id" yaml:"node_id"`

	//data dir
	DaftDir string `json:"raft_dir"`

	//用于和其他主机通信的地址，默认为GrpcEndpoint
	LocalAddr string `json:"local_addr"`

	GrpcEndpoint string `json:"grpc_endpoint"`

	HttpEndpoint string `json:"http_endpoint"`

	//CertFile TLS cert file path
	CertFile string `json:"cert_file"`
	//KeyFile TLS key file path
	KeyFile string `json:"key_file"`

	BootstrapCluster bool `json:"bootstrap_cluster"`

	ApplyTimeout time.Duration `json:"apply_timeout"`

	JoinAddr string `json:"join_addr"`

	LoggerLevel LoggerLevel `json:"logger_level"`

	PeerToken string `json:"peer_token"`

	//TokenTimeout is the jwt token expired time
	TokenTimeout time.Duration `json:"token-timeout"`
}

func GetDefaultConfig() *Config {
	ip, _ := getClientIp()
	if ip == "" {
		ip = "127.0.0.1"
	}
	return &Config{
		NodeID:           "",
		DaftDir:          "",
		LocalAddr:        fmt.Sprintf("%s:6233", ip),
		GrpcEndpoint:     "0.0.0.0:6233",
		HttpEndpoint:     "0.0.0.0:7233",
		BootstrapCluster: false,
		ApplyTimeout:     1 * time.Second,
		JoinAddr:         "",
		LoggerLevel:      LoggerLevelInfo,
		TokenTimeout:     48 * time.Hour,
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
func getClientIp() (string, error) {
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return "", err
	}
	for _, address := range addrs {
		// 检查ip地址判断是否回环地址
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String(), nil
			}

		}
	}
	return "", errors.New("can not find the client ip address")
}
