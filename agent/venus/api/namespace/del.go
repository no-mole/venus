package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// Del
// @Summary 删除命名空间
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security Basic
// @Param namespace path string true "命名空间"
// @Success 200 {object} emptypb.Empty
// @Router /namespace/{namespace} [Delete]
func Del(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		_, err := s.NamespaceDel(ctx, &pbnamespace.NamespaceDelRequest{Namespace: namespace})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, "success")
	}
}
