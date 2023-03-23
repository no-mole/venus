package main

import (
	"context"
	"fmt"
	clientv1 "github.com/no-mole/venus/client/v1"
	"time"
)

func main() {
	ctx := context.Background()
	//ctx, _ = context.WithTimeout(ctx, 20*time.Second)
	conf := clientv1.Config{
		Endpoints:   []string{"127.0.0.1:6233"},
		DialTimeout: time.Second,
		PeerToken:   "biomind",
		Context:     ctx,
	}
	cli, err := clientv1.NewClient(conf)
	if err != nil {
		panic(err)
	}
	item, err := cli.Grant(ctx, 5)
	if err != nil {
		panic(err)
	}
	fmt.Printf("grant %+v\n", item)
	for {
		<-time.After(time.Second)
		err = cli.KeepaliveOnce(ctx, item.LeaseId)
		if err != nil {
			panic(err)
		}
		println("KeepaliveOnce")
	}
}
