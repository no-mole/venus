package sysconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbsysconfig"
)

// Update
// @Summary 更新系统配置
// @Description qiuzhi.lu
// @Tags sys_config
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param object body pbsysconfig.SysConfig true "参数"
// @Success 200 {object} pbsysconfig.SysConfig
// @Router /sys_config [Post]
func Update(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := &pbsysconfig.SysConfig{}
		err := ctx.BindJSON(p)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		resp, err := s.Update(ctx, p)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
