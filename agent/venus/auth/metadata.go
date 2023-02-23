package auth

import (
	"google.golang.org/grpc/metadata"
	"strings"
)

const (
	GrpcTokenMetadataKey = "authorization"
)

func TokenStringFromGrpcMetadata(md metadata.MD) string {
	headers := md.Get(GrpcTokenMetadataKey)
	if len(headers) == 0 {
		return ""
	}
	tk := strings.TrimPrefix(headers[0], "Bearer ")
	tk = strings.TrimPrefix(tk, "bearer ")
	tk = strings.Trim(tk, " ")
	return tk
}

func WithGrpcMetadata(md metadata.MD, token string) {
	md.Set(GrpcTokenMetadataKey, token)
}
