package main

import (
	pb "commons/api"
	"context"
	"log"

	"google.golang.org/grpc"
)

type grpcHandler struct {
	pb.UnimplementedOrderServiceServer
}

func NewGRPCHandler(grpcServer *grpc.Server) *grpcHandler {
	handler := &grpcHandler{}
	pb.RegisterOrderServiceServer(grpcServer, handler)
	return &grpcHandler{}
}

func (h *grpcHandler) CreateOrder(ctx context.Context, p *pb.CreateOrderRequest) (*pb.Order, error) {
	log.Println("new order received! order %v", p)
	o := &pb.Order{
		ID: "42",
	}
	return o, nil
}
