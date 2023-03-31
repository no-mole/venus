package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// HistoryList
// @Summary 获取某配置历史列表
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param key path string true "key"
// @Success 200 {object} pbkv.KvHistoryListResponse
// @Router /kv/history/{namespace}/{key} [Get]
func HistoryList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.KvHistoryList(ctx, &pbkv.KvHistoryListRequest{
			Namespace: ctx.Param("namespace"),
			Key:       ctx.Param("key"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Items)
	}
}
