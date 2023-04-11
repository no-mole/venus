package clientv1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

type SysConfig interface {
	SysConfigUpdate(ctx context.Context, sysConfig *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error)
	SysConfigGet(ctx context.Context) (*pbsysconfig.SysConfig, error)
}

func NewSysConfig(c *Client, logger *zap.Logger) SysConfig {
	return &sysConfig{
		remote:   pbsysconfig.NewSysConfigServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("sys_config"),
	}
}

var _ SysConfig = &sysConfig{}

type sysConfig struct {
	remote pbsysconfig.SysConfigServiceClient

	callOpts []grpc.CallOption

	logger *zap.Logger
}

func (s *sysConfig) SysConfigUpdate(ctx context.Context, sysConfig *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	s.logger.Debug("SysConfigUpdate", zap.Any("sysConfig", sysConfig))
	return s.remote.SysConfigUpdate(ctx, sysConfig, s.callOpts...)
}

func (s *sysConfig) SysConfigGet(ctx context.Context) (*pbsysconfig.SysConfig, error) {
	return s.remote.SysConfigGet(ctx, &emptypb.Empty{}, s.callOpts...)
}
