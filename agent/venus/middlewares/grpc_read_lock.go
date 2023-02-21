package middlewares

import (
	"context"
	"google.golang.org/grpc"
	"sync"
)

func ReadLock(lock *sync.RWMutex) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp interface{}, err error) {
		lock.RLock()
		defer lock.RUnlock()
		return handler(ctx, req)
	}
}
