一个多模块的工程，一个模块调用另一个模块

需要使用到  go mod edit -replace 命令  和  go mod tidy  命令，
让 go.mod 文件中，生成模块依赖信息

