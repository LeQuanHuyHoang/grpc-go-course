package main

import (
	"google.golang.org/grpc"
	"log"
	"net"

	pb "github.com/LeQuanHuyHoang/grpc-go-course/greet/proto"
)

var addr string = "0.0.0.0:8080"

type Server struct {
	pb.GreetServiceServer
}

func main() {
	lis, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("Failed to listen on: %v\n", err)
	}

	log.Printf("Listening on: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v\n", err)
	}
}
