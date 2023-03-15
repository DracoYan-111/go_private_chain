package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"go_private_chain/internal/model"
	"go_private_chain/internal/service"
)

// 创建type类型的sMiddleware结构体
type (
	sMiddleware struct{}
)

func init() {
	service.RegisterMiddleware(New())
}

func New() service.IMiddleware {
	return &sMiddleware{}
}

// Ctx 方法被s*sMiddleware调用,传入的参数为*ghttp.Request
func (s *sMiddleware) Ctx(req *ghttp.Request) {
	customCtx := &model.Context{
		Session: req.Session,
	}
	service.TesDBCtx().Init(req, customCtx)
	// 继续执行下一个中间件。
	req.Middleware.Next()
}
