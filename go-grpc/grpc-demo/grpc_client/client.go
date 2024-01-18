package main

import (
	"context"
	"fmt"
	"grpc-demo/proto/helloworld"
	"log"

	"google.golang.org/grpc"
)

func main() {
	// 连接 gRPC 服务
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	// 初始化具体的 gRPC 业务服务客户端
	client := helloworld.NewGreeterClient(conn)
	// 调用具体的 gRPC 业务服务的接口
	reply, err := client.SayHello(context.Background(), &helloworld.HelloRequest{Name: "zhangsan峰"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(reply.Message)
}
