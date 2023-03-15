package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (r *Remote) AddOrUpdateSysConfig(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	return r.client.AddOrUpdateSysConfig(ctx, req)
}
