package main

import (
	"context"
	"log"
	"time"

	pd "github.com/GueyeCoder/grpc-demo/proto"
)

func callSayHello(client pd.GreetServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	res, err := client.SayHello(ctx, &pd.NoParam{})
	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("%s", res.Message)
}
