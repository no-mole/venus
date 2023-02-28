package api

import (
	"github.com/gin-gonic/gin"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/output"
	"github.com/no-mole/venus/agent/venus/auth"
	"strings"
)

const cookieKey = "venus-authorization"
const headerKey = "Authorization"

// MustLogin parse header and set token into context
// [Authorization: Bearer]
func MustLogin(aor auth.Authenticator) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		_, exist := auth.FromContext(ctx)
		if exist {
			return
		}
		tokenString, _ := ctx.Cookie(cookieKey)
		if tokenString == "" {
			tokenString = strings.TrimPrefix(ctx.Request.Header.Get(headerKey), "Bearer ")
		}
		if len(tokenString) == 0 {
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
