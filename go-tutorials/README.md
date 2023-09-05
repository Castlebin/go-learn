# Go 常用命令

1. go run .
可以直接运行当前模块下的 main 函数   
如果想直接运行某个文件，比如 hello.go ，也可以直接使用  `go run hello.go`


2. go mod init xxx/yyy            
用于初始化一个模块，调用后，会在当前目录下生成一个 go.mod 文件，记录模块信息及使用的go版本


3. go mod tidy 
如果使用了其他外部包，使用该命令，会在模块信息描述文件 go.mod 中添加模块依赖信息，并且生成或者更新 go.sum 文件，描述了对模块的详细依赖（认证信息）

4. go mod edit -replace github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module/greetings=../greetings
一个模块，依赖另一个当前工程下的自定义模块，需要告诉所依赖的这个模块名字和它的目录位置。这样才能知道，依赖的是哪里的模块
运行该命令之后，会在当前模块的 go.mod 文件里生成模块的依赖信息。例如：  
replace github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module => ../greetings  

    同样的，在这之后，应该再运行 `go mod tidy` ，会在 go.mod 文件中，生成依赖模块的版本信息（当前工程的模块，所以这里是一个虚拟版本号）
    require github.com/Castlebin/go-learn/go-tutorials/Create_a_Go_module/greetings v0.0.0-00010101000000-000000000000
