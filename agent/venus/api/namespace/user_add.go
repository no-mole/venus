package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

func UserAdd(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbnamespace.NamespaceUserInfo{}
		err := ctx.ShouldBindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		_, err = s.NamespaceAddUser(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, "success")
	}
}
