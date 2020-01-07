package main

import (
	"fmt"
	"net"
	
	pb "gRPC_demo/helloworld/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct{}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {
	// 监听本地的 9652 端口
	lis, err := net.Listen("tcp", ":9652")
	if err != nil {
		fmt.Printf("failed to listen: %v", err)
		return
	}

	// 创建gRPC服务器
	s := grpc.NewServer()
	// 在gRPC服务端注册服务
	pb.RegisterGreeterServer(s, &server{})

	// 在给定的gRPC服务器上注册服务器反射服务
	reflection.Register(s)
	// Serve方法在lis上接受传入连接，为每个连接创建一个 ServerTransport和server的goroutine
	// 该goroutine读取gRPC请求，然后调用已注册的处理程序来响应他们。
	err = s.Serve(lis)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
