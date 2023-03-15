package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// AccessKeyAdd
// @Summary 命名空间下新增accessKey
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param ak path string true "accessKey"
// @Success 200 {object} emptypb.Empty
// @Router /namespace/{namespace}/access_key/{ak} [Post]
func AccessKeyAdd(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.NamespaceAddAccessKey(ctx, &pbnamespace.NamespaceAccessKeyInfo{
			Ak:           ctx.Param("ak"),
			NamespaceUid: ctx.Param("namespace"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
