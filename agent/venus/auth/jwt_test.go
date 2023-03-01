package auth

import (
	"context"
	"testing"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func TestTokenSign(t *testing.T) {
	ctx := context.Background()
	token := NewJwtTokenWithClaim(time.Now().Add(time.Second), "aaa", "ccc", TokenTypeUser, map[string]Permission{"default": PermissionWriteRead})
	aor := genAuth()
	tokenString, err := aor.Sign(ctx, token)
	if err != nil {
		t.Fatal(err)
	}
	newToken, err := aor.Parse(ctx, tokenString)
	if err != nil {
		t.Fatal(err)
	}
	if claims, ok := newToken.Claims.(*Claims); !ok {
		t.Fatal("not *Claims")
	} else if claims.UniqueID != "aaa" || claims.Name != "ccc" || claims.NamespaceRoles["default"] != PermissionWriteRead {
		t.Fatal(claims)
	}

	<-time.After(1100 * time.Millisecond)
	_, err = aor.Parse(ctx, tokenString)
	if err == nil {
		t.Fatal("not validate expire time")
	}
	if jwtErr, ok := err.(*jwt.ValidationError); !ok {
		t.Fatal("not jwt.ValidationError")
	} else if !jwtErr.Is(jwt.ErrTokenExpired) {
		t.Fatal("!jwtErr.Is(jwt.ErrTokenExpired)")
	}

}
