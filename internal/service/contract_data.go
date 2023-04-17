package service

import (
	"context"
)

type (
	IContractData interface {
		CreateJob(ctx context.Context, req string) error
	}
)

var (
	localContract IContractData
)

func GoTestDb() IContractData {
	if localContract == nil {
		panic("implement not found for interface IGoTestDb, forgot register?")
	}
	return localContract
}

func RegisterGoTestDb(i IContractData) {
	localContract = i
}
