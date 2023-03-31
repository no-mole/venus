package cluster

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List
// @Summary 获取cluster node 列表
// @Tags cluster
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} []pbcluster.Node
// @Router /cluster [get]
func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.Nodes(ctx, &emptypb.Empty{})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp.Nodes)
	}
}
