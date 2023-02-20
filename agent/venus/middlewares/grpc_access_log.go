package middlewares

import (
	"context"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"time"
)

type LogOption func(opt *ServerOptions) zap.Field

type ServerOptions struct {
	ServerName string
	TimeStart  time.Time
}

func UnaryServerAccessLog(logger *zap.Logger, opts ...LogOption) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		defer func() {
			logger.Debug("grpc service caller", zap.String("serviceName", info.FullMethod), zap.String("duration", time.Now().Sub(start).String()))
		}()
		return handler(ctx, req)
	}
}

func StreamServerAccessLog(logger *zap.Logger, opts ...LogOption) grpc.StreamServerInterceptor {
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		defer func() {
			logger.Info("grpc stream service caller", zap.String("serviceName", info.FullMethod), zap.String("duration", time.Now().Sub(start).String()))
		}()
		return handler(srv, ss)
	}
}
