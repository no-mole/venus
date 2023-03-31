package user

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/api/server"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Detail
// @Summary 获取用户详情
// @Tags user
// @Accept application/json
// @Produce application/json
// @Security ApiKeyAuth
// @Param uid path string true "用户uid"
// @Success 200 {object} pbuser.LoginResponse
// @Router /user/{uid} [get]
func Detail(s server.Server) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		resp, err := s.UserDetails(ctx, &emptypb.Empty{})
		if err != nil {
			output.Json(ctx, err, nil)
			return
		}
		output.Json(ctx, nil, resp)
	}
}
