package kv

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbkv"
)

func Fetch(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.FetchKey(ctx, &pbkv.FetchKeyRequest{
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
