package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbuser"
)

func ChangePassword(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbuser.UserInfo{}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item.Uid = ctx.Param("uid")
		item, err = s.UserRegister(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, item)
	}
}
