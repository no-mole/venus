package clientv1

import (
	"google.golang.org/grpc"
)

type Client struct {
	callOpts []grpc.CallOption
	conn     grpc.ClientConnInterface
}
