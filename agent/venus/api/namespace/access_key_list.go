package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// AccessKeyList
// @Summary 命名空间下accessKey列表
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Success 200 {object} pbnamespace.NamespaceAccessKeyListResponse
// @Router /namespace/{namespace}/access_key [Get]
func AccessKeyList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.NamespaceAccessKeyList(ctx, &pbnamespace.NamespaceAccessKeyListRequest{Namespace: ctx.Param("namespace")})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
