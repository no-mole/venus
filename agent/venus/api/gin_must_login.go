package api

import (
	"context"
	"encoding/json"
	"github.com/coreos/go-oidc/v3/oidc"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/proto/pbsysconfig"
	"golang.org/x/oauth2"
	"io"
	"net/http"
	"net/url"
	"strings"
)

const cookieKey = "venus-authorization"
const headerKey = "Authorization"

// MustLogin parse header and set token into context
// [Authorization: Bearer]
func MustLogin(aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		err := parseToken(ctx, aor)
		if err != nil {
			//跳登陆页面不
			output.Json(ctx, err, nil)
			ctx.Abort()
			return
		}
	}
}

func parseToken(ctx *gin.Context, aor auth.Authenticator) error {
	tokenString, _ := ctx.Cookie(cookieKey)
	if tokenString == "" {
		tokenString = strings.TrimPrefix(ctx.Request.Header.Get(headerKey), "Bearer ")
	}
	if len(tokenString) == 0 {
		return errors.ErrorNotLogin
	}
	tokenString = strings.Trim(tokenString, " ")
	jwtToken, err := aor.Parse(ctx, tokenString)
	if err != nil {
		return err
	}
	ctx.Set(auth.TokenContextKey, jwtToken)
	return nil
}

type providerJSON struct {
	Issuer      string   `json:"issuer"`
	AuthURL     string   `json:"authorization_endpoint"`
	TokenURL    string   `json:"token_endpoint"`
	JWKSURL     string   `json:"jwks_uri"`
	UserInfoURL string   `json:"userinfo_endpoint"`
	Algorithms  []string `json:"id_token_signing_alg_values_supported"`
}

func oidcLogin(ctx context.Context, conf *pbsysconfig.SysConfig) (provider *oidc.Provider, oauth2Conf *oauth2.Config, err error) {
	wellKnown := strings.TrimSuffix(conf.Oidc.OauthServer, "/") + "/.well-known/openid-configuration"
	resp, err := http.Get(wellKnown)
	if err != nil {
		return provider, oauth2Conf, err
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return provider, oauth2Conf, err
	}
	err = resp.Body.Close()
	if err != nil {
		return provider, oauth2Conf, err
	}
	openidConf := &providerJSON{}
	err = json.Unmarshal(body, openidConf)
	if err != nil {
		return provider, oauth2Conf, err
	}

	syncScheme(conf.Oidc.OauthServer, openidConf)
	providerConfig := oidc.ProviderConfig{
		IssuerURL:   openidConf.Issuer,
		AuthURL:     openidConf.AuthURL,
		TokenURL:    openidConf.TokenURL,
		UserInfoURL: openidConf.UserInfoURL,
		JWKSURL:     openidConf.JWKSURL,
		Algorithms:  openidConf.Algorithms,
	}
	provider = providerConfig.NewProvider(ctx)
	oauth2Conf = &oauth2.Config{
		ClientID:     conf.Oidc.ClientId,
		ClientSecret: conf.Oidc.ClientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  conf.Oidc.RedirectUri,
		Scopes:       []string{oidc.ScopeOpenID, "email"},
	}
	return provider, oauth2Conf, nil
}

func syncScheme(server string, eps *providerJSON) {
	u, _ := url.Parse(server)
	if u.Scheme == "http" {
		eps.Issuer = strings.Replace(eps.Issuer, "https", "http", 1)
		eps.AuthURL = strings.Replace(eps.AuthURL, "https", "http", 1)
		eps.TokenURL = strings.Replace(eps.TokenURL, "https", "http", 1)
		eps.JWKSURL = strings.Replace(eps.JWKSURL, "https", "http", 1)
		eps.UserInfoURL = strings.Replace(eps.UserInfoURL, "https", "http", 1)
	}
}
