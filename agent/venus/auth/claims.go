package auth

import (
	"github.com/golang-jwt/jwt/v4"
	"github.com/no-mole/venus/agent/errors"
	"strconv"
	"time"
)

type TokenType string
type Permission string

const (
	TokenTypeAccessKey     TokenType = "ak"
	TokenTypeUser          TokenType = "us"
	TokenTypeAdministrator TokenType = "ad"

	PermissionWriteRead Permission = "wr"
	PermissionReadOnly  Permission = "r"
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

func IsClaims(token *jwt.Token) (*Claims, bool) {
	if c, ok := token.Claims.(*Claims); ok {
		return c, true
	}
	return nil, false
}

func NewJwtTokenWithClaim(expiresAt time.Time, uniqueID, name string, tt TokenType, namespaceRoles map[string]Permission) *jwt.Token {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &Claims{
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        strconv.Itoa(time.Now().Nanosecond()),
			ExpiresAt: jwt.NewNumericDate(expiresAt),
		},
		UniqueID:       uniqueID,
		Name:           name,
		TokenType:      tt,
		NamespaceRoles: namespaceRoles,
	})
	return token
}

type Claims struct {
	jwt.RegisteredClaims `json:"rc,omitempty"`
	UniqueID             string                `json:"uid,omitempty"`
	Name                 string                `json:"nm,omitempty"`
	TokenType            TokenType             `json:"tt,omitempty"`
	NamespaceRoles       map[string]Permission `json:"nr,omitempty"`
}

func (i *Claims) Valid() error {
	err := i.RegisteredClaims.Valid()
	if err != nil {
		return err
	}
	if !IsExpectedTokenType(i.TokenType) {
		return errors.ErrorTokenNotValid
	}
	return nil
}

func (i *Claims) Type() TokenType {
	return i.TokenType
}

func (i *Claims) Writable(namespace string) bool {
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

func (i *Claims) Readable(namespace string) bool {
	if i.Type() == TokenTypeAdministrator {
		return true
	}
	if _, ok := i.NamespaceRoles[namespace]; ok {
		return true
	}
	return false
}
