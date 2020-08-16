package main

import (
	"context"
	"log"
	"net"

	"github.com/bradleybonitatibus/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

type calcServer struct{}

func (*calcServer) Sum(ctx context.Context, req *calculatorpb.SumRequest) (*calculatorpb.SumResponse, error) {
	return &calculatorpb.SumResponse{
		Answer: int64(req.GetA() + req.GetB()),
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	s := grpc.NewServer()

	calculatorpb.RegisterCalculatorServiceServer(s, &calcServer{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to bind to server :%v", err)
	}
}
