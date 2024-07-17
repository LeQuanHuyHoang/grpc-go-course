package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/blog/proto"
	"log"
)

func readBlog(c pb.BlogServiceClient, id string) *pb.Blog {
	log.Println("---Read blog---")

	blogId := &pb.BlogId{
		Id: id,
	}

	res, err := c.ReadBlog(context.Background(), blogId)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog was read: %v\n", res)
	return res
}
