package middlewares

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/venus/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
	"strings"
)

func serverMustLogin(ctx context.Context, fullMethodName string, aor auth.Authenticator) (context.Context, error) {
	//all Login method do not need login
	if strings.Contains(fullMethodName, "Login") {
		return ctx, nil
	}
	if strings.Contains(fullMethodName, "RaftTransport") {
		return ctx, nil
	}
	meta, has := metadata.FromIncomingContext(ctx)
	if !has {
		return nil, errors.ErrorGrpcNotLogin
	}
	tokenStr := auth.TokenStringFromGrpcMetadata(meta)
	token, err := aor.Parse(ctx, tokenStr)
	if err != nil {
		return nil, err
	}
	return auth.WithContext(ctx, token), nil
}

func MustLoginUnaryServerInterceptor(aor auth.Authenticator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = serverMustLogin(ctx, info.FullMethod, aor)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
func MustLoginStreamServerInterceptor(aor auth.Authenticator) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := serverMustLogin(ss.Context(), info.FullMethod, aor)
		if err != nil {
			return err
		}
		return handler(srv, &grpcMiddleware.WrappedServerStream{
			ServerStream:   ss,
			WrappedContext: ctx,
		})
	}
}

// clientMustLogin 如果ctx中存在token，则传递raw token，与authTokenBundle同时生效
// (1)grpc raft transport只使用了authTokenBundle
// (2)普通client 只使用了authTokenBundle
// (3)proxy server 使用了 clientMustLogin,因为server的拦截器会把token解析并种到ctx
// (4)join server使用了base token，也是 clientMustLogin 起作用
func clientMustLogin(ctx context.Context) (context.Context, error) {
	token, has := auth.FromContext(ctx)
	if !has {
		return ctx, nil
	}

	meta, has := metadata.FromOutgoingContext(ctx)
	if !has {
		meta = metadata.MD{}
	}
	auth.WithGrpcMetadata(meta, token.Raw)
	ctx = metadata.NewOutgoingContext(ctx, meta)
	return ctx, nil
}

func MustLoginUnaryClientInterceptor() grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		ctx, err = clientMustLogin(ctx)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func MustLoginStreamClientInterceptor() grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (_ grpc.ClientStream, err error) {
		ctx, err = clientMustLogin(ctx)
		if err != nil {
			return nil, err
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
