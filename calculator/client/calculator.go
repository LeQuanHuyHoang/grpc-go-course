package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/calculator/proto"
	"io"
	"log"
	"time"
)

func doSum(c pb.CalculatorServiceClient) {
	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNumber:  5,
		SecondNumber: 7,
	})

	if err != nil {
		log.Fatalf("Failed to calculate: %v", err)
	}

	log.Printf("Sum: %d\n", res.Sum)
}

func doPrime(c pb.CalculatorServiceClient) {
	log.Println("DoPrime was invoked")

	req := &pb.PrimeRequest{
		Number: 120,
	}

	stream, err := c.Prime(context.Background(), req)

	if err != nil {
		log.Fatalf("Error while calling Prime: %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Error while reading stream: %v", err)
		}

		log.Printf("Prime received: %s", msg)
	}
}

func doAvg(c pb.CalculatorServiceClient) {
	log.Println("DoLongGreet was invoked")

	reqs := []*pb.AverageRequest{
		{Number: 1},
		{Number: 2},
		{Number: 3},
	}

	stream, err := c.Avg(context.Background())

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

	log.Printf("Avg: %f", res.Result)
}
