package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// ResetPassword
// @Summary 重置密码
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param uid path string true "用户uid"
// @Success 200 {object} pbuser.UserInfo
// @Router /user/{uid} [Put]
func ResetPassword(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.UserResetPassword(ctx, &pbuser.ResetPasswordRequest{Uid: ctx.Param("uid")})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
