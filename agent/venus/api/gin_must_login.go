package api

import (
	"context"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"github.com/no-mole/venus/agent/venus/api/server"
	"net/http"
	"strings"
	"sync"

	"github.com/coreos/go-oidc"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/proto/pbsysconfig"
	"golang.org/x/oauth2"
)

const cookieKey = "venus-authorization"
const headerKey = "Authorization"

var (
	provider     *oidc.Provider
	sysConfHash  string
	oauth2Config oauth2.Config
	lock         sync.RWMutex
)

func OIDCMustLogin(s server.Server, aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		shouldOidcLogin, err := oidcLogin(ctx, s)
		if err != nil {
			output.Json(ctx, err, nil)
			ctx.Abort()
			return
		}
		if !shouldOidcLogin {
			return
		}
		err = parseToken(ctx, aor)
		if err != nil {
			ctx.Redirect(http.StatusFound, oauth2Config.AuthCodeURL("venus"))
			ctx.Abort()
			return
		}
	}
}

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

func oidcLogin(ctx context.Context, s server.Server) (bool, error) {
	sysConf := s.GetSysConfig()
	var err error
	// 获取当前配置的hash值
	str := hashConfig(sysConf)
	if str != "" && str != sysConfHash {
		lock.Lock()
		defer lock.Unlock()
		if str != sysConfHash {
			sysConfHash = str
			if sysConf != nil && sysConf.Oidc != nil && sysConf.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable {
				provider, err = oidc.NewProvider(ctx, sysConf.Oidc.OauthServer)
				if err != nil {
					return false, err
				}
				oauth2Config = oauth2.Config{
					ClientID:     sysConf.Oidc.ClientId,
					ClientSecret: sysConf.Oidc.ClientSecret,
					Endpoint:     provider.Endpoint(),
					RedirectURL:  sysConf.Oidc.RedirectUri,
					Scopes:       []string{oidc.ScopeOpenID, "email"},
				}
			}
		}
	}
	return sysConf != nil && sysConf.Oidc != nil && sysConf.Oidc.OidcStatus == pbsysconfig.OidcStatus_OidcStatusEnable, nil
}

func setCookie(tokenString string) {

}
