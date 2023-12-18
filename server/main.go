package main

import (
	"calc/pb"
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCalculatorServer
}

func (s *server) Add(
	ctx context.Context, in *pb.CalcRequest,
) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{
		Result: in.A + in.B,
	}, nil
}

func (s *server) Sub(
	ctx context.Context, in *pb.CalcRequest,
) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{
		Result: in.A - in.B,
	}, nil
}

func (s *server) Multiply(
	ctx context.Context, in *pb.CalcRequest,
) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{
		Result: in.A * in.B,
	}, nil
}

func (s *server) Divide(
	ctx context.Context, in *pb.CalcRequest,
) (*pb.CalcResponse, error) {
	return &pb.CalcResponse{
		Result: in.A / in.B,
	}, nil
}

func main() {
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("failed to create listener", err)
	}

	s := grpc.NewServer()
	reflection.Register(s)

	pb.RegisterCalculatorServer(s, &server{})
	if err := s.Serve(listener); err != nil {
		log.Fatal("failed to serve", err)
	}
}
