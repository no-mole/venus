package middlewares

import (
	"context"
	"time"

	"google.golang.org/grpc/peer"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

type LogOption func(opt *ServerOptions) zap.Field

type ServerOptions struct {
	ServerName string
	TimeStart  time.Time
}

func apply(logger *zap.Logger, opts ...LogOption) *zap.Logger {
	if len(opts) == 0 {
		return logger
	}
	serverOpts := &ServerOptions{}
	var fields []zap.Field
	for _, opt := range opts {
		fields = append(fields, opt(serverOpts))
	}
	return logger.With(fields...)
}

func UnaryServerAccessLog(logger *zap.Logger, opts ...LogOption) grpc.UnaryServerInterceptor {
	logger = apply(logger, opts...)
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		start := time.Now()
		ip := ""
		p, ok := peer.FromContext(ctx)
		if ok {
			ip = p.Addr.String()
		}
		defer func() {
			logger.Debug("grpc service caller",
				zap.String("remoteAddr", ip),
				zap.String("serviceName", info.FullMethod),
				zap.String("duration", time.Since(start).String()),
			)
		}()
		return handler(ctx, req)
	}
}

func StreamServerAccessLog(logger *zap.Logger, opts ...LogOption) grpc.StreamServerInterceptor {
	logger = apply(logger, opts...)
	return func(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo, handler grpc.StreamHandler) error {
		start := time.Now()
		ip := ""
		p, ok := peer.FromContext(ss.Context())
		if ok {
			ip = p.Addr.String()
		}
		defer func() {
			logger.Info("grpc stream service caller",
				zap.String("remoteAddr", ip),
				zap.String("serviceName", info.FullMethod),
				zap.String("duration", time.Since(start).String()),
			)
		}()
		return handler(srv, ss)
	}
}
