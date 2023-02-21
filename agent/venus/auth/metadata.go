package auth

import (
	"google.golang.org/grpc/metadata"
)

const (
	GrpcMetadataKey = "authorization"
)

func FromGrpcMetadata(md metadata.MD) string {
	headers := md.Get(GrpcMetadataKey)
	if len(headers) == 0 {
		return ""
	}
	return headers[0]
}

func WithGrpcMetadata(md metadata.MD, token string) {
	md.Set(GrpcMetadataKey, token)
}
