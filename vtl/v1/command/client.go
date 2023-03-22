package command

import (
	"context"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/spf13/viper"
	"strings"
)

const (
	EnvPrefix            = "VTL"
	Endpoint             = "endpoint"
	DialTimeout          = "dial-timeout"
	DialKeepaliveTime    = "dial-keepalive-time"
	DialKeepaliveTimeout = "dial-keepalive-timeout"
	MaxCallSendMsgSize   = "max-call-send-msg-size"
	MaxCallRecvMsgSize   = "max-call-recv-msg-size"
	UserName             = "username"
	RootPassword         = "root-password"
	PeerToken            = "peer-token"
	AccessKey            = "access-key"
	AccessKeySecret      = "access-key-secret"
)

func getClientConfigFromFlags() (clientv1.Config, error) {
	viper.SetEnvPrefix(EnvPrefix)
	viper.SetDefault("endpoint", "127.0.0.1:3333")
	viper.SetDefault("dial-timeout", "1s")
	viper.SetDefault("dial-keepalive-time", "10s")
	viper.SetDefault("dial-keepalive-timeout", "1s")
	cfg := clientv1.Config{
		Endpoints:            strings.Split(viper.GetString(Endpoint), ","),
		DialTimeout:          viper.GetDuration(DialTimeout),
		DialKeepAliveTime:    viper.GetDuration(DialKeepaliveTime),
		DialKeepAliveTimeout: viper.GetDuration(DialKeepaliveTimeout),
		MaxCallSendMsgSize:   viper.GetInt(MaxCallSendMsgSize),
		MaxCallRecvMsgSize:   viper.GetInt(MaxCallRecvMsgSize),
		Username:             viper.GetString(UserName),
		Password:             viper.GetString(RootPassword),
		PeerToken:            viper.GetString(PeerToken),
		AccessKey:            viper.GetString(AccessKey),
		AccessKeySecret:      viper.GetString(AccessKeySecret),
		Context:              context.Background(),
	}
	return cfg, nil
}

func getClient(cfg clientv1.Config) (*clientv1.Client, error) {
	client, err := clientv1.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}

func getClientFromFlags() (*clientv1.Client, error) {
	cfg, err := getClientConfigFromFlags()
	if err != nil {
		return nil, err
	}
	client, err := getClient(cfg)
	if err != nil {
		return nil, err
	}
	return client, nil
}
