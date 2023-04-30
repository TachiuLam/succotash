package service

import (
	"context"

	pb "github.com/TachiuLam/succotash/grpc/protobuf"
)

type GreeterService struct {
}

func (s *GreeterService) SayHello(_ context.Context, args *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + args.GetName()}, nil
}
