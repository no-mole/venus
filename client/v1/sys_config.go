package clientv1

import (
	"context"

	"google.golang.org/protobuf/types/known/emptypb"

	"go.uber.org/zap"
	"google.golang.org/grpc"

	"github.com/no-mole/venus/proto/pbconfig"
)

type Oidc interface {
	AddOrUpdateOidc(ctx context.Context, oauthServer, clientId, clientSecret, redirectUri string) (*pbconfig.Oidc, error)
	ChangeOidcStatus(ctx context.Context, oidcStatus pbconfig.OidcStatus) (*pbconfig.Oidc, error)
	LoadOidcConfig(ctx context.Context) (*pbconfig.Oidc, error)
}

func NewOidc(c *Client, logger *zap.Logger) Oidc {
	return &oidc{
		remote:   pbconfig.NewConfigServiceClient(c.conn),
		callOpts: c.callOpts,
		logger:   logger.Named("oidc"),
	}
}

var _ Oidc = &oidc{}

type oidc struct {
	remote pbconfig.ConfigServiceClient

	callOpts []grpc.CallOption

	logger *zap.Logger
}

func (o *oidc) AddOrUpdateOidc(ctx context.Context, oauthServer, clientId, clientSecret, redirectUri string) (*pbconfig.Oidc, error) {
	o.logger.Debug("AddOrUpdateOidc", zap.String("oauthServer", oauthServer), zap.String("clientId", clientId),
		zap.String("clientSecret", clientSecret), zap.String("redirectUri", redirectUri))
	return o.remote.AddOrUpdateOidc(ctx, &pbconfig.Oidc{
		OauthServer:  oauthServer,
		ClientId:     clientId,
		ClientSecret: clientSecret,
		RedirectUri:  redirectUri,
		OidcStatus:   pbconfig.OidcStatus_OidcStatusDisable,
	}, o.callOpts...)
}

func (o *oidc) ChangeOidcStatus(ctx context.Context, oidcStatus pbconfig.OidcStatus) (*pbconfig.Oidc, error) {
	o.logger.Debug("ChangeStatus", zap.String("OidcStatus", pbconfig.OidcStatus_name[int32(oidcStatus)]))
	return o.remote.ChangeOidcStatus(ctx, &pbconfig.ChangeOidcStatusRequest{Status: oidcStatus}, o.callOpts...)
}

func (o *oidc) LoadOidcConfig(ctx context.Context) (*pbconfig.Oidc, error) {
	return o.remote.LoadOidcConfig(ctx, &emptypb.Empty{}, o.callOpts...)
}
