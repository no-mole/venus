package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbmicroservice"
)

func Versions(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		serviceName := ctx.Param("service_name")
		resp, err := s.ListServiceVersions(ctx, &pbmicroservice.ListServiceVersionsRequest{
			Namespace:   namespace,
			ServiceName: serviceName,
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
