package main

import (
	"context"
	"github.com/Jille/raftadmin/proto"
	"github.com/no-mole/venus/proto/pbnamespace"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
	"log"
)

func main() {
	ctx := context.Background()
	endpoint := "127.0.0.1:3333"

	clientConn, err := grpc.Dial(endpoint, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("grpc.Dial(%s):%v\n", endpoint, err)
	}
	client := proto.NewRaftAdminClient(clientConn)
	status, err := client.Stats(ctx, &proto.StatsRequest{})
	if err != nil {
		log.Fatalf("grpc.Dial(%s):%v\n", endpoint, err)
	}
	log.Printf("%+v", status)

	namespaceClient := pbnamespace.NewNamespaceClient(clientConn)
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
	listResp, err := namespaceClient.ListNamespaces(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatalf("namespaceClient.ListNamespaces(%+v):%v\n", listResp, err)
	}
	log.Printf("cur list:%+v\n", listResp)

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
