package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"go_private_chain/internal/model/entity"
)

type GetUserAddressReq struct {
	g.Meta   `path:"/user/get-user-address" method:"get" tags:"GetUserAddress" summary:"查询用户的账户地址"`
	Nickname string `v:"required" summary:"用户昵称"`
}

type GetUserAddressRes struct {
	OK             bool   `v:"required" summary:"任务状态"`
	UserAddress    string `v:"required" summary:"用户对应的账户地址"`
	UserCreateHash string `v:"required" summary:"用户地址的hash"`
}

type GetUserCollectionReq struct {
	g.Meta   `path:"/user/get-user-address" method:"get" tags:"GetUserAddress" summary:"查询用户的账户地址"`
	Nickname string `v:"required" summary:"用户昵称"`
}

type GetUserCollectionRes struct {
	OK             bool                     `v:"required" summary:"任务状态"`
	UserCollection []entity.UserCollections `v:"required" summary:"用户对应的账户地址"`
}
