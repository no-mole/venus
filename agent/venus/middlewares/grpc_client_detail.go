package middlewares

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/no-mole/venus/agent/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func UnaryClientWithCallerDetail() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.MD{}
		}
		_, _, ok = utils.FromMetadata(md)
		if !ok {
			//如果ctx的md中不存在client detail，则写入md
			//如果存在，则默认转发
			utils.WithMetadata(md)
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func StreamClientWithCallerDetail() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (grpc.ClientStream, error) {
		md, ok := metadata.FromOutgoingContext(ctx)
		if !ok {
			md = metadata.MD{}
		}
		_, _, ok = utils.FromMetadata(md)
		if !ok {
			//如果ctx的md中不存在client detail，则写入md
			//如果存在，则默认转发
			utils.WithMetadata(md)
			ctx = metadata.NewOutgoingContext(ctx, md)
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}

func UnaryServerWithCallerDetail() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = utils.WithContext(ctx, md)
		}
		return handler(ctx, req)
	}
}

func StreamServerWithCallerDetail() grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx := ss.Context()
		md, ok := metadata.FromIncomingContext(ctx)
		if ok {
			ctx = utils.WithContext(ctx, md)
		}
		return handler(srv, &grpcMiddleware.WrappedServerStream{
			ServerStream:   ss,
			WrappedContext: ctx,
		})
	}
}
