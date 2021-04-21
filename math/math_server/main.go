package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc_tourist/math/mathpb"
	"io"
	"log"
	"net"
)

type server struct {
}

func (s *server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error) {
	fmt.Printf("--- grpc Server-side Streaming RPC ---\n")
	fmt.Printf("request received: %v\n", in)
	return &pb.SumResponse{Result: in.FirstNum + in.SecondNum}, nil
}

func (s *server) PrimeFactors(in *pb.PrimeFactorsRequest, stream pb.Math_PrimeFactorsServer) error {
	fmt.Printf("--- grpc Server-side Streaming RPC ---\n")
	fmt.Printf("request received: %v\n", in)

	num := in.Num
	factor := int32(2)
	for num > 1 {
		if num%factor == 0 {
			err := stream.Send(&pb.PrimeFactorsResponse{Result: factor})
			if err != nil {
				log.Fatalf("failed to send stream: %v", err)
				return err
			}
			num = num / factor
			continue
		}
		factor += 1
	}
	return nil
}

func (s *server) Average(stream pb.Math_AverageServer) error {
	var sum int32
	count := 0
	for {
		num, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("Receiving client streaming data completed\n")
			average := float32(sum) / float32(count)
			return stream.SendAndClose(&pb.AverageResponse{Result: average})
		}

		fmt.Printf("request received: %v\n", num)
		if err != nil {
			log.Fatalf("Error while receiving client streaming data: %v", err)
		}
		sum += num.Num
		count ++
	}
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
