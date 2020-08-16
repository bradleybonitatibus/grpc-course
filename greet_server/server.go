package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/bradleybonitatibus/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

type server struct{}

func (*server) Greet(c context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Println("Greet function was invoked with ", req.String())
	res := &greetpb.GreetResponse{
		Result: fmt.Sprintf("Hello %v", req.GetGreeting().GetFirstName()),
	}
	return res, nil
}

func (*server) GreetManyTimes(req *greetpb.GreetManyTimesRequest, stream greetpb.GreetService_GreetManyTimesServer) error {
	firstName := req.GetGreeting().GetFirstName()
	for i := 0; i < 200000; i++ {
		r := fmt.Sprintf("Hello %v number %v", firstName, i)
		res := &greetpb.GreetManyTimesResponse{
			Result: r,
		}
		stream.Send(res)
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen %v", err)
		os.Exit(1)
	}
	s := grpc.NewServer()

	greetpb.RegisterGreetServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to bind to server %v", err)
	}
}
