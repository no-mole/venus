package venus

import (
	"context"
	"github.com/no-mole/venus/agent/errors"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbsysconfig"
)

func (s *Server) AddOrUpdateSysConfig(ctx context.Context, req *pbsysconfig.SysConfig) (*pbsysconfig.SysConfig, error) {
	isAdmin, err := s.authenticator.IsAdministratorContext(ctx)
	if err != nil {
		return &pbsysconfig.SysConfig{}, err
	}
	if !isAdmin {
		return &pbsysconfig.SysConfig{}, errors.ErrorGrpcPermissionDenied
	}
	return s.server.AddOrUpdateSysConfig(ctx, req)
}

func (s *Server) LoadSysConfig(ctx context.Context, _ *emptypb.Empty) (*pbsysconfig.SysConfig, error) {
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
