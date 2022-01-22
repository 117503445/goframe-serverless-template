package main

import (
	_ "goframe-serverless-template/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"
	"goframe-serverless-template/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
