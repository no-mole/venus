package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// Gen
// @Summary 创建accessKey
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param alias path string true "access key alias"
// @Param object body pbaccesskey.AccessKeyInfo true "参数"
// @Success 200 {object} pbaccesskey.AccessKeyInfo
// @Router /access_key/{namespace}/{alias} [Post]
func Gen(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		alias := ctx.Param("alias")
		resp, err := s.AccessKeyGen(ctx, &pbaccesskey.AccessKeyInfo{Alias: alias})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		_, err = s.NamespaceAddAccessKey(ctx, &pbnamespace.NamespaceAccessKeyInfo{
			Ak:           resp.Ak,
			NamespaceUid: ctx.Param("namespace"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, err, resp)
	}
}
