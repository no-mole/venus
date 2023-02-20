package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/no-mole/venus/agent/errors"
	"time"
)

type TokenType string
type Permission string

const (
	TokenTypeAccessKey     = "ak"
	TokenTypeUser          = "us"
	TokenTypeAdministrator = "ad"

	PermissionWriteRead = "wr"
	PermissionReadOnly  = "r"
)

var tokenTypes = map[TokenType]struct{}{
	TokenTypeAccessKey:     {},
	TokenTypeUser:          {},
	TokenTypeAdministrator: {},
}

func IsExpectedTokenType(tt TokenType) bool {
	if _, ok := tokenTypes[tt]; ok {
		return true
	}
	return false
}

func IsClaims(token *jwt.Token) (Claims, bool) {
	if c, ok := token.Claims.(Claims); ok {
		return c, true
	}
	return Claims{}, false
}

func NewJwtTokenWithClaim(expiresAt time.Time, tt TokenType, namespaceRoles map[string]Permission) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, Claims{
		RegisteredClaims: &jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		TokenType:      tt,
		NamespaceRoles: namespaceRoles,
	})
	return token
}

type Claims struct {
	*jwt.RegisteredClaims
	TokenType      TokenType             `json:"tt"`
	NamespaceRoles map[string]Permission `json:"nr"`
}

func (i Claims) Valid() error {
	err := i.RegisteredClaims.Valid()
	if err != nil {
		return err
	}
	if !IsExpectedTokenType(i.TokenType) {
		return errors.ErrorTokenNotValid
	}
	return nil
}

func (i Claims) Type() TokenType {
	return i.TokenType
}

func (i Claims) Writable(namespace string) bool {
	if i.Type() == TokenTypeAdministrator {
		return true
	}
	if role, ok := i.NamespaceRoles[namespace]; ok {
		if role == PermissionWriteRead {
			return true
		}
	}
	return false
}

func (i Claims) Readable(namespace string) bool {
	if i.Type() == TokenTypeAdministrator {
		return true
	}
	if _, ok := i.NamespaceRoles[namespace]; ok {
		return true
	}
	return false
}
