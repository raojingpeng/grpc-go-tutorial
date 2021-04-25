package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	pb "grpc_tourist/features/echopb"
	"log"
	"os"
)

func main() {
	addr := flag.String("addr", "localhost:50051", "the address to connect to")
	flag.Parse()

	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewEchoClient(conn)
	msg := "Madman"
	if len(os.Args) > 1 {
		msg = os.Args[1]
	}
	resp, err := client.UnaryEcho(context.Background(), &pb.EchoRequest{Message: msg})
	if err != nil {
		//log.Fatalf("failed to call UnaryEcho: %v", err)
		errStatus, _ := status.FromError(err)
		fmt.Printf("Error Code: %v\n", errStatus.Code())
		fmt.Printf("Error Description: %v\n\n", errStatus.Message())

		if codes.InvalidArgument == errStatus.Code() {
			fmt.Println("You can take specific action based on specific error!")
		}

	}
	fmt.Printf("response:\n")
	fmt.Printf(" - %v\n", resp.GetMessage())
}