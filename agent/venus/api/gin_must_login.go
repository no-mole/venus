package api

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"golang.org/x/oauth2"

	"github.com/coreos/go-oidc"

	"github.com/no-mole/venus/agent/structs"

	"github.com/no-mole/venus/proto/pbsysconfig"

	"github.com/no-mole/venus/agent/venus/server"

	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
)

const cookieKey = "venus-authorization"
const headerKey = "Authorization"
const issuerUrl = "https://smart.gitlab.biomind.com.cn"
const authorizationEndpoint = "https://smart.gitlab.biomind.com.cn/oauth/authorize"
const tokenEndpoint = "https://smart.gitlab.biomind.com.cn/oauth/token"

var s server.Server
var Provider *oidc.Provider
var oidcConfHash string
var Oauth2Config oauth2.Config

// MustLogin parse header and set token into context
// [Authorization: Bearer]
func MustLogin(aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sysConf, err := s.LoadSysConfig(context.Background(), &pbsysconfig.LoadSysConfigRequest{ConfigName: structs.OidcConfigKey})
		if err != nil {
			return
		}
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		// 获取当前配置的hash值
		str := hashConfig(sysConf)
		if str != "" && str != oidcConfHash {
			oidcConfHash = str
			if sysConf.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable {
				Provider, err = oidc.NewProvider(ctx, issuerUrl)
				if err != nil {
					output.Json(ctx, err, nil)
					ctx.Abort()
					return
				}
				Oauth2Config = oauth2.Config{
					ClientID:     sysConf.Oidc.ClientId,
					ClientSecret: sysConf.Oidc.ClientSecret,
					Endpoint: oauth2.Endpoint{
						AuthURL:   authorizationEndpoint,
						TokenURL:  tokenEndpoint,
						AuthStyle: 0,
					},
					RedirectURL: sysConf.Oidc.RedirectUri,
					// todo
					Scopes: nil,
				}
			}
		}
		tokenString, _ := ctx.Cookie(cookieKey)
		if tokenString == "" {
			tokenString = strings.TrimPrefix(ctx.Request.Header.Get(headerKey), "Bearer ")
		}

		if len(tokenString) == 0 {
			if sysConf.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable {
				ctx.Redirect(http.StatusMovedPermanently, Oauth2Config.AuthCodeURL(""))
			}
			output.Json(ctx, errors.ErrorGrpcNotLogin, nil)
			ctx.Abort()
			return
		}
		tokenString = strings.Trim(tokenString, " ")
		jwtToken, err := aor.Parse(ctx, tokenString)
		if err != nil {
			output.Json(ctx, err, nil)
			ctx.Abort()
			return
		}
		ctx.Set(auth.TokenContextKey, jwtToken)
	}
}

func hashConfig(config *pbsysconfig.SysConfig) string {
	data, err := json.Marshal(config.Oidc)
	if err != nil {
		return ""
	}
	shaData := sha256.Sum256(data)
	return base64.RawURLEncoding.EncodeToString(shaData[:])
}
