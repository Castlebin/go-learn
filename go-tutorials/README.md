# Go 常用命令

1. go run .
可以直接运行当前模块下的 main 函数   
如果想直接运行某个文件，比如 hello.go ，也可以直接使用  `go run hello.go`


2. go mod init xxx/yyy            
用于初始化一个模块，调用后，会在当前目录下生成一个 go.mod 文件，记录模块信息及使用的go版本


3. go mod tidy 
如果使用了其他外部包，使用该命令，会在模块信息描述文件 go.mod 中添加模块依赖信息，并且生成或者更新 go.sum 文件，描述了对模块的详细依赖（认证信息）

