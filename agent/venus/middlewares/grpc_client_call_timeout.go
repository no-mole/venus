package middlewares

import (
	"context"
	"google.golang.org/grpc"
	"time"
)

// DefaultTimeoutUnaryClientInterceptor returns a new unary client interceptor that sets a timeout on the request context(when ctx.Deadline() is false).
func DefaultTimeoutUnaryClientInterceptor(timeout time.Duration) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		if _, ok := ctx.Deadline(); !ok {
			var cancel context.CancelFunc
			ctx, cancel = context.WithTimeout(ctx, timeout)
			defer cancel()
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}
