package main

import (
	"context"
	"fmt"
	"os"

	"github.com/bradleybonitatibus/grpc-course/calculator/calculatorpb"
	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		fmt.Println(err.Error())
	}

	defer con.Close()

	client := calculatorpb.NewCalculatorServiceClient(con)
	callSum(client)
}

func callSum(client calculatorpb.CalculatorServiceClient) {
	req := &calculatorpb.SumRequest{
		A: 420,
		B: 69,
	}

	res, err := client.Sum(context.Background(), req)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	fmt.Println(fmt.Sprintf("The sum of %v and %v is %v", req.GetA(), req.GetB(), res.GetAnswer()))
}
