package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

func Login(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req := &pbuser.LoginRequest{}
		err := ctx.BindJSON(req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		req.Uid = ctx.Param("uid")
		resp, err := s.UserLogin(ctx, req)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
