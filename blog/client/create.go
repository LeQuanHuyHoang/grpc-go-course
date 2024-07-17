package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/blog/proto"
	"log"
)

func createBlog(c pb.BlogServiceClient) string {
	log.Println("---Creating blog---")

	blog := &pb.Blog{
		AuthorId: "HuyHoang",
		Title:    "First Post",
		Content:  "Content of the first post",
	}

	res, err := c.CreateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Printf("Blog has been created: %s\n", res.Id)
	return res.Id
}
