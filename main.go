package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/os/gctx"
	"go_private_chain/internal/cmd"
	_ "go_private_chain/internal/logic"
)

func main() {
	//cmd.Test.Run(gctx.New())
	cmd.Main.Run(gctx.New())
	//timed.UpdateLibrary()
}

//[{"chainId":2,"id":"1605133286670282753","name":"创世纪念勋章"},{"chainId":2,"id":"1605133287098101762","name":"石氏星经"}]
//{"contract": "0x94cdfCbB76E674cc92E670aF89efc66C5865d09e","tos": ["0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65","0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266"],"tokenIds": [1, 2],"uris": ["asdfasdfasdfadf", "asdfasdfasdfadf"]}
