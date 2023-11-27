package main

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/param_verify"
	"go-gin-api/app/route"
	"go-gin-api/middleware/logger"
	"go-gin-api/middleware/runtime_panic"
)

func main() {
	r := gin.Default()

	// 设置日志中间件（记录gin 访问日志）
	r.Use(logger.SetUp())
	// 捕获异常
	r.Use(runtime_panic.SetUp())

	// 设置路由
	route.SetUpRoute(r)

	// 绑定参数验证器
	param_verify.BindingVerify()

	r.Run() // listen and serve on 0.0.0.0:8080
}
