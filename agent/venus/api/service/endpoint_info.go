package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbmicroservice"
)

// EndpointInfo
// @Summary 服务入口详情
// @Tags service
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param service_name path string true "服务名称"
// @Param service_version path string true "服务版本"
// @Param service_endpoint path string true "服务版本"
// @Success 200 {object} pbmicroservice.ServiceEndpointInfo
// @Router /service/{namespace}/{service_name}/{service_version}/{service_endpoint} [Get]
func EndpointInfo(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.ServiceDesc(ctx, &pbmicroservice.ServiceInfo{
			Namespace:       ctx.Param("namespace"),
			ServiceName:     ctx.Param("service_name"),
			ServiceVersion:  ctx.Param("service_version"),
			ServiceEndpoint: ctx.Param("service_endpoint"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
