package middleware

import (
	"github.com/gogf/gf/v2/net/ghttp"
)

// 创建type类型的sMiddleware结构体
type (
	sMiddleware struct{}
)

// Ctx 方法被s*sMiddleware调用,传入的参数为*ghttp.Request
func (s *sMiddleware) Ctx(r *ghttp.Request) {
	//customCtx := &model.Context{
	//	Session: r.Session,
	//}
	//service.TesDBCtx().Init(r, customCtx)
	//if service.S {
	//
	//}
}
