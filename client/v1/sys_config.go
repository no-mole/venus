package clientv1

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

type SysConfig interface {
	AddOrUpdateSysConfig(ctx context.Context, configName string, oidc *pbsysconfig.Oidc) (*pbsysconfig.SysConfig, error)
	ChangeOidcStatus(ctx context.Context, oidcStatus pbsysconfig.OidcStatus) (*pbsysconfig.SysConfig, error)
	LoadSysConfig(ctx context.Context, configName string) (*pbsysconfig.SysConfig, error)
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

func (s *sysConfig) AddOrUpdateSysConfig(ctx context.Context, configName string, oidc *pbsysconfig.Oidc) (*pbsysconfig.SysConfig, error) {
	s.logger.Debug("AddOrUpdateSysConfig", zap.String("configName", configName), zap.Any("oidc", oidc))
	return s.remote.AddOrUpdateSysConfig(ctx, &pbsysconfig.SysConfig{
		ConfigName: configName,
		Oidc:       oidc,
	}, s.callOpts...)
}

func (s *sysConfig) ChangeOidcStatus(ctx context.Context, oidcStatus pbsysconfig.OidcStatus) (*pbsysconfig.SysConfig, error) {
	s.logger.Debug("ChangeStatus", zap.String("OidcStatus", pbsysconfig.OidcStatus_name[int32(oidcStatus)]))
	return s.remote.ChangeOidcStatus(ctx, &pbsysconfig.ChangeOidcStatusRequest{Status: oidcStatus}, s.callOpts...)
}

func (s *sysConfig) LoadSysConfig(ctx context.Context, configName string) (*pbsysconfig.SysConfig, error) {
	s.logger.Debug("LoadSysConfig", zap.String("configName", configName))
	return s.remote.LoadSysConfig(ctx, &pbsysconfig.LoadSysConfigRequest{ConfigName: configName}, s.callOpts...)
}
