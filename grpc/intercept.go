package main

import (
	"context"
	"log"

	"google.golang.org/grpc"
)

// NewServerIntercept grpc  拦截器
func NewServerIntercept() grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler) (resp interface{}, err error) {
		//这个方法的具体实现
		err = record(ctx) // 中间件
		if err != nil {
			return
		}
		// 校验通过后继续处理请求
		return handler(ctx, req)
	}
}

func record(ctx context.Context) error {
	log.Printf("%s request ...", ctx.Value("user"))
	return nil
}
