package server

import (
	"context"
	"testing"
	"time"

	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
)

var clientConn *grpc.ClientConn
var client pbkv.KVServiceClient

func init() {
	start := time.Now()
	endpoint := "127.0.0.1:3333"
	var err error
	clientConn, err = grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	for clientConn.GetState() != connectivity.Ready {
		clientConn.Connect()
	}
	client = pbkv.NewKVServiceClient(clientConn)
	println("conn", time.Since(start).String())
}
func BenchmarkLeaderWrite(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.AddKV(ctx, &pbkv.KVItem{
				Namespace: "default",
				Key:       "key1",
				DataType:  "json",
				Value:     "111",
				Version:   "v11",
			})
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}

func BenchmarkLeaderRead(b *testing.B) {
	ctx := context.Background()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			_, err := client.FetchKey(ctx, &pbkv.FetchKeyRequest{
				Namespace: "default",
				Key:       "key1",
			})
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
