package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// WatchKeyClientList
// @Summary 获取配置项监听客户端信息
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "namespace"
// @Param key path string true "key"
// @Success 200 {object} pbkv.WatchKeyClientListResponse
// @Router /kv/watch/{namespace}/{key} [Get]
func WatchKeyClientList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.WatchKeyClientList(ctx, &pbkv.WatchKeyClientListRequest{
			Namespace: ctx.Param("namespace"),
			Key:       ctx.Param("key"),
			Diffusion: true,
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Items)
	}
}
