package main

import (
	pb "github.com/LeQuanHuyHoang/grpc-go-course/greet/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"
)

var addr string = "0.0.0.0:8080"

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		log.Fatalf("Failed to connect: %v\n", err)
	}

	defer conn.Close()
	c := pb.NewGreetServiceClient(conn)

	//doGreet(c)
	//doGreetManyTimes(c)
	//doLongGreet(c)
	doGreetEveryone(c)
}
