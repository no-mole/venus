package auth

import (
	"context"

	"github.com/golang-jwt/jwt/v4"
)

const (
	TokenContextKey = "jwtToken"
)

func FromContext(ctx context.Context) (*jwt.Token, bool) {
	val := ctx.Value(TokenContextKey)
	if val == nil {
		return nil, false
	}
	if v, ok := val.(*jwt.Token); ok {
		return v, true
	}
	return nil, false
}

func WithContext(ctx context.Context, token *jwt.Token) context.Context {
	return context.WithValue(ctx, TokenContextKey, token)
}
