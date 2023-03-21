package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	goTestDb "go_private_chain/internal/controller/go_test_db"
	"go_private_chain/internal/service"
	"go_private_chain/internal/timed"
	"log"
)

var (
	// Main 主启动命令
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of simple goframe demos",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			timed.UpdateLibrary()
			//gsvc.SetRegistry(etcd.New(`127.0.0.1:2379`))
			s := g.Server(`hello.svc`)
			s.Use(ghttp.MiddlewareHandlerResponse)
			log.Println("来了来了来了来了来了来了来了来了来了1")
			s.Group("/", func(group *ghttp.RouterGroup) {
				log.Println("来了来了来了来了来了来了来了来了来了2")

				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)
				log.Println("来了来了来了来了来了来了来了来了来了3")

				// 注册中间件
				group.Bind(
					goTestDb.New(),
				)
			})
			s.Run()
			return nil
		},
	}
)
