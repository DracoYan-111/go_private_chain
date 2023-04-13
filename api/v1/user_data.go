package v1

import (
	"github.com/gogf/gf/v2/frame/g"
)

type NewUserAddressReq struct {
	g.Meta     `path:"/user/new-user-address" method:"post" tags:"CreateUserAccounts" summary:"创建一个用户对应的新账户"`
	Ciphertext string `v:"required" summary:"账户密文"`
}

type NewUserAddressRes struct {
	OK          bool   `v:"required" summary:"任务状态"`
	UserAddress string `v:"required" summary:"用户对应的账户地址"`
}

type NewBatchCastNftReq struct {
	g.Meta     `path:"/user/new-added-task" method:"post" tags:"NewNftIssued" summary:"进行一个新的藏品增发任务"`
	Ciphertext string `v:"required" summary:"json格式密文"`
}

type NewBatchCastNftRes struct {
	OK      bool     `v:"required" summary:"任务状态"`
	Hash    string   `v:"required" summary:"交易hash"`
	IdArray []string `v:"required" summary:"返回的id数组"`
}

type NewTransferNftReq struct {
	g.Meta     `path:"/user/new-transfer-task" method:"post" tags:"NewNftTransfer" summary:"进行一个新的藏品转移任务"`
	Ciphertext string `v:"required" summary:"json格式密文"`
}

type NewTransferNftRes struct {
	OK      bool     `v:"required" summary:"任务状态"`
	Hash    string   `v:"required" summary:"交易hash"`
	IdArray []string `v:"required" summary:"转移成功的id数组"`
}
