package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbkv"
)

// Put
// @Summary 新增配置
// @Description qiuzhi.lu
// @Tags kv
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param key path string true "配置key"
// @Param object body pbkv.KVItem true "参数"
// @Success 200 {object} pbkv.KVItem
// @Router /kv/{namespace}/{key} [Put]
func Put(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbkv.KVItem{
			Namespace: ctx.Param("namespace"),
			Key:       ctx.Param("key"),
		}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item, err = s.AddKV(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, item)
	}
}
