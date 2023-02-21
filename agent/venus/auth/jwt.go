package auth

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"github.com/no-mole/venus/agent/errors"
)

type TokenProvider interface {
	Sign(ctx context.Context, token *jwt.Token) (string, error)
	Parse(ctx context.Context, tokenString string) (*jwt.Token, error)
}

func NewTokenProvider(sampleSecret []byte) TokenProvider {
	return &tokenJwt{
		sampleSecret: sampleSecret,
	}
}

type tokenJwt struct {
	sampleSecret []byte
}

func (t *tokenJwt) Sign(ctx context.Context, token *jwt.Token) (string, error) {
	return token.SignedString(t.sampleSecret)
}

func (t *tokenJwt) Parse(ctx context.Context, tokenString string) (*jwt.Token, error) {
	if tokenString == "" {
		return nil, errors.ErrorTokenNotValid
	}
	token, err := jwt.ParseWithClaims(tokenString, Claims{RegisteredClaims: &jwt.RegisteredClaims{}}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.ErrorTokenUnexpectedSigningMethod
		}
		return t.sampleSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.ErrorTokenNotValid
	}

	if _, ok := token.Claims.(Claims); ok {
		return nil, errors.ErrorTokenNotValid
	}
	return token, err
}
