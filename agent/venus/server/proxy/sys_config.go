package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (s *Remote) SysConfigUpdate(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	return s.client.SysConfigUpdate(ctx, req)
}
