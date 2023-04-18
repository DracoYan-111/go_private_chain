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
	userAddress, err := service.GetUserData().GetUserAddress(ctx, req.Nickname)
	res = &v1.GetUserAddressRes{
		OK:             err == nil,
		UserAddress:    userAddress[0],
		UserCreateHash: userAddress[1],
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
