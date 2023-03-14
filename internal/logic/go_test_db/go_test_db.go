package go_test_db

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"go_private_chain/internal/dao"
	"go_private_chain/internal/service"
)

// 创建type类型的sGoTestDb结构体
type (
	sGoTestDb struct{}
)

// init方法中调用service.RegisterGoTestDb方法,传入的参数为service.IGoTestDb类型的指针
func init() {
	service.RegisterGoTestDb(New())
}

// New 方法返回service.IGoTestDb类型的指针
func New() service.IGoTestDb {
	return &sGoTestDb{}
}

// Create 方法被s *sUser调用,传入的参数为context.Context,model.UserCreateInput
func (s *sGoTestDb) Create(ctx context.Context, req string) (err error) {

	return dao.GoTestDb.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		_, err = dao.GoTestDb.Ctx(ctx).Data(req).Batch(len(req)).Insert()
		return err
	})
}
