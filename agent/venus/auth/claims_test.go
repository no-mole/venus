package auth

import (
	"testing"
	"time"
)

var sampleSecret = []byte("venus")

func genAuth() Authenticator {
	tp := NewTokenProvider(sampleSecret)
	return NewAuthenticator(tp)
}

func TestAdminTokenReadWrite(t *testing.T) {
	adminToken := NewJwtTokenWithClaim(time.Now().Add(10*time.Second), "venus", "venus", TokenTypeAdministrator, map[string]Permission{
		"default": PermissionWriteRead,
		"named":   PermissionReadOnly,
	})
	aor := genAuth()
	if !aor.Readable(adminToken, "default") {
		t.Fatal()
	}
	if !aor.Writable(adminToken, "default") {
		t.Fatal()
	}
	if !aor.Readable(adminToken, "named") {
		t.Fatal()
	}
	if !aor.Writable(adminToken, "named") {
		t.Fatal()
	}
	if !aor.Readable(adminToken, "aaa") {
		t.Fatal()
	}
	if !aor.Writable(adminToken, "aaa") {
		t.Fatal()
	}
}

func TestUserTokenReadWrite(t *testing.T) {
	adminToken := NewJwtTokenWithClaim(time.Now().Add(10*time.Second), "venus", "venus", TokenTypeUser, map[string]Permission{
		"default": PermissionWriteRead,
		"named":   PermissionReadOnly,
	})
	aor := genAuth()
	if !aor.Readable(adminToken, "default") {
		t.Fatal()
	}
	if !aor.Writable(adminToken, "default") {
		t.Fatal()
	}
	if !aor.Readable(adminToken, "named") {
		t.Fatal()
	}
	if aor.Writable(adminToken, "named") {
		t.Fatal()
	}
	if aor.Readable(adminToken, "aaa") {
		t.Fatal()
	}
	if aor.Writable(adminToken, "aaa") {
		t.Fatal()
	}
}
