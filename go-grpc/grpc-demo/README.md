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

