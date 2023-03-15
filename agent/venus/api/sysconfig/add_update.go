package sysconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbsysconfig"
)

// AddUpdate
// @Summary 新增或更新系统配置
// @Description qiuzhi.lu
// @Tags sys_config
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param object body pbsysconfig.SysConfig true "参数"
// @Success 200 {object} pbsysconfig.SysConfig
// @Router /sys_config [Post]
func AddUpdate(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		p := &pbsysconfig.SysConfig{}
		err := ctx.BindJSON(p)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		resp, err := s.AddOrUpdateSysConfig(ctx, p)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
