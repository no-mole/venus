package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// Delete
// @Summary 删除用户
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param uid path string true "用户uid"
// @Router /user/{uid} [delete]
func Delete(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := s.UserUnregister(ctx, &pbuser.UserInfo{
			Uid:      ctx.Param("uid"),
			Password: "default",
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
	}
}
