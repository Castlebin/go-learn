package main

import (
	"context"
	"fmt"
	"grpc-demo/proto/helloworld"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// greeterServer 实现了 helloworld.GreeterServer 接口，即实线了我们 gRPC proto 文件中定义的 gRPC 服务
type greeterServer struct {
	// 这个匿名对象，是接口的默认实现（调用都是报错），如果没有实现接口的方法，就会调用这个匿名对象的方法
	helloworld.UnimplementedGreeterServer // 嵌入这个匿名对象（gRPC 这么设计，升级方便，升级的，以前的接口的实现，不会编译错误）
}

func (*greeterServer) SayHello(ctx context.Context, req *helloworld.HelloRequest) (*helloworld.HelloReply, error) {
	// 打印下请求参数
	fmt.Println(req)

	// 具体的业务逻辑
	reply := &helloworld.HelloReply{Message: "hello, " + req.Name + " !"}

	// 返回响应
	return reply, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	// 创建一个 gRPC 服务（可以有多个具体的 gRPC 业务服务注册到它上面）
	server := grpc.NewServer()
	// 注册 grpcurl 所需的 reflection 服务
	reflection.Register(server)

	// 注册具体的一个 gRPC 业务服务   (可以看到，如果有多个 gRPC 业务服务，都可以这么注册在一起)
	helloworld.RegisterGreeterServer(server, &greeterServer{})

	// 启动 gRPC 服务
	fmt.Println("grpc server start ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
