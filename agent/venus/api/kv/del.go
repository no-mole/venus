package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// Del
// @Summary 删除配置
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param key path string true "配置key"
// @Success 200 {object} emptypb.Empty
// @Router /kv/{namespace}/{key} [Delete]
func Del(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.DelKey(ctx, &pbkv.DelKeyRequest{
			Namespace: ctx.Param("namespace"),
			Key:       ctx.Param("key"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
