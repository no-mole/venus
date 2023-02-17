package main

import (
	"context"
	"github.com/no-mole/venus/internal/proto/pbcluster"
	"github.com/no-mole/venus/proto/pbkv"
	"google.golang.org/grpc"
	"google.golang.org/grpc/connectivity"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
	"time"
)

func main() {
	ctx := context.Background()
	endpoint := "127.0.0.1:3333"

	clientConn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial(%s):%v\n", endpoint, err)
	}
	for clientConn.GetState() != connectivity.Ready {
		clientConn.Connect()
	}
	println("start")
	//kvClient1 := pbkv.NewKVClient(clientConn)
	//for i := 0; i < 1000; i++ {
	//	start := time.Now()
	//	_, err = kvClient1.AddKV(ctx, &pbkv.KVItem{
	//		Namespace: "default",
	//		Key:       "key1",
	//		DataType:  "json",
	//		Value:     time.Now().String(),
	//		Version:   "v11",
	//	})
	//	if err != nil {
	//		log.Fatalf("kvClient.AddKV(%s):%v\n", endpoint, err)
	//	}
	//	println(time.Now().String(), time.Now().Sub(start).String())
	//}
	//os.Exit(0)

	client := pbcluster.NewClusterClient(clientConn)
	status, err := client.Stats(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("grpc.Dial(%s):%v\n", endpoint, err)
	}
	log.Printf("%+v\n", status)

	kvClient := pbkv.NewKVClient(clientConn)
	//item, err := kvClient.AddKV(ctx, &pbkv.KVItem{
	//	Namespace: "default",
	//	Key:       "key1",
	//	DataType:  "json",
	//	Value:     time.Now().String(),
	//})
	//if err != nil {
	//	log.Fatalf("kvClient.AddKV(%s):%v\n", endpoint, err)
	//}
	//log.Printf("%+v", item)
	item, err := kvClient.FetchKey(ctx, &pbkv.FetchKeyRequest{
		Namespace: "default",
		Key:       "key1",
	})
	if err != nil {
		log.Fatalf("kvClient.FetchKey(%s):%v\n", endpoint, err)
	}
	log.Printf("%+v\n", item)

	item, err = kvClient.AddKV(ctx, &pbkv.KVItem{
		Namespace: "default",
		Key:       "key1",
		DataType:  "json",
		Value:     time.Now().String(),
		Version:   "v11",
	})
	if err != nil {
		log.Fatalf("kvClient.AddKV(%s):%v\n", endpoint, err)
	}
	log.Printf("%+v\n", item)

	//namespaceClient := pbnamespace.NewNamespaceClient(clientConn)
	//item, err := namespaceClient.AddNamespace(ctx, &pbnamespace.NamespaceItem{
	//	NamespaceCn: "测试name888",
	//	NamespaceEn: "test_name_888",
	//	Creator:     "zdd",
	//	CreateTime:  time.Now().Format(time.RFC3339Nano),
	//})
	//if err != nil {
	//	log.Fatalf("namespaceClient.AddNamespace(%+v):%v\n", item, err)
	//}
	//log.Printf("add item:%+v", item)
	//listResp, err := namespaceClient.ListNamespaces(ctx, &emptypb.Empty{})
	//if err != nil {
	//	log.Fatalf("namespaceClient.ListNamespaces(%+v):%v\n", listResp, err)
	//}
	//log.Printf("cur list:%+v\n", listResp)

	//item, err = namespaceClient.AddNamespace(ctx, &pbnamespace.NamespaceItem{
	//	NamespaceCn: "测试name2",
	//	NamespaceEn: "test_name_2",
	//	Creator:     "zdd",
	//	CreateTime:  time.Now().Format(time.RFC3339Nano),
	//})
	//if err != nil {
	//	log.Fatalf("namespaceClient.AddNamespace(%+v):%v\n", item, err)
	//}
	//log.Printf("add item:%+v", item)
	//listResp, err = namespaceClient.ListNamespaces(ctx, &emptypb.Empty{})
	//if err != nil {
	//	log.Fatalf("namespaceClient.ListNamespaces(%+v):%v\n", listResp, err)
	//}
	//log.Printf("cur list:%+v", listResp)
	//
	//item, err = namespaceClient.AddNamespace(ctx, &pbnamespace.NamespaceItem{
	//	NamespaceCn: "测试name5",
	//	NamespaceEn: "test_name_5",
	//	Creator:     "sdgf",
	//	CreateTime:  time.Now().Format(time.RFC3339Nano),
	//})
	//if err != nil {
	//	log.Fatalf("namespaceClient.AddNamespace(%+v):%v\n", item, err)
	//}
	//log.Printf("add item:%+v", item)
	//listResp, err = namespaceClient.ListNamespaces(ctx, &emptypb.Empty{})
	//if err != nil {
	//	log.Fatalf("namespaceClient.ListNamespaces(%+v):%v\n", listResp, err)
	//}
	//log.Printf("cur list:%+v", listResp)
}
