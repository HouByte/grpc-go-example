package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	pb "grpc-go-example/example/server/proto"
	"net"
)

// hello server
type server struct {
	pb.UnimplementedSayHelloServer
}

func (s *server) SayHello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {

	return &pb.HelloResponse{Msg: "hello " + request.Name}, nil
}

func main() {
	//开启端口
	listen, _ := net.Listen("tcp", ":9090")
	// 创建grpc服务
	grpcServer := grpc.NewServer()
	//注册服务到rpc
	pb.RegisterSayHelloServer(grpcServer, &server{})

	// 启动服务
	err := grpcServer.Serve(listen)
	if err != nil {
		fmt.Printf("failed to serve: %v", err)
		return
	}
}
