package main

import (
	pb "github.com/LeQuanHuyHoang/grpc-go-course/blog/proto"
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
	c := pb.NewBlogServiceClient(conn)

	id := createBlog(c)
	readBlog(c, id)
	//readBlog(c, "invalidID")
	//updateBlog(c, id, "Thu", "2015")
	listBlog(c)
	deleteBlog(c, id)
	listBlog(c)
}
