package middlewares

import (
	"context"
	grpcMiddleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"github.com/no-mole/venus/agent/errors"
	"github.com/no-mole/venus/agent/venus/auth"
	"google.golang.org/grpc"
	"google.golang.org/grpc/metadata"
)

func serverMustLogin(ctx context.Context, aor auth.Authenticator) (context.Context, error) {
	meta, has := metadata.FromIncomingContext(ctx)
	if !has {
		return nil, errors.ErrorGrpcNotLogin
	}
	tokenStr := auth.FromGrpcMetadata(meta)
	token, err := aor.Parse(ctx, tokenStr)
	if err != nil {
		return nil, err
	}
	return auth.WithContext(ctx, token), nil
}

func MustLoginUnaryServerInterceptor(aor auth.Authenticator) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		ctx, err = serverMustLogin(ctx, aor)
		if err != nil {
			return nil, err
		}
		return handler(ctx, req)
	}
}
func MustLoginStreamServerInterceptor(aor auth.Authenticator) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		ctx, err := serverMustLogin(ss.Context(), aor)
		if err != nil {
			return err
		}
		return handler(srv, &grpcMiddleware.WrappedServerStream{
			ServerStream:   ss,
			WrappedContext: ctx,
		})
	}
}

func clientMustLogin(ctx context.Context, aor auth.Authenticator) (context.Context, error) {
	token, has := auth.FromContext(ctx)
	if !has {
		return nil, errors.ErrorNotLogin
	}
	tokenStr, err := aor.Sign(ctx, token)
	if err != nil {
		return nil, err
	}
	meta, has := metadata.FromOutgoingContext(ctx)
	if !has {
		meta = metadata.MD{}
	}
	auth.WithGrpcMetadata(meta, tokenStr)
	ctx = metadata.NewOutgoingContext(ctx, meta)
	return ctx, nil
}

func MustLoginUnaryClientInterceptor(aor auth.Authenticator) grpc.UnaryClientInterceptor {
	return func(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn, invoker grpc.UnaryInvoker, opts ...grpc.CallOption) (err error) {
		ctx, err = clientMustLogin(ctx, aor)
		if err != nil {
			return err
		}
		return invoker(ctx, method, req, reply, cc, opts...)
	}
}

func MustLoginStreamClientInterceptor(aor auth.Authenticator) grpc.StreamClientInterceptor {
	return func(ctx context.Context, desc *grpc.StreamDesc, cc *grpc.ClientConn, method string, streamer grpc.Streamer, opts ...grpc.CallOption) (_ grpc.ClientStream, err error) {
		ctx, err = clientMustLogin(ctx, aor)
		if err != nil {
			return nil, err
		}
		return streamer(ctx, desc, cc, method, opts...)
	}
}
