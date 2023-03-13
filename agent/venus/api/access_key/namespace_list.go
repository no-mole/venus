package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
)

// NamespaceList
// @Summary accessKey namespace 列表
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param ak path string true "access_key"
// @Success 200 {object} pbaccesskey.AccessKeyNamespaceListResponse
// @Router /access_key/{ak}/namespace [Get]
func NamespaceList(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.AccessKeyNamespaceList(ctx, &pbaccesskey.AccessKeyNamespaceListRequest{
			Ak: ctx.Param("ak"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
