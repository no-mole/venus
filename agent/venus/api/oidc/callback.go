package oidc

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/server"
)

// Callback
// @Summary 登陆接口
// @Description qiuzhi.lu
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param code path string true "auth code"
// @Success 200
// @Router /auth/callback/{code} [Get]
func Callback(s server.Server, aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authCode := ctx.Param("code")
		token, err := api.Oauth2Config.Exchange(ctx, authCode)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		userInfo, err := api.Provider.UserInfo(ctx, api.Oauth2Config.TokenSource(ctx, token))
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		c := &claim{}
		err = userInfo.Claims(c)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}

		tk := auth.NewJwtTokenWithClaim(time.Now().Add(2*time.Hour), userInfo.Email, c.Name, auth.TokenTypeUser, nil)
		tokenString, err := aor.Sign(ctx, tk)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		ctx.SetCookie("venus-authorization", tokenString, 7200, "/", "", false, true)
	}
}

type claim struct {
	Name string `json:"name"`
}
