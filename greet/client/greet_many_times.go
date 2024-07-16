package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/greet/proto"
	"io"
	"log"
)

func doGreetManyTimes(c pb.GreetServiceClient) {
	log.Println("DoGreetManyTimes was invoked")

	req := &pb.GreetRequest{
		FirstName: "Hoang",
	}

	stream, err := c.GreetManyTimes(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling GreetManyTimes: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("GreetManyTimes received: %s", msg)
	}
}
