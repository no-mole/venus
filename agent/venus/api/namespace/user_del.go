package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

func UserDel(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbnamespace.NamespaceUserInfo{}
		err := ctx.ShouldBindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		_, err = s.NamespaceDelUser(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, "success")
	}
}
