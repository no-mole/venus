package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// Fetch
// @Summary 查看配置
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param key path string true "配置key"
// @Success 200 {object} pbkv.KVItem
// @Router /kv/{namespace}/{key} [Get]
func Fetch(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.FetchKey(ctx, &pbkv.FetchKeyRequest{
			Namespace: ctx.Param("namespace"),
			Key:       ctx.Param("key"),
		})
		output.Json(ctx, err, resp)
	}
}
