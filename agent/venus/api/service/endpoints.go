package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbservice"
)

func Endpoints(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.DiscoveryOnce(ctx, &pbservice.ServiceInfo{
			Namespace:      ctx.Param("namespace"),
			ServiceName:    ctx.Param("service_name"),
			ServiceVersion: ctx.Param("service_version"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
