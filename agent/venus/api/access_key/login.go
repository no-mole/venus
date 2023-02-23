package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbaccesskey"
)

// Login
// @Summary accessKey登陆
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param ak path string true "access_key"
// @Param object body pbaccesskey.AccessKeyLoginRequest true "参数"
// @Success 200 {object} pbaccesskey.AccessKeyLoginResponse
// @Router /access_key/login/{ak} [Post]
func Login(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbaccesskey.AccessKeyLoginRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		req.Ak = ctx.Param("ak")
		resp, err := s.AccessKeyLogin(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
