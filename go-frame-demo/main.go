package main

import (
	_ "go-frame-demo/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go-frame-demo/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.GetInitCtx())
}
