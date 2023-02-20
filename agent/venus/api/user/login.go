package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// Login
// @Summary 登陆
// @Description qiuzhi.lu@neptune
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security Basic
// @Param uid path string true "用户uid"
// @Param object body pbuser.LoginRequest true "参数"
// @Success 200 {object} pbuser.UserInfo
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
		output.Json(ctx, nil, resp)
	}
}
