package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/blog/proto"
	"google.golang.org/protobuf/types/known/emptypb"
	"io"
	"log"
)

func listBlog(c pb.BlogServiceClient) {
	log.Println("listBlog")

	stream, err := c.ListBlog(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error while calling ListBlog: %v\n", err)
	}

	for {
		res, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Failed to connect to Blog: %v\n", err)
		}

		log.Println(res)
	}
}
