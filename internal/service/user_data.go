package service

import (
	"context"
)

type (
	IUserData interface {
		CreateUserAddress(ctx context.Context, req string) (string, error)
		BatchCastingNft(ctx context.Context, req string) (string, []string, error)
		BatchTransferNft(ctx context.Context, req string) (string, []string, error)
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
