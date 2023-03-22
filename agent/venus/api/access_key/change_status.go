package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
)

// ChangeStatus
// @Summary accessKey修改状态
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param ak path string true "access_key"
// @Param object body pbaccesskey.AccessKeyStatusChangeRequest true "参数"
// @Success 200 {object} emptypb.Empty
// @Router /access_key/{ak} [Put]
func ChangeStatus(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbaccesskey.AccessKeyStatusChangeRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		req.Ak = ctx.Param("ak")
		_, err = s.AccessKeyChangeStatus(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, nil)
	}
}
