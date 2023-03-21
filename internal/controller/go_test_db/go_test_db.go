package go_test_db

import (
	"context"
	v1 "go_private_chain/api/v1"
	"go_private_chain/internal/service"
	"log"
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

// NewTestTask 是上传新任务的API。
func (c *Controller) NewTestTask(ctx context.Context, req *v1.NewTestReq) (res *v1.NewTestRes, err error) {
	log.Println("来了来了来了来了来了来了来了来了来了4")

	res = &v1.NewTestRes{
		OK: "啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦啦",
	}
	return
}
