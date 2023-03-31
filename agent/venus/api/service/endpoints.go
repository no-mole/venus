package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbmicroservice"
)

// Endpoints
// @Summary 服务入口
// @Description qiuzhi.lu
// @Tags service
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param service_name path string true "服务名称"
// @Param service_version path string true "服务版本"
// @Success 200 {object} pbmicroservice.DiscoveryServiceResponse.Endpoints
// @Router /service/{namespace}/{service_name}/{service_version} [Get]
func Endpoints(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.Discovery(ctx, &pbmicroservice.ServiceInfo{
			Namespace:      ctx.Param("namespace"),
			ServiceName:    ctx.Param("service_name"),
			ServiceVersion: ctx.Param("service_version"),
		})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Endpoints)
	}
}
