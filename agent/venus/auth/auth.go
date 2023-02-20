package auth

import (
	"github.com/golang-jwt/jwt/v4"
)

type Authenticator interface {
	TokenProvider
	Writable(token *jwt.Token, namespace string) bool
	Readable(token *jwt.Token, namespace string) bool
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
