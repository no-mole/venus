package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbkv"
)

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
