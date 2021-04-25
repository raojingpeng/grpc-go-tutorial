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
	"net"
)

type server struct {
}

func (s *server) UnaryEcho(ctx context.Context, request *pb.EchoRequest) (*pb.EchoResponse, error) {
	if len(request.GetMessage()) > 10 {
		return nil, status.Errorf(codes.InvalidArgument, "Length of `Message` cannot be more than 10 characters")
	}
	return &pb.EchoResponse{Message: request.Message}, nil
}

func (s *server) ServerStreamingEcho(request *pb.EchoRequest, echoServer pb.Echo_ServerStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
}

func (s *server) ClientStreamingEcho(echoServer pb.Echo_ClientStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
}

func (s *server) BidirectionalStreamingEcho(echoServer pb.Echo_BidirectionalStreamingEchoServer) error {
	return status.Errorf(codes.Unimplemented, "method UnaryEcho not implemented")
}

func main() {
	port := flag.Int("port", 50051, "the port to server on")
	flag.Parse()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Printf("server listening at %v\n", lis.Addr())

	s := grpc.NewServer()
	pb.RegisterEchoServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Printf("failed to server: %v", err)
	}
}
