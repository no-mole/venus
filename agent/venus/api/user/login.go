package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// Login
// @Summary 登陆
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param uid path string true "用户uid"
// @Param object body pbuser.LoginRequest true "参数"
// @Success 200 {object} LoginResp
// @Router /user/login/{uid} [Post]
func Login(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbuser.LoginRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		req.Uid = ctx.Param("uid")
		resp, err := s.UserLogin(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		userNamespaceListResp, err := s.UserNamespaceList(ctx, &pbuser.UserNamespaceListRequest{Uid: resp.Uid})
		if err != nil {
			output.Json(ctx, err, nil)
		}
		ctx.SetCookie("venus-authorization", resp.AccessToken, int(resp.ExpiredIn), "/", "", false, true)
		output.Json(ctx, nil, &LoginResp{
			UserInfo:          resp,
			UserNamespaceList: userNamespaceListResp,
		})
	}
}

type LoginResp struct {
	UserInfo          *pbuser.LoginResponse             `json:"user_info"`
	UserNamespaceList *pbuser.UserNamespaceListResponse `json:"user_namespace_list"`
}
