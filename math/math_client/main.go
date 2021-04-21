package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_tourist/math/mathpb"
	"io"
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
		log.Printf("failed to call Sum: %v", err)
	}

	fmt.Printf("response:\n")
	fmt.Printf(" - %v\n", resp.Result)
}

func serverSideStreamingCall(c pb.MathClient) {
	fmt.Printf("--- gRPC Server-side Streaming RPC Call ---\n")
	req := &pb.PrimeFactorsRequest{Num: 48}
	stream, err := c.PrimeFactors(context.Background(), req)
	if err != nil {
		log.Fatalf("failed to call PrimeFactors: %v", err)
	}

	var rpcStatus error
	fmt.Printf("response:\n")
	for {
		resp, err := stream.Recv()
		if err != nil {
			rpcStatus = err
			break
		}
		fmt.Printf(" - %v\n", resp.Result)
	}

	if rpcStatus != io.EOF {
		log.Fatalf("failed to finish server-side streaming: %v", rpcStatus)
	}

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

	// Contact the server and print out its response.
	// 1 Unary RPC Call
	// unaryCall(c)
	// 2 Server-side Streaming RPC Call
	serverSideStreamingCall(c)
}
