package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	pb "github.com/atheerauribi/server2/proto"
)

type calculatorServer struct {
	pb.UnimplementedCalculatorServer
}

func (s *calculatorServer) Add(ctx context.Context, req *pb.AddRequest) (*pb.OperationResponse, error) {
	return &pb.OperationResponse{Result: req.Number1 + req.Number2}, nil
}

func (s *calculatorServer) Subtract(ctx context.Context, req *pb.SubtractRequest) (*pb.OperationResponse, error) {
	return &pb.OperationResponse{Result: req.Number1 - req.Number2}, nil
}

func (s *calculatorServer) Multiply(ctx context.Context, req *pb.MultiplyRequest) (*pb.OperationResponse, error) {
	return &pb.OperationResponse{Result: req.Number1 * req.Number2}, nil
}

// number1 / number2
func (s *calculatorServer) Divide(ctx context.Context, req *pb.DivideRequest) (*pb.OperationResponse, error) {
	if req.Number2 == 0 {
		return nil, status.Errorf(codes.InvalidArgument, "cannot divide by zero")
	}
	return &pb.OperationResponse{Result: req.Number1 / req.Number2}, nil
}

// RunServer runs the gRPC server
func RunServer() error {
	fmt.Println("Running Server\nListening on port 8889...")
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8889))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption

	s := grpc.NewServer(opts...)
	pb.RegisterCalculatorServer(s, &calculatorServer{})
	return s.Serve(lis)
}
