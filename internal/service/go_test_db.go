package service

import (
	"context"
	"go_private_chain/internal/model/entity"
)

type (
	IGoTestDb interface {
		CreateJob(ctx context.Context, req string) error
		UndoneJob() ([]*entity.GoTestDb, error)
	}
)

var (
	localGoTestDb IGoTestDb
)

func GoTestDb() IGoTestDb {
	if localGoTestDb == nil {
		panic("implement not found for interface IGoTestDb, forgot register?")
	}
	return localGoTestDb
}

func RegisterGoTestDb(i IGoTestDb) {
	localGoTestDb = i
}
