package api

import (
	"fmt"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/proto/pbuser"
	"google.golang.org/protobuf/types/known/emptypb"
	"net/http"
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
// @Param object query CallbackParam true "入参"
// @Success 200
// @Router /oauth2/callback [Get]
func Callback(s server.Server, aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := &CallbackParam{}
		err := ctx.BindQuery(p)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		lock.RLock()
		defer lock.RUnlock()
		token, err := oauth2Config.Exchange(ctx, p.Code)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		userInfo, err := provider.UserInfo(ctx, oauth2Config.TokenSource(ctx, token))
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

		_, err = s.UserDetails(ctx, &emptypb.Empty{})
		if err != nil {
			if err != errors.ErrorGrpcUserNotExist {
				output.Json(ctx, err, nil)
				return
			}
			//todo 优化oidc用户数据同步问题
			adminToken := auth.NewJwtTokenWithClaim(time.Now().Add(30*time.Second), "venus", "venus", auth.TokenTypeAdministrator, nil)
			ctx.Set(auth.TokenContextKey, adminToken)
			_, err = s.UserRegister(ctx, &pbuser.UserInfo{
				Uid:  userInfo.Email,
				Name: c.Name,
			})
			if err != nil {
				output.Json(ctx, err, nil)
				return
			}
		}

		ctx.SetCookie(cookieKey, tokenString, 7200, "/", "", false, true)

		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}
		redirect := fmt.Sprintf("%s://%s/%s",
			scheme,
			ctx.Request.Host,
			"ui/index.html",
		)
		ctx.Redirect(http.StatusFound, redirect)
	}
}

type claim struct {
	Name string `json:"name"`
}
