package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"math/big"
)

type NewUserAddressReq struct {
	g.Meta     `path:"/user/new-user-address" method:"post" tags:"AddedAccounts" summary:"创建一个新账户"`
	Ciphertext string `v:"required"`
}

type NewUserAddressRes struct {
	OK          bool   `summary:"任务状态"`
	UserAddress string `summary:"用户地址"`
}

type NewBatchCastNftReq struct {
	g.Meta     `path:"/user/new-added-task" method:"post" tags:"AddedService" summary:"上传一个新的藏品增发任务"`
	Ciphertext string `v:"required"`
}

type NewBatchCastNftRes struct {
	OK     bool       `summary:"任务状态"`
	Hash   string     `summary:"交易hash"`
	IdData []*big.Int `summary:"返回的id数组"`
}
