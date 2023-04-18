package service

import (
	"context"
	"go_private_chain/internal/model/entity"
)

type (
	IGetUserData interface {
		GetUserAddress(ctx context.Context, req string) (resData []string, err error)
		GetUserCollection(ctx context.Context, req string) (userCollectionList []entity.UserCollections, err error)
	}
)

var (
	localGetUserData IGetUserData
)

func GetUserData() IGetUserData {
	if localGetUserData == nil {
		panic("implement not found for interface IGetUserData, forgot register?")
	}
	return localGetUserData
}

func RegisterGetUserData(i IGetUserData) {
	localGetUserData = i
}
