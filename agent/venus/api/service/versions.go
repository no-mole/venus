package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"github.com/no-mole/venus/proto/pbmicroservice"
)

// Versions
// @Summary 服务版本
// @Description qiuzhi.lu
// @Tags service
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Param service_name path string true "服务名称"
// @Success 200 {object} pbmicroservice.ListServiceVersionsResponse
// @Router /service/{namespace}/{service_name} [Get]
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
