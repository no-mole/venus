package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
)

// Gen
// @Summary 新增accessKey
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security Basic
// @Param ak path string true "access_key"
// @Param object body pbaccesskey.AccessKeyInfo true "参数"
// @Success 200 {object} pbaccesskey.AccessKeyInfo
// @Router /access_key/{ak} [Post]
func Gen(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbaccesskey.AccessKeyInfo{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		req.Ak = ctx.Param("ak")
		resp, err := s.AccessKeyGen(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
