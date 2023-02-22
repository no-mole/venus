package secret

import (
	"crypto/sha256"
	"encoding/base64"
)

func Confusion(id, secret string) string {
	data := sha256.Sum256([]byte(id + secret))
	return base64.RawURLEncoding.EncodeToString(data[:])
}
