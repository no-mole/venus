package local

import (
	"context"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbconfig"
)

func (l *Local) AddOrUpdateOidc(_ context.Context, req *pbconfig.Oidc) (*pbconfig.Oidc, error) {
	req.OidcStatus = pbconfig.OidcStatus_OidcStatusDisable
	data, err := codec.Encode(structs.OidcAddRequestType, req)
	if err != nil {
		return &pbconfig.Oidc{}, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return &pbconfig.Oidc{}, f.Error()
	}
	return req, nil
}

func (l *Local) ChangeOidcStatus(ctx context.Context, req *pbconfig.ChangeOidcStatusRequest) (*pbconfig.Oidc, error) {
	item := &pbconfig.Oidc{}
	buf, err := l.fsm.State().Get(ctx, []byte(structs.ConfigBucketName), []byte(structs.OidcConfigKey))
	if err != nil {
		return item, err
	}
	err = codec.Decode(buf, item)
	if err != nil {
		return item, err
	}
	item.OidcStatus = req.Status
	data, err := codec.Encode(structs.OidcAddRequestType, req)
	if err != nil {
		return item, err
	}
	f := l.r.Apply(data, l.applyTimeout)
	if f.Error() != nil {
		return item, f.Error()
	}
	return item, nil
}
