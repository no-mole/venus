package auth

import (
	"context"

	"github.com/no-mole/venus/agent/errors"

	"github.com/golang-jwt/jwt/v4"
)

type Authenticator interface {
	TokenProvider
	Writable(token *jwt.Token, namespace string) bool
	WritableContext(ctx context.Context, namespace string) (bool, error)

	Readable(token *jwt.Token, namespace string) bool
	ReadableContext(ctx context.Context, namespace string) (bool, error)

	IsAdministrator(token *jwt.Token) bool
	IsAdministratorContext(ctx context.Context) (bool, error)
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

func (a *authenticator) IsAdministratorContext(ctx context.Context) (bool, error) {
	token, has := FromContext(ctx)
	if !has {
		return false, errors.ErrorNotLogin
	}
	return a.IsAdministrator(token), nil
}

func (a *authenticator) IsAdministrator(token *jwt.Token) bool {
	claims, ok := IsClaims(token)
	if !ok {
		return false
	}
	return claims.TokenType == TokenTypeAdministrator
}
