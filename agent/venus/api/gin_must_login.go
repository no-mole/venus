package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
	"strings"
)

// MustLogin parse header and set token into context
// [Authorization: Bearer]
func MustLogin(aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		token := strings.Trim(ctx.Request.Header.Get("Authorization"), " ")
		if (len(token) == 0) || (!strings.HasPrefix(token, "Bearer ")) {
			output.Json(ctx, errors.ErrorGrpcNotLogin, nil)
			ctx.Abort()
			return
		}
		token = strings.TrimPrefix(token, "Bearer ")
		jwtToken, err := aor.Parse(ctx, token)
		if err != nil {
			output.Json(ctx, err, nil)
			ctx.Abort()
			return
		}
		ctx.Set(auth.TokenContextKey, jwtToken)
	}
}
