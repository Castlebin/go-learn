package route

import (
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/product"
)

func SetUpRoute(engine *gin.Engine) {
	ProductRouter := engine.Group("")
	{
		// 新增产品
		ProductRouter.POST("/product", product.Add)

		// 更新产品
		ProductRouter.PUT("/product/:id", product.Edit)

		// 删除产品
		ProductRouter.DELETE("/product/:id", product.Delete)

		// 获取产品详情
		ProductRouter.GET("/product/:id", product.Detail)

		// 测试下捕获异常   curl --location 'http://localhost:8080/product/panic'
		ProductRouter.GET("/product/panic", product.Panic)
	}
}
