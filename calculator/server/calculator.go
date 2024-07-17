package main

import (
	"context"
	"fmt"
	pb "github.com/LeQuanHuyHoang/grpc-go-course/calculator/proto"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"io"
	"log"
	"math"
	"slices"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	rs := in.FirstNumber + in.SecondNumber
	return &pb.SumResponse{
		Sum: rs,
	}, nil
}

func (s *Server) Prime(in *pb.PrimeRequest, stream pb.CalculatorService_PrimeServer) error {
	k := 2
	N := in.Number

	for {
		if N <= 1 {
			break
		}
		if N%int32(k) == 0 {
			stream.Send(&pb.PrimeResponse{
				Result: int32(k),
			})
			N = N / int32(k)
		} else {
			k++
		}
	}

	return nil
}

func (s *Server) Avg(stream pb.CalculatorService_AvgServer) error {
	log.Println("Avg function is invoked")

	var rs []int32
	var res float64

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			for _, v := range rs {
				res += float64(v)
			}
			res = res / float64(len(rs))
			return stream.SendAndClose(&pb.AverageResponse{
				Result: res,
			})
		}

		if err != nil {
			log.Fatalf("Error while reading client stream: %v", err)
		}

		rs = append(rs, req.Number)
	}
}

func (s *Server) Maximum(stream pb.CalculatorService_MaximumServer) error {
	log.Println("Maximum function")
	var reqs []int32

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading from client: %v", err)
		}

		reqs = append(reqs, req.Number)
		res := slices.Max(reqs)
		err = stream.Send(&pb.MaximumResponse{
			Max: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client: %v", err)
		}
	}
}

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {
	log.Println("Sqrt was invoked")

	num := in.Number

	if num < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("Received invalid number: %d\n", num))
	}

	return &pb.SqrtResponse{
		Result: math.Sqrt(float64(num)),
	}, nil
}
