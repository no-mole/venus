package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

// UserList
// @Summary 命名空间下用户列表
// @Description qiuzhi.lu
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Success 200 {object} pbnamespace.NamespaceUserListResponse
// @Router /namespace/{namespace}/user [Get]
func UserList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		resp, err := s.NamespaceUserList(ctx, &pbnamespace.NamespaceUserListRequest{NamespaceUid: namespace})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
