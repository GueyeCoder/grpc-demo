package main

import (
	"context"

	pd "github.com/GueyeCoder/grpc-demo/proto"
)

func (h *helloServer) SayHello(ctx context.Context, req *pd.NoParam) (*pd.HelloResponse, error) {
	return &pd.HelloResponse{
		Message: "Hello",
	}, nil
}
