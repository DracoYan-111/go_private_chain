package user_data

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

// NewUserAddress 是生成一个新用户地址。
func (c *UserController) NewUserAddress(ctx context.Context, req *v1.NewUserAddressReq) (res *v1.NewUserAddressRes, err error) {

	userAddress, err := service.UserData().CreateUserAddress(ctx, req.Ciphertext)

	res = &v1.NewUserAddressRes{
		OK:          err == nil,
		UserAddress: userAddress,
	}
	return
}
