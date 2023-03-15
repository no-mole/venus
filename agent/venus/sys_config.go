package venus

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (s *Server) AddOrUpdateSysConfig(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	return s.server.AddOrUpdateSysConfig(ctx, req)
}

func (s *Server) ChangeOidcStatus(ctx context.Context, req *pbsysconfig.ChangeOidcStatusRequest) (*pbsysconfig.SysConfig, error) {
	return s.server.ChangeOidcStatus(ctx, req)
}

var SysConfig *pbsysconfig.SysConfig

func (s *Server) LoadSysConfig(ctx context.Context, _ *emptypb.Empty) (*pbsysconfig.SysConfig, error) {
	item := &pbsysconfig.SysConfig{}

	if SysConfig != nil && SysConfig.Oidc != nil && SysConfig.Oidc.OidcStatus != pbsysconfig.OidcStatus_OidcStatusNil {
		return SysConfig, nil
	}
	buf, err := s.fsm.State().Get(ctx, []byte(structs.SysConfigBucketName), []byte(structs.SysConfigBucketName))
	if err != nil {
		return item, err
	}
	err = codec.Decode(buf, item)
	if err != nil {
		return item, err
	}
	SysConfig = item
	return item, nil
}
