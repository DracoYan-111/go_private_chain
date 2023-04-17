package contract_data

import (
	"context"
	v1 "go_private_chain/api/v1"
	"go_private_chain/internal/service"
)

type (
	JobController struct{}
)

func New() *JobController {
	return &JobController{}
}

// NewJobTask 提交新的藏品部署任务
func (c *JobController) NewJobTask(ctx context.Context, req *v1.NewJobTaskReq) (res *v1.NewJobTaskRes, err error) {

	err = service.GoTestDb().CreateJob(ctx, req.Ciphertext)

	res = &v1.NewJobTaskRes{
		OK: err == nil,
	}
	return
}
