package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbnamespace"
)

func Add(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		item := &pbnamespace.NamespaceItem{}
		err := ctx.BindJSON(item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		item.NamespaceEn = ctx.Param("namespace")
		item, err = s.NamespaceAdd(ctx, item)
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, item)
	}

}
