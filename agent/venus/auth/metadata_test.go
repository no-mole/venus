package auth

import (
	"testing"

	"google.golang.org/grpc/metadata"
)

func TestTokenStringFromGrpcMetadata(t *testing.T) {
	ss := []struct {
		prefix string
		suffix string
	}{
		{"Bearer ", "   "},
		{"Bearer ", ""},
		{"bearer ", "   "},
		{"bearer ", ""},
		{"", "   "},
		{"  ", ""},
	}
	tokenStr := "xxxxx"
	for _, s := range ss {
		input := s.prefix + tokenStr + s.suffix
		md := metadata.MD{
			"authorization": []string{input},
		}
		output := TokenStringFromGrpcMetadata(md)
		if output != tokenStr {
			t.Errorf("input:%s,expected:%s,got:%s\n", input, tokenStr, output)
		}
	}

}
