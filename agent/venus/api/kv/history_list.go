package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// HistoryList
// @Summary 配置管理历史列表
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Success 200 {object} pbkv.HistoryListResponse
// @Router /kv/history/{namespace} [Get]
func HistoryList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.HistoryList(ctx, &pbkv.HistoryListRequest{Namespace: ctx.Param("namespace")})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Items)
	}
}
