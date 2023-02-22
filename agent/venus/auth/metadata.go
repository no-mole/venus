package auth

import (
	"google.golang.org/grpc/metadata"
)

const (
	GrpcTokenMetadataKey     = "authorization"
	GrpcPeerTokenMetadataKey = "peer-token"
)

func TokenStringFromGrpcMetadata(md metadata.MD) string {
	headers := md.Get(GrpcTokenMetadataKey)
	if len(headers) == 0 {
		return ""
	}
	return headers[0]
}

func WithGrpcMetadata(md metadata.MD, token string) {
	md.Set(GrpcTokenMetadataKey, token)
}
