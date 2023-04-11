package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"go_private_chain/internal/cmd"
	_ "go_private_chain/internal/logic"
)

func main() {
	cmd.Main.Run(gctx.New())
}
