package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	goTestDb "go_private_chain/internal/controller/go_test_db"
	goUserData "go_private_chain/internal/controller/user_data"
	"go_private_chain/internal/service"
	"go_private_chain/internal/timed"
)

var (
	// Main 主启动命令
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server of simple goframe",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			timed.UpdateLibrary()
			//gsvc.SetRegistry(etcd.New(`127.0.0.1:2379`))
			s := g.Server()
			s.Use(ghttp.MiddlewareHandlerResponse)
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(
					service.Middleware().Ctx,
					ghttp.MiddlewareCORS,
				)

				// 注册中间件
				group.Bind(
					goTestDb.New(),
					goUserData.New(),
				)
			})
			s.Run()
			//timed.UpdateLibrary()
			return nil
		},
	}
)
