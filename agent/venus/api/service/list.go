package service

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbmicroservice"
)

// List
// @Summary 服务列表
// @Description qiuzhi.lu
// @Tags service
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param namespace path string true "命名空间"
// @Success 200 {object} pbmicroservice.ListServicesResponse.Services
// @Router /service/{namespace} [Get]
func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		namespace := ctx.Param("namespace")
		resp, err := s.ListServices(ctx, &pbmicroservice.ListServicesRequest{Namespace: namespace})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Services)
	}
}
