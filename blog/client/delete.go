package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/blog/proto"
	"log"
)

func deleteBlog(c pb.BlogServiceClient, in string) {
	log.Println("deleteBlog")

	_, err := c.DeleteBlog(context.Background(), &pb.BlogId{Id: in})

	if err != nil {
		log.Fatalf("error deleting blog: %v", err)
	}

	log.Println("deleteBlog success")
}
