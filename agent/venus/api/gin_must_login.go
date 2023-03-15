package api

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"net/http"
	"strings"

	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbsysconfig"
	"golang.org/x/oauth2"
)

const cookieKey = "venus-authorization"
const headerKey = "Authorization"

var Provider *oidc.Provider
var sysConfHash string
var Oauth2Config oauth2.Config

// MustLogin parse header and set token into context
// [Authorization: Bearer]
func MustLogin(s server.Server, aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		sysConf, err := s.LoadSysConfig(ctx, &emptypb.Empty{})
		if err != nil {
			return
		}
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		// 获取当前配置的hash值
		str := hashConfig(sysConf)
		if str != "" && str != sysConfHash {
			sysConfHash = str
			if sysConf != nil && sysConf.Oidc != nil && sysConf.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable {
				Provider, err = oidc.NewProvider(ctx, sysConf.Oidc.OauthServer)
				if err != nil {
					output.Json(ctx, err, nil)
					ctx.Abort()
					return
				}
				Oauth2Config = oauth2.Config{
					ClientID:     sysConf.Oidc.ClientId,
					ClientSecret: sysConf.Oidc.ClientSecret,
					Endpoint: oauth2.Endpoint{
						AuthURL:   Provider.Endpoint().AuthURL,
						TokenURL:  Provider.Endpoint().TokenURL,
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
			if sysConf != nil && sysConf.Oidc != nil && sysConf.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable {
				ctx.Redirect(http.StatusMovedPermanently, Oauth2Config.AuthCodeURL(""))
				return
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
	if config == nil {
		return ""
	}
	data, err := json.Marshal(config)
	if err != nil {
		return ""
	}
	shaData := sha256.Sum256(data)
	return base64.RawURLEncoding.EncodeToString(shaData[:])
}
