package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbmicroservice"
)

func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		resp, err := s.ListServices(ctx, &pbmicroservice.ListServicesRequest{Namespace: namespace})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
