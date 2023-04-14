package service

import (
	"context"
)

type (
	IGoTestDb interface {
		CreateJob(ctx context.Context, req string) error
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
