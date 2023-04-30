package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/TachiuLam/succotash/grpc/protobuf"
)

func StartClient() {
	//连接grpc的服务端, 取消传输证书校验
	conn, err := grpc.Dial("127.0.0.1:8090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Panic(err)
	}
	//创建客户端存根对象
	c := pb.NewGreeterClient(conn)
	// Contact the server and print out its response.
	//设置超时时间
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	//客户端往服务端发送10次消息
	for i := 0; i < 5; i++ {
		ctxChild := context.WithValue(ctx, "user", fmt.Sprintf("%d mimi", i))
		//调用服务端的方法
		r, err := c.SayHello(ctxChild, &pb.HelloRequest{
			Name: "Bob",
		})
		if err != nil {
			log.Fatalf("%v", err)
			return
		}
		log.Println(r.GetMessage())
	}

	//关闭连接
	defer conn.Close()
}

func main() {
	{
		StartClient()
	}
}
