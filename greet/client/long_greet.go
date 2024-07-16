package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/greet/proto"
	"log"
	"time"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Println("DoLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Huy"},
		{FirstName: "Hoang"},
		{FirstName: "Thu"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error while calling LongGreet: %v", err)
	}

	for _, req := range reqs {
		log.Printf("Sending req: %v\n", req)
		stream.Send(req)
		time.Sleep(time.Second)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error while recieving stream: %v", err)
	}

	log.Printf("LongGreet: %s", res.Result)
}
