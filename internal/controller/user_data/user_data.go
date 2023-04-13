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

// NewUserAddress 新的用户地址生成任务
func (c *UserController) NewUserAddress(ctx context.Context, req *v1.NewUserAddressReq) (res *v1.NewUserAddressRes, err error) {

	userAddress, err := service.UserData().CreateUserAddress(ctx, req.Ciphertext)

	res = &v1.NewUserAddressRes{
		OK:          err == nil,
		UserAddress: userAddress,
	}
	return
}

// NewBatchCastNft 新的批量创建nft任务
func (c UserController) NewBatchCastNft(ctx context.Context, req *v1.NewBatchCastNftReq) (res *v1.NewBatchCastNftRes, err error) {
	hash, idData, err := service.UserData().BatchCastingNft(ctx, req.Ciphertext)
	res = &v1.NewBatchCastNftRes{
		Hash:    hash,
		IdArray: idData,
		OK:      err == nil,
	}
	return
}

// NewTransferNft 新的转移nft任务
func (c UserController) NewTransferNft(ctx context.Context, req *v1.NewTransferNftReq) (res *v1.NewTransferNftRes, err error) {
	hash, idArray, err := service.UserData().BatchTransferNft(ctx, req.Ciphertext)
	res = &v1.NewTransferNftRes{
		Hash:    hash,
		IdArray: idArray,
		OK:      err == nil,
	}
	return
}
