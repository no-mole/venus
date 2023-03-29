package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// HistoryDetail
// @Summary 获取某配置历史记录详情
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param key path string true "key"
// @Param version path string true "version" 版本号
// @Success 200 {object} pbkv.KVItem
// @Router /kv/history/{namespace}/{key}/{version} [Get]
func HistoryDetail(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.KvHistoryDetail(ctx, &pbkv.GetHistoryDetailRequest{
			Version:   ctx.Param("version"),
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
