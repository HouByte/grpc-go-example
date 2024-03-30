package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	pb "grpc-go-example/example/client/proto"
	"log"
)

func main() {
	// 连接服务端，这里先不进行加密处理
	conn, err := grpc.Dial("127.0.0.1:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	defer conn.Close()

	// 建立连接
	client := pb.NewSayHelloClient(conn)

	//远程调用请求
	resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "go"})
	if err != nil {
		return
	}

	fmt.Println(resp.Msg)
}
