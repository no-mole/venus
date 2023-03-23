package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (s *Remote) Update(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	return s.client.Update(ctx, req)
}
