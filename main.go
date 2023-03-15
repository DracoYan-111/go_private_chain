package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	_ "go_private_chain/internal/logic"
	_ "go_private_chain/internal/packed"

	"github.com/gogf/gf/v2/os/gctx"

	"go_private_chain/internal/cmd"
)

func main() {
	cmd.Main.Run(gctx.New())
}
