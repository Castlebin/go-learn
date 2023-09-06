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

5. go test  运行单测

6. go work init ./hello   
创建一个 go 工程文件，并且包含当前目录下的 ./hello 文件夹中的模块。
运行该命令后，会在目录下生成一个 go.work 文件，格式和 go.mod 类似。
这是 go 工程的工程文件。
`use ./hello` 表示该工程使用了 ./hello 目录

有了这个文件之后，你就可以直接在工程目录下执行 `go run ./hello` 来运行 hello 中的 go 程序了

7. go work use ./example/hello 为工程添加要使用的目录，给 go 的编译器来做识别的

使用 `go work` 是一种更好的工程管理方式，这样可以不用再使用 `go mod edit -replace` 命令来在 go.mod 中生成 replace 指令 将依赖本地工程中的包指向到本地工程的目录。更符合工程管理的规范

