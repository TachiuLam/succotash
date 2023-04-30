package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	pb "github.com/TachiuLam/succotash/grpc/protobuf"
	"github.com/TachiuLam/succotash/grpc/service"
)

func StartServer() {
	//开启服务端监听
	lis, err := net.Listen("tcp", "127.0.0.1:8090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	//创建grpcServer 将拦截器添加进去
	gRpcServer := grpc.NewServer([]grpc.ServerOption{grpc.UnaryInterceptor(NewServerIntercept())}...)

	//注册服务实现
	pb.RegisterGreeterServer(gRpcServer, &service.GreeterService{})
	//启动服务
	gRpcServer.Serve(lis)
}

func main() {
	StartServer()
}
