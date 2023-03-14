package main

import (
	_ "go_private_chain/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go_private_chain/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
