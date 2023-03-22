package command

import (
	"context"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/spf13/viper"
	"strings"
)

const (
	EnvPrefix                    = "VTL"
	FlagNameEndpoint             = "endpoint"
	FlagNameDialTimeout          = "dial-timeout"
	FlagNameDialKeepaliveTime    = "dial-keepalive-time"
	FlagNameDialKeepaliveTimeout = "dial-keepalive-timeout"
	FlagNameMaxCallSendMsgSize   = "max-call-send-msg-size"
	FlagNameMaxCallRecvMsgSize   = "max-call-recv-msg-size"
	FlagNameUserName             = "username"
	FlagNameRootPassword         = "root-password"
	FlagNamePeerToken            = "peer-token"
	FlagNameAccessKey            = "access-key"
	FlagNameAccessKeySecret      = "access-key-secret"
)

func getClientConfigFromFlags() (clientv1.Config, error) {
	viper.SetEnvPrefix(EnvPrefix)
	viper.SetDefault("endpoint", "127.0.0.1:3333")
	viper.SetDefault("dial-timeout", "1s")
	viper.SetDefault("dial-keepalive-time", "10s")
	viper.SetDefault("dial-keepalive-timeout", "1s")
	cfg := clientv1.Config{
		Endpoints:            strings.Split(viper.GetString(FlagNameEndpoint), ","),
		DialTimeout:          viper.GetDuration(FlagNameDialTimeout),
		DialKeepAliveTime:    viper.GetDuration(FlagNameDialKeepaliveTime),
		DialKeepAliveTimeout: viper.GetDuration(FlagNameDialKeepaliveTimeout),
		MaxCallSendMsgSize:   viper.GetInt(FlagNameMaxCallSendMsgSize),
		MaxCallRecvMsgSize:   viper.GetInt(FlagNameMaxCallRecvMsgSize),
		Username:             viper.GetString(FlagNameUserName),
		Password:             viper.GetString(FlagNameRootPassword),
		PeerToken:            viper.GetString(FlagNamePeerToken),
		AccessKey:            viper.GetString(FlagNameAccessKey),
		AccessKeySecret:      viper.GetString(FlagNameAccessKeySecret),
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
