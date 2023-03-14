package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (r *Remote) ChangeOidcStatus(ctx context.Context, req *pbsysconfig.ChangeOidcStatusRequest) (*pbsysconfig.SysConfig, error) {
	return r.client.ChangeOidcStatus(ctx, req.Status)
}

func (r *Remote) AddOrUpdateSysConfig(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	return r.client.AddOrUpdateSysConfig(ctx, req.ConfigName, req.Oidc)
}
