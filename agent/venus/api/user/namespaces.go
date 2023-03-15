package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

// NamespaceList
// @Summary user namespace 列表
// @Description qiuzhi.lu
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param uid path string true "access_key"
// @Success 200 {object} pbnamespace.NamespaceAccessKeyListResponse
// @Router /user/{uid}/namespace [Get]
func NamespaceList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.UserNamespaceList(ctx, &pbuser.UserNamespaceListRequest{
			Uid: ctx.Param("uid"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Items)
	}
}
