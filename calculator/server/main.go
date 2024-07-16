package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/LeQuanHuyHoang/grpc-go-course/calculator/proto"
)

var addr string = "0.0.0.0:8081"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen to: %v\n", err)
	}

	log.Printf("Listening on port: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
