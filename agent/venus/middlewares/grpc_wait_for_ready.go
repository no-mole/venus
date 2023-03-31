package middlewares

import "google.golang.org/grpc"

func ClientWaitForReady() grpc.DialOption {
	return grpc.WithDefaultCallOptions(
		grpc.WaitForReady(true),
	)
}
