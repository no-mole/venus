package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbsysconfig"
	"github.com/no-mole/venus/proto/pbuser"
)

// Login
// @Summary 登陆
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Param object body pbuser.LoginRequest true "参数"
// @Success 200 {object} pbuser.LoginResponse
// @Router /login [Post]
func Login(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbuser.LoginRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		resp, err := s.UserLogin(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		ctx.SetCookie(cookieKey, resp.AccessToken, int(resp.ExpiredIn), "/", "", false, true)
		output.Json(ctx, nil, resp)
	}
}

// OidcLogin
// @Summary oidc登陆
// @Description zhouguokang
// @Tags user
// @Accept application/json
// @Produce application/json
// @Router /oidc_login [Get]
func OidcLogin(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		conf := s.GetSysConfig()
		if conf == nil || conf.Oidc == nil || conf.Oidc.OidcStatus != pbsysconfig.OidcStatus_OidcStatusEnable || conf.Oidc.OauthServer == "" {
			output.Json(ctx, nil, nil)
			return
		}
		_, oauth2Conf, err := oidcLogin(ctx, conf)
		if err != nil {
			output.Json(ctx, nil, nil)
			return
		}
		output.Json(ctx, nil, oauth2Conf.AuthCodeURL("venus"))
	}
}
