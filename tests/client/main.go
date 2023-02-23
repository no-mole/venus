package main

import (
	"context"
	"fmt"
	"time"

	clientv1 "github.com/no-mole/venus/client/v1"
)

func main() {
	ctx := context.Background()

	cfg := clientv1.Config{
		//Endpoints:      []string{"127.0.0.1:6233", "127.0.0.1:6234", "127.0.0.1:6235"},
		//Endpoints:      []string{"127.0.0.1:6233"},
		Endpoints:      []string{"127.0.0.1:6234"},
		DialTimeout:    time.Second,
		MaxRetries:     5,
		PerCallTimeout: time.Second,
		Context:        ctx,
		PeerToken:      "zddno1",
	}
	client, err := clientv1.NewClient(cfg)
	if err != nil {
		panic(err)
	}
	namespace := "default"
	key := "key112"
	item, err := client.FetchKey(ctx, namespace, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("FetchKey %+v\n", item)

	item, err = client.AddKV(ctx, namespace, key, "json", time.Now().Format(time.RFC3339))
	if err != nil {
		panic(err)
	}
	fmt.Printf("AddKV %+v\n", item)
	<-time.After(time.Second)
	item, err = client.FetchKey(ctx, namespace, key)
	if err != nil {
		panic(err)
	}
	fmt.Printf("FetchKey %+v\n", item)
}
