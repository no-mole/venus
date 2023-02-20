package namespace

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List
// @Summary 命名空间列表
// @Description qiuzhi.lu@neptune
// @Tags namespace
// @Accept application/json
// @Produce application/json
// @Security Basic
// @Success 200 {object} pbnamespace.NamespacesListResponse
// @Router /namespace [Get]
func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.NamespacesList(ctx, &emptypb.Empty{})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
