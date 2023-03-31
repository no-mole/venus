package local

import (
	"context"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (l *Local) SysConfigUpdate(_ context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	data, err := codec.Encode(structs.SysConfigAddRequestType, req)
	if err != nil {
		return &pbsysconfig.SysConfig{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &pbsysconfig.SysConfig{}, f.Error()
	}
	return req, nil
}
