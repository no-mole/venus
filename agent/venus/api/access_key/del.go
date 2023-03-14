package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
)

// Del
// @Summary 删除accessKey
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param ak path string true "access_key"
// @Success 200 {object} emptypb.Empty
// @Router /access_key/{ak} [Delete]
func Del(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, err := s.AccessKeyDel(ctx, &pbaccesskey.AccessKeyDelRequest{Ak: ctx.Param("ak")})
		output.Json(ctx, err, nil)
	}
}
