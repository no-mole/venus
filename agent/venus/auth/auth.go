package auth

import (
	"context"

	"github.com/no-mole/venus/agent/errors"

	"github.com/golang-jwt/jwt/v4"
)

type Authenticator interface {
	TokenProvider
	Writable(token *jwt.Token, namespace string) bool
	Readable(token *jwt.Token, namespace string) bool
	WritableContext(ctx context.Context, namespace string) (bool, error)
	ReadableContext(ctx context.Context, namespace string) (bool, error)
}

type authenticator struct {
	TokenProvider
}

func NewAuthenticator(provider TokenProvider) Authenticator {
	return &authenticator{
		TokenProvider: provider,
	}
}

func (a *authenticator) Writable(token *jwt.Token, namespace string) bool {
	if claims, ok := IsClaims(token); ok {
		return claims.Writable(namespace)
	}
	return false
}

func (a *authenticator) Readable(token *jwt.Token, namespace string) bool {
	if claims, ok := IsClaims(token); ok {
		return claims.Readable(namespace)
	}
	return false
}

func (a *authenticator) WritableContext(ctx context.Context, namespace string) (bool, error) {
	token, has := FromContext(ctx)
	if !has {
		return false, errors.ErrorNotLogin
	}
	writable := a.Writable(token, namespace)
	return writable, nil
}

func (a *authenticator) ReadableContext(ctx context.Context, namespace string) (bool, error) {
	token, has := FromContext(ctx)
	if !has {
		return false, errors.ErrorNotLogin
	}
	readable := a.Readable(token, namespace)
	return readable, nil
}
