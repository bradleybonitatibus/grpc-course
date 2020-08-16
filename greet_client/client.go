package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/bradleybonitatibus/grpc-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {

	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Failed to connect to grpc server: %v", err)
	}

	defer con.Close()

	client := greetpb.NewGreetServiceClient(con)
	// unaryRPC(client)

	serverStream(client)

}

func unaryRPC(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Bradley",
			LastName:  "Bonitatibus",
		},
	}
	res, err := client.Greet(context.Background(), req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Printf("Message received from gRCP server: %v", res.GetResult())
}

func serverStream(client greetpb.GreetServiceClient) {
	req := &greetpb.GreetManyTimesRequest{
		Greeting: &greetpb.Greeting{
			FirstName: "Bradley",
			LastName:  "Bonitatibus",
		},
	}
	svc, err := client.GreetManyTimes(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
	}

	for {
		msg, err := svc.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			fmt.Println(err.Error())
		}

		fmt.Println(msg.GetResult())
	}
}
