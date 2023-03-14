package api

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
	"github.com/no-mole/venus/agent/venus/server"
)

type CallbackParam struct {
	Code string `json:"code" form:"code" binding:"required"`
}

// Callback
// @Summary 登陆接口
// @Description qiuzhi.lu
// @Tags auth
// @Accept application/json
// @Produce application/json
// @Param object body CallbackParam true "入参"
// @Success 200
// @Router /auth/callback/{code} [Get]
func Callback(s server.Server, aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := &CallbackParam{}
		err := ctx.BindQuery(p)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		token, err := Oauth2Config.Exchange(ctx, p.Code)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		userInfo, err := Provider.UserInfo(ctx, Oauth2Config.TokenSource(ctx, token))
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
		ctx.SetCookie(cookieKey, tokenString, 7200, "/", "", false, true)
	}
}

type claim struct {
	Name string `json:"name"`
}