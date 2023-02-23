package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// AccessKeyDel
// @Summary 命名空间下删除accessKey
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param ak path string true "accessKey"
// @Param object body pbnamespace.NamespaceAccessKeyInfo true "参数"
// @Success 200 {object} emptypb.Empty
// @Router /namespace/{namespace}/access_key/{ak} [Delete]
func AccessKeyDel(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbnamespace.NamespaceAccessKeyInfo{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		req.Namespace = ctx.Param("namespace")
		req.Ak = ctx.Param("ak")
		resp, err := s.NamespaceDelAccessKey(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
