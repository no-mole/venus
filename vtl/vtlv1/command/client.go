package command

import (
	"context"
	clientv1 "github.com/no-mole/venus/client/v1"
	"github.com/spf13/viper"
	"strings"
)

const (
	EnvPrefix = "VTL"
)

func getClientConfigFromFlags() (clientv1.Config, error) {
	viper.SetEnvPrefix(EnvPrefix)
	viper.SetDefault("endpoint", "127.0.0.1:3333")
	viper.SetDefault("dial-timeout", "1s")
	viper.SetDefault("dial-keepalive-time", "10s")
	viper.SetDefault("dial-keepalive-timeout", "1s")
	cfg := clientv1.Config{
		Endpoints:            strings.Split(viper.GetString("endpoint"), ","),
		DialTimeout:          viper.GetDuration("dial-timeout"),
		DialKeepAliveTime:    viper.GetDuration("dial-keepalive-time"),
		DialKeepAliveTimeout: viper.GetDuration("dial-keepalive-timeout"),
		MaxCallSendMsgSize:   viper.GetInt("max-call-send-msg-size"),
		MaxCallRecvMsgSize:   viper.GetInt("max-call-recv-msg-size"),
		Username:             viper.GetString("username"),
		Password:             viper.GetString("password"),
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
