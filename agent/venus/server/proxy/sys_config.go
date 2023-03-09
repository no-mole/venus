package proxy

import (
	"context"

	"github.com/no-mole/venus/proto/pbconfig"
)

func (r *Remote) ChangeOidcStatus(ctx context.Context, req *pbconfig.ChangeOidcStatusRequest) (*pbconfig.Oidc, error) {
	return r.client.ChangeOidcStatus(ctx, req.Status)
}

func (r *Remote) AddOrUpdateOidc(ctx context.Context, req *pbconfig.Oidc) (*pbconfig.Oidc, error) {
	return r.client.AddOrUpdateOidc(ctx, req.OauthServer, req.ClientId, req.ClientSecret, req.RedirectUri)
}
