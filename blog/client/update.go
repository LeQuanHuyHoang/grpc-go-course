package main

import (
	"context"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/blog/proto"
	"log"
)

func updateBlog(c pb.BlogServiceClient, id string, newAuthorId string, newTitle string) {
	log.Println("---Update blog---")

	blog := &pb.Blog{
		Id:       id,
		AuthorId: newAuthorId,
		Title:    newTitle,
	}

	_, err := c.UpdateBlog(context.Background(), blog)

	if err != nil {
		log.Fatalf("Unexpected error: %v\n", err)
	}

	log.Println("Blog has been updated")
}
