package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// ChangePassword
// @Summary 修改密码
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param object body pbuser.ChangePasswordRequest true "参数"
// @Success 200 {object} pbuser.UserInfo
// @Router /change_password [Put]
func ChangePassword(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbuser.ChangePasswordRequest{}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		resp, err := s.UserChangePassword(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
