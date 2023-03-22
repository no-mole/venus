package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// Add
// @Summary 新增命名空间
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param object body pbnamespace.NamespaceItem true "参数"
// @Success 200 {object} pbnamespace.NamespaceItem
// @Router /namespace/{namespace} [Post]
func Add(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbnamespace.NamespaceItem{}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item.NamespaceUid = ctx.Param("namespace")
		item, err = s.NamespaceAdd(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, item)
	}

}
