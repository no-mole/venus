package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// ChangePassword
// @Summary 修改密码
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security Basic
// @Param uid path string true "用户uid"
// @Param object body pbuser.UserInfo true "参数"
// @Success 200 {object} pbuser.UserInfo
// @Router /user/{uid} [Put]
func ChangePassword(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbuser.UserInfo{}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item.Uid = ctx.Param("uid")
		item, err = s.UserRegister(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, item)
	}
}
