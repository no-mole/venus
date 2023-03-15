package local

import (
	"context"

	"github.com/no-mole/venus/agent/errors"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (l *Local) AddOrUpdateSysConfig(_ context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
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

func (l *Local) ChangeOidcStatus(ctx context.Context, req *pbsysconfig.ChangeOidcStatusRequest) (*pbsysconfig.SysConfig, error) {
	item := &pbsysconfig.SysConfig{}
	buf, err := l.fsm.State().Get(ctx, []byte(structs.SysConfigBucketName), []byte(structs.SysConfigBucketName))
	if err != nil {
		return item, err
	}
	err = codec.Decode(buf, item)
	if err != nil {
		return item, err
	}
	if item == nil || item.Oidc == nil {
		return item, errors.ErrorGrpcSysOrOidcConfigNotExist
	}
	item.Oidc.OidcStatus = req.Status
	data, err := codec.Encode(structs.SysConfigAddRequestType, req)
	if err != nil {
		return item, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return item, f.Error()
	}
	return item, nil
}
