package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
)

// Logout
// @Summary 退出登陆
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200
// @Router /logout [Delete]
func Logout() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.SetCookie(cookieKey, "", -1, "/", "", false, true)
		output.Json(ctx, nil, "success")
	}
}
