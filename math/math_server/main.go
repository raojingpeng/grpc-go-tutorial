package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_tourist/math/mathpb"
	"log"
	"net"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	return &pb.SumResponse{Result: in.FirstNum + in.SecondNum}, nil
}

func main() {
	port := flag.Int("port", 50051, "the port to server on")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listen at %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterMathServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
