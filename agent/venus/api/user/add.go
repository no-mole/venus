package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

func Add(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbuser.UserInfo{}
		err := ctx.ShouldBindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item, err = s.UserRegister(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, item)
	}
}
