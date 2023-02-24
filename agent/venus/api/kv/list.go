package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// List
// @Summary 配置列表
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Success 200 {object} pbkv.ListKeysResponse
// @Router /kv/{namespace} [Get]
func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.ListKeys(ctx, &pbkv.ListKeysRequest{Namespace: ctx.Param("namespace")})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
