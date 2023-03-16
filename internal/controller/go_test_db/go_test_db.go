package go_test_db

import (
	"context"
	v1 "go_private_chain/api/v1"
	"go_private_chain/internal/service"
)

type (
	Controller struct{}
)

func New() *Controller {
	return &Controller{}
}

// NewJobTask 是上传新任务的API。
func (c *Controller) NewJobTask(ctx context.Context, req *v1.NewJobTaskReq) (res *v1.NewJobTaskRes, err error) {

	err = service.GoTestDb().CreateJob(ctx, req.Ciphertext)

	res = &v1.NewJobTaskRes{
		OK: err == nil,
	}
	return
}
