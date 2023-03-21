package service

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"log"
)

type (
	IMiddleware interface {
		Ctx(req *ghttp.Request)
	}
)

var (
	localMiddleware IMiddleware
)

func Middleware() IMiddleware {
	log.Println("来了来了来了来了来了来了来了来了来了")
	if localMiddleware == nil {
		panic("implement not found for interface IMiddleware, forgot register?")
	}
	return localMiddleware
}

func RegisterMiddleware(i IMiddleware) {
	localMiddleware = i
}
