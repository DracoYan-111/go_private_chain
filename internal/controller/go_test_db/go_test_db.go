package go_test_db

import (
	"context"
	"go_private_chain/api/v1"
	"go_private_chain/internal/service"
	"log"
)

type Controller struct{}

func New() *Controller {
	return &Controller{}
}

// SignUp is the API for user sign up.
func (c *Controller) SignUp(ctx context.Context, req *v1.NewJobTaskReq) (res *v1.NewJobTaskRes, err error) {
	log.Println("++++++++++++++++++")
	err = service.GoTestDb().Create(ctx, req.Ciphertext)
	return
}
