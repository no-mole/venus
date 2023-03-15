package sysconfig

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Load
// @Summary 获取系统配置
// @Description qiuzhi.lu
// @Tags sys_config
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} pbsysconfig.SysConfig
// @Router /sys_config [Get]
func Load(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.LoadSysConfig(ctx, &emptypb.Empty{})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
