# gRPC 快速上手

## 入门
### 1. 定义一个 gRPC 服务 (proto 文件 )
在项目文件夹下建 proto\hellworld 目录，然后在目录下创建一个 helloworld.proto 文件，内容如下：
```proto
syntax = "proto3";

option go_package = "grpc-demo/helloworld";
option java_multiple_files = true;
option java_package = "grpc.helloworld";
option java_outer_classname = "HelloWorldProto";

package helloworld;

// The greeting service definition.
service Greeter {
  // Sends a greeting
  rpc SayHello (HelloRequest) returns (HelloReply) {}
}

// The request message containing the user's name.
message HelloRequest {
  string name = 1;
}

// The response message containing the greetings
message HelloReply {
  string message = 1;
}
```
### 2. 生成 gRPC 代码
进入proto\helloworld目录，执行以下命令，生成 gRPC 代码:
```shell
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative helloworld.proto
```

注意:  
需要提前安装 protoc 工具，以及 protoc-gen-go 插件，具体参考：https://grpc.io/docs/languages/go/quickstart/  
protoc 工具下载地址：
https://github.com/protocolbuffers/protobuf/releases   
安装 protoc-gen-go 插件：
```shell
go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2
```
将 protoc-gen-go、 protoc-gen-go-grpc、protoc 都放到一起，并加入到环境变量中，这样就可以在任意目录下执行 protoc 命令了。


生成了两个文件，分别是 helloworld.pb.go 和 helloworld_grpc.pb.go，这两个文件是 gRPC 的代码，我们可以直接使用它们来实现 gRPC 服务。

其中在 helloworld_grpc.pb.go 文件中，我们可以看到 GreeterServer 和 GreeterClient 两个接口定义，分别是服务端和客户端的接口定义。

GreeterServer 接口定义：
```go
type GreeterServer interface {
	// Sends a greeting
	SayHello(context.Context, *HelloRequest) (*HelloReply, error)
	mustEmbedUnimplementedGreeterServer()
}
```

GreeterClient 接口定义：
```go
type GreeterClient interface {
	// Sends a greeting
	SayHello(ctx context.Context, in *HelloRequest, opts ...grpc.CallOption) (*HelloReply, error)
}
```

下面就可以开始实现 gRPC 服务了。

### 3. 实现 gRPC 服务
在项目文件夹下创建 grpc_server\helloworld 目录，然后在目录下创建一个 server.go 文件，内容如下：
```go
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
```

### 4. 编写启动 gRPC 服务的代码
```go
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
```

### 5. 启动 gRPC 服务
```shell
go run grpc_server\server.go
```

### 6. 测试 gRPC 服务
```shell
# 安装 grpcurl
go install github.com/fullstorydev/grpcurl/cmd/grpcurl@latest
```
```shell
# grpcurl 命令测试 gRPC 服务
grpcurl -plaintext localhost:50051 list
```
输出如下：
```
grpc.reflection.v1.ServerReflection
grpc.reflection.v1alpha.ServerReflection
helloworld.Greeter
```

```shell
# grpcurl 测试 SayHello 接口
grpcurl -plaintext -d '{"name": "grpc"}' localhost:50051 helloworld.Greeter/SayHello
```


### 7. 编写客户端代码
在项目文件夹下创建 grpc_client\helloworld 目录，然后在目录下创建一个 client.go 文件，内容如下：
```go
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
```
