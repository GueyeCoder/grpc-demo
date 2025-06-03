package main

import (
	"io"
	"log"

	pb "github.com/GueyeCoder/grpc-demo/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {
	var messages []string

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.MessageList{Message: messages})
		}
		if err != nil {
			return err
		}
		log.Printf("Got request with names: %v", req.Name)
		messages = append(messages, "hello ", req.Name)

	}
}
