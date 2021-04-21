package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_tourist/math/mathpb"
	"log"
)

func unaryCall(c pb.MathClient) {
	fmt.Printf("--- gRPC Unary RPC Call ---\n")
	req := &pb.SumRequest{
		FirstNum:  -5,
		SecondNum: 6,
	}
	resp, err := c.Sum(context.Background(), req)
	if err != nil {
		fmt.Printf("failed to call Sum: %v", err)
	}

	fmt.Printf("response:\n")
	fmt.Printf(" - %v\n", resp.Result)
}

func main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	c := pb.NewMathClient(conn)
	unaryCall(c)
}
