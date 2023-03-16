package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (r *Remote) Update(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	return r.client.Update(ctx, req)
}
