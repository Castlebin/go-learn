package product

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"go-gin-api/app/controller/param_bind"
	"go-gin-api/app/util"
)

func Add(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	form := &param_bind.ProductAdd{}
	if err := c.ShouldBind(form); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "", nil)
}

func Edit(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	form := &param_bind.ProductAdd{}
	if err := c.ShouldBind(form); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "", nil)
}

func Delete(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	form := &param_bind.ProductAdd{}
	if err := c.ShouldBind(form); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "", nil)
}

func Detail(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}
	form := &param_bind.ProductAdd{}
	if err := c.ShouldBind(form); err != nil {
		utilGin.Response(-1, err.Error(), nil)
		return
	}
	utilGin.Response(0, "", nil)
}

func Panic(c *gin.Context) {
	utilGin := util.Gin{Ctx: c}

	a := 1
	b := 0
	c1 := a / b
	fmt.Println(c1)

	utilGin.Response(0, "", nil)
}
