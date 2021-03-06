package cmd

import (
	"context"

	"goframe-serverless-template/internal/controller"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/genv"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "start http server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {

			dblink := genv.GetWithCmd("dblink").String()
			g.Log().Line().Debug(ctx, "dblink = ", dblink)

			// g.Log().Line().Debug(ctx, "dbpass = ", genv.GetWithCmd("dbpass").String())

			if dblink != "" {
				gdb.SetConfig(gdb.Config{
					"default": gdb.ConfigGroup{
						gdb.ConfigNode{
							Link: dblink,
							Type: "mysql",
						},
					},
				})
			}

			s := g.Server()
			s.Group("/", func(group *ghttp.RouterGroup) {
				group.Middleware(ghttp.MiddlewareHandlerResponse)
				group.Bind(
					controller.Hello,
					controller.Task,
				)
			})
			s.Run()

			return nil
		},
	}
)
