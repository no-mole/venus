package venus

import (
	"context"
	"github.com/coreos/go-oidc"
	"github.com/no-mole/venus/agent/errors"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (s *Server) Update(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx)
	if err != nil {
		return &pbsysconfig.SysConfig{}, err
	}
	if !isAdmin {
		return &pbsysconfig.SysConfig{}, errors.ErrorGrpcPermissionDenied
	}
	if req.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable {
		_, err = oidc.NewProvider(ctx, req.Oidc.OauthServer)
		if err != nil {
			return req, errors.ToGrpcError(err)
		}
	}
	return s.server.Update(ctx, req)
}

func (s *Server) Get(ctx context.Context, _ *emptypb.Empty) (*pbsysconfig.SysConfig, error) {
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx)
	if err != nil {
		return &pbsysconfig.SysConfig{}, err
	}
	if !isAdmin {
		return &pbsysconfig.SysConfig{}, errors.ErrorGrpcPermissionDenied
	}
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.sysConfig, nil
}

func (s *Server) loadSysConf(ctx context.Context) (*pbsysconfig.SysConfig, error) {
	s.rwLock.Lock()
	defer s.rwLock.Unlock()
	item := &pbsysconfig.SysConfig{}
	buf, err := s.fsm.State().Get(ctx, []byte(structs.SysConfigBucketName), []byte(structs.SysConfigBucketName))
	if err != nil {
		return item, err
	}
	err = codec.Decode(buf, item)
	if err != nil {
		return item, err
	}
	return item, nil
}

func (s *Server) GetSysConfig() *pbsysconfig.SysConfig {
	s.rwLock.RLock()
	defer s.rwLock.RUnlock()
	return s.sysConfig
}
