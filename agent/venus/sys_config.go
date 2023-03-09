package venus

import (
	"context"

	"github.com/no-mole/venus/agent/codec"
	"github.com/no-mole/venus/agent/structs"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/no-mole/venus/proto/pbconfig"
)

func (s *Server) AddOrUpdateOidc(ctx context.Context, req *pbconfig.Oidc) (*pbconfig.Oidc, error) {
	return s.server.AddOrUpdateOidc(ctx, req)
}

func (s *Server) ChangeOidcStatus(ctx context.Context, req *pbconfig.ChangeOidcStatusRequest) (*pbconfig.Oidc, error) {
	return s.server.ChangeOidcStatus(ctx, req)
}

func (s *Server) LoadOidcConfig(ctx context.Context, _ *emptypb.Empty) (*pbconfig.Oidc, error) {
	item := &pbconfig.Oidc{}
	buf, err := s.fsm.State().Get(ctx, []byte(structs.ConfigBucketName), []byte(structs.OidcConfigKey))
	if err != nil {
		return item, err
	}
	err = codec.Decode(buf, item)
	if err != nil {
		return item, err
	}
	return item, nil
}
