package cmd

import (
	"context"
	"github.com/gogf/gf/contrib/registry/file/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/gsvc"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	goTestDb "go_private_chain/internal/controller/go_test_db"
	goUserData "go_private_chain/internal/controller/user_data"
	"go_private_chain/internal/service"
	"go_private_chain/internal/timed"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "Start the internal service of Jingping Cloud Chain",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			timed.UpdateLibrary()
			gsvc.SetRegistry(file.New(gfile.Temp("go_private_chain")))
			s := g.Server("jingping.chain")
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
			return nil
		},
	}
)
