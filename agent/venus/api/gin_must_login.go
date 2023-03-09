package api

import (
	"context"
	"net/http"
	"strings"

	"github.com/no-mole/venus/proto/pbconfig"

	"github.com/no-mole/venus/agent/venus/server"
	"google.golang.org/protobuf/types/known/emptypb"

	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
)

const cookieKey = "venus-authorization"
const headerKey = "Authorization"

var s server.Server

// MustLogin parse header and set token into context
// [Authorization: Bearer]
func MustLogin(aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		oidcConf, err := s.LoadOidcConfig(context.Background(), &emptypb.Empty{})
		if err != nil {
			return
		}
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		tokenString, _ := ctx.Cookie(cookieKey)
		if tokenString == "" {
			tokenString = strings.TrimPrefix(ctx.Request.Header.Get(headerKey), "Bearer ")
		}
		if len(tokenString) == 0 {
			if oidcConf.OidcStatus == pbconfig.OidcStatus_OidcStatusEnable {
				// todo uri
				ctx.Redirect(http.StatusMovedPermanently, "")
			}
			output.Json(ctx, errors.ErrorGrpcNotLogin, nil)
			ctx.Abort()
			return
		}
		tokenString = strings.Trim(tokenString, " ")
		jwtToken, err := aor.Parse(ctx, tokenString)
		if err != nil {
			output.Json(ctx, err, nil)
			ctx.Abort()
			return
		}
		ctx.Set(auth.TokenContextKey, jwtToken)
	}
}
