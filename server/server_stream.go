package main

import (
	"log"
	"time"

	pb "github.com/GueyeCoder/grpc-demo/proto"
)

func (h *helloServer) callSayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Got request with names: %v", req.Names)

	for _, name := range req.Names {
		res := &pb.HelloRequest{
			Message: "hello " + name,
		}
		if err := stream.Send((*pb.HelloResponse)(res)); err != nil {
			return err
		}
		time.Sleep(2 * time.Second)
	}
	return nil
}
