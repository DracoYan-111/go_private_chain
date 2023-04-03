package service

import (
	"context"
	"math/big"
)

type (
	IUserData interface {
		CreateUserAddress(ctx context.Context, req string) (string, error)
		BatchCastingNft(ctx context.Context, req string) (string, []*big.Int, error)
	}
)

var (
	localUserData IUserData
)

func UserData() IUserData {
	if localUserData == nil {
		panic("implement not found for interface IUserData, forgot register?")
	}
	return localUserData
}

func RegisterUserData(i IUserData) {
	localUserData = i
}
