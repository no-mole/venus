package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// UserDel
// @Summary 命名空间下删除用户
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param uid path string true "用户uid"
// @Success 200 {object} emptypb.Empty
// @Router /namespace/{namespace}/user/{uid} [Delete]
func UserDel(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := s.NamespaceDelUser(ctx, &pbnamespace.NamespaceUserDelRequest{
			NamespaceUid: ctx.Param("namespace"),
			Uid:          ctx.Param("uid"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, nil)
	}
}
