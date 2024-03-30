# gRPC

## 简介

gRPC 是一个高性能、开源、通用的RPC框架，由Google推出，基于HTTP2协议标准设计开发，默认采用Protocol Buffers数据序列化协议，支持多种开发语言。 gRPC提供了一种简单的方法来精确的定义服务，并且为客户端和服务端自动生成可靠的功能库。

## 环境安装

### Protocol Buffers 安装

> protocol 编译器安装

1. 下载对应系统的版本：[protobuf/releases](https://github.com/protocolbuffers/protobuf/releases)
2. 解压到合适的位置
3. `bin`目录加入环境变量
4. `命令行工具`执行`protoc --version`,和自己下载版本一致说明配置成功

> 关于Protocol Buffers详细描述见：[Protocol Buffers](./docs/Protocol%20Buffers.md)

### gRPC 安装

> 国内需要配置代理地址
> 
> go env -w  GOPROXY=https://goproxy.cn,direct

安装核心库

```shell
go get google.golang.org/grpc
```

### 生成工具安装

> 前面已经安装了编译器，还需要安装相应版本的代码生成工具 
> 
> 安装新版本

```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest
```

> 安装grpc的时候已经下载下来，此时直接安装即可
> 
> 其中旧版本为`github.com/protobuf/cmd/protoc-gen-go`
> 
> 安装后会在$GOPATH下看见两个新安装的文件


## Example

### proto 编写
1. 创建目录,为了例子更清晰

```text
\---example
    +---client
    |   \---proto
    \---server
        \---proto
```

2. proto目录下编写例子hello.proto，服务端和客户端一致
```protobuf
//声明使用的语法版本是proto3语法
syntax = "proto3";

//指定生成目录和包名；.代表当前目录生成,service代表生成的go文件的包名
option go_package = ".;service";

//消息体，映射代码中的结构体
message HelloRequest {
  //定义一个字符串类型的变量， 后面的'赋值'这里代表变量在消息的位置
  string name = 1;
}

message HelloResponse {
  string msg = 1;
}

//定义一个服务
service SayHello {
  //定义rpc方法，接受参数和返回参数
  rpc SayHello(HelloRequest) returns (HelloResponse) {}
}
```
3. 生成代码

> 直接进入相应目录，生成代码
> 
> 服务端和客户端可以分别生成

```shell
protoc --go_out=. hello.proto
protoc --go-grpc_out=. hello.proto
```
> 生成后目录中多了`hello.pb.go`和`hello_grpc.pb.go`文件

### 服务端功能编写

服务端编写步骤

- 创建gRPC Server对象
- 将server注册到gRPC Server的内部注册中心。内部会进行路由处理，转发到相应的逻辑处理
- 创建Listen，监听TCP端口
- gRPC Server开始监听，直到停止

代码编写:
```go
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

```

> pb "grpc-go-example/example/server/proto" 生成rpc相应的代码步骤
> 
> 注册9090端口

### 客户端功能编写

客户端编写步骤
- 创建连接服务端
- 创建客户端对象
- 发送RPC请求，等待同步响应，得到回调后返回响应结果
- 处理相应结果

代码实现
```go
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
```