package middlewares

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RecoverHandleFunc func(ctx context.Context, fullMethodName string, p interface{}) (err error)

func ZapLoggerRecoverHandle(logger *zap.Logger) RecoverHandleFunc {
	return func(ctx context.Context, fullMethodName string, p interface{}) (err error) {
		logger.Error("grpc server panic")
		return status.Errorf(codes.Unknown, "panic triggered: %v", p)
	}
}

func UnaryServerRecover(fn RecoverHandleFunc) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		panicked := true
		defer func() {
			if r := recover(); r != nil || panicked {
				err = fn(ctx, info.FullMethod, r)
			}
		}()
		resp, err = handler(ctx, req)
		panicked = false
		return resp, err
	}
}

func StreamServerRecover(fn RecoverHandleFunc) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) (err error) {
		panicked := true
		defer func() {
			if r := recover(); r != nil || panicked {
				err = fn(ss.Context(), info.FullMethod, r)
			}
		}()
		err = handler(srv, ss)
		panicked = false
		return err
	}
}
