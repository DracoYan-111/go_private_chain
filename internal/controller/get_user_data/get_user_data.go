// Package get_user_data 针对获取用户信息的控制层
package get_user_data

import (
	"context"
	v1 "go_private_chain/api/v1"
	"go_private_chain/internal/service"
)

type (
	UserController struct{}
)

func New() *UserController {
	return &UserController{}
}

// GetUserAddress 通过昵称查询用户地址
func (c *UserController) GetUserAddress(ctx context.Context, req *v1.GetUserAddressReq) (res *v1.GetUserAddressRes, err error) {
	userData, err := service.GetUserData().GetUserAddress(ctx, req.Nickname)
	var userAddress string
	var userCreateHash string
	if userData != nil {
		userAddress = userData[0]
		userCreateHash = userData[1]
	}
	res = &v1.GetUserAddressRes{
		OK:             err == nil,
		UserAddress:    userAddress,
		UserCreateHash: userCreateHash,
	}
	return
}

// GetUserCollection 通过昵称查询用户拥有的藏品
func (c *UserController) GetUserCollection(ctx context.Context, req *v1.GetUserCollectionReq) (res *v1.GetUserCollectionRes, err error) {
	userCollectionRes, err := service.GetUserData().GetUserCollection(ctx, req.Nickname)
	res = &v1.GetUserCollectionRes{
		OK:             err == nil,
		UserCollection: userCollectionRes,
	}
	return
}
