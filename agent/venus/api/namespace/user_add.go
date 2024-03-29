package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// UserAdd
// @Summary 命名空间下新增用户
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param uid path string true "用户uid"
// @Param object body pbnamespace.NamespaceUserInfo true "参数"
// @Success 200 {object} emptypb.Empty
// @Router /namespace/{namespace}/user/{uid} [Post]
func UserAdd(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbnamespace.NamespaceUserInfo{}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item.NamespaceUid = ctx.Param("namespace")
		item.Uid = ctx.Param("uid")
		_, err = s.NamespaceAddUser(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, nil)
	}
}
