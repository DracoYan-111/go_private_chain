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
	localContractData IContractData
)

func ContractData() IContractData {
	if localContractData == nil {
		panic("implement not found for interface IContractData, forgot register?")
	}
	return localContractData
}

func RegisterContractData(i IContractData) {
	localContractData = i
}
