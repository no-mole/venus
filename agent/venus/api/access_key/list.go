package access_key

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List
// @Summary accessKey列表
// @Description qiuzhi.lu
// @Tags access_key
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} pbaccesskey.AccessKeyListResponse
// @Router /access_key [Get]
func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.AccessKeyList(ctx, &emptypb.Empty{})
		if err != nil || len(resp.Items) == 0 {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
