package middlewares

import (
	grpcRetry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
	"google.golang.org/grpc"
	"time"
)

func ClientRetry(maxTimes uint, backoffExponential time.Duration) grpc.DialOption {
	// Give up after 5 retries.
	retryOpts := []grpcRetry.CallOption{
		grpcRetry.WithBackoff(grpcRetry.BackoffExponential(backoffExponential)),
		grpcRetry.WithMax(maxTimes),
	}
	return grpc.WithUnaryInterceptor(grpcRetry.UnaryClientInterceptor(retryOpts...))
}
