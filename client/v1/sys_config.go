package clientv1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

type SysConfig interface {
	AddOrUpdateSysConfig(ctx context.Context, sysConfig *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error)
	ChangeOidcStatus(ctx context.Context, oidcStatus pbsysconfig.OidcStatus) (*pbsysconfig.SysConfig, error)
	LoadSysConfig(ctx context.Context) (*pbsysconfig.SysConfig, error)
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

func (s *sysConfig) AddOrUpdateSysConfig(ctx context.Context, sysConfig *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	s.logger.Debug("AddOrUpdateSysConfig", zap.Any("sysConfig", sysConfig))
	return s.remote.AddOrUpdateSysConfig(ctx, sysConfig, s.callOpts...)
}

func (s *sysConfig) ChangeOidcStatus(ctx context.Context, oidcStatus pbsysconfig.OidcStatus) (*pbsysconfig.SysConfig, error) {
	s.logger.Debug("ChangeStatus", zap.String("OidcStatus", pbsysconfig.OidcStatus_name[int32(oidcStatus)]))
	return s.remote.ChangeOidcStatus(ctx, &pbsysconfig.ChangeOidcStatusRequest{Status: oidcStatus}, s.callOpts...)
}

func (s *sysConfig) LoadSysConfig(ctx context.Context) (*pbsysconfig.SysConfig, error) {
	return s.remote.LoadSysConfig(ctx, &emptypb.Empty{}, s.callOpts...)
}
