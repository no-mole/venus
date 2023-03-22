package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"github.com/no-mole/venus/proto/pbcluster"
)

// Stats
// @Summary 获取cluster node 详情
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param id path string true "node id"
// @Success 200 {object} pbcluster.StatsResponse.Stats
// @Router /cluster/{id} [get]
func Stats(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.Stats(ctx, &pbcluster.StatsRequest{NodeId: ctx.Param("id")})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Stats)
	}
}
