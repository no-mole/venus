package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbsysconfig"
	"github.com/no-mole/venus/proto/pbuser"
	"net/http"
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
func Callback(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		scheme := "http"
		if ctx.Request.TLS != nil {
			scheme = "https"
		}
		redirect := fmt.Sprintf("%s://%s/%s",
			scheme,
			ctx.Request.Host,
			"ui/index.html",
		)

		p := &CallbackParam{}
		err := ctx.BindQuery(p)
		if err != nil {
			ctx.Redirect(http.StatusFound, redirect)
			return
		}
		conf := s.GetSysConfig()
		if conf == nil || conf.Oidc == nil || conf.Oidc.OidcStatus != pbsysconfig.OidcStatus_OidcStatusEnable || conf.Oidc.OauthServer == "" {
			ctx.Redirect(http.StatusFound, redirect)
			return
		}
		provider, oauth2Config, err := oidcLogin(ctx, conf)

		token, err := oauth2Config.Exchange(ctx, p.Code)
		if err != nil {
			ctx.Redirect(http.StatusFound, redirect)
			return
		}
		userInfo, err := provider.UserInfo(ctx, oauth2Config.TokenSource(ctx, token))
		if err != nil {
			ctx.Redirect(http.StatusFound, redirect)
			return
		}
		c := &claim{}
		err = userInfo.Claims(c)
		if err != nil {
			ctx.Redirect(http.StatusFound, redirect)
			return
		}

		resp, err := s.UserSync(&pbuser.UserInfo{
			Uid:  userInfo.Email,
			Name: c.Name,
		})

		if err != nil {
			ctx.Redirect(http.StatusFound, redirect)
			return
		}

		ctx.SetCookie(cookieKey, resp.AccessToken, int(resp.ExpiredIn), "/", "", false, true)
		ctx.Redirect(http.StatusFound, redirect)
	}
}

type claim struct {
	Name string `json:"name"`
}
