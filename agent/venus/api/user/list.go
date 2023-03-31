package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

// List
// @Summary 用户列表
// @Description by zgk
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Success 200 {object} pbuser.UserListResponse
// @Router /user [get]
func List(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.UserList(ctx, &emptypb.Empty{})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
