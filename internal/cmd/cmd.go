// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"
	"devinggo/modules/system/pkg/utils"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Main = &gcmd.Command{
		Description: `默认启动所有服务`,
		Arguments: []gcmd.Argument{
			{
				Name:   "config",
				Short:  "c",
				Brief:  "config file (default config.yaml)",
				IsArg:  false,
				Orphan: false,
			},
		},
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			return All.Func(ctx, parser)
		},
	}

	Help = &gcmd.Command{
		Name:  "help",
		Brief: "查看帮助",
		Description: `
		命令提示符
		---------------------------------------------------------------------------------
		启动服务
		>> 所有服务  [go run main.go]   热编译  [gf run main.go]
		>> 初始化配置以及其他必要文件  [go run main.go unpack]
		>> 迁移数据库文件  [go run main.go migrate:xxx]
		  migrate:create -name name Create a set of timestamped up/down migrations titled NAME
		  migrate:goto -v version      Migrate to version V
		  migrate:up [-n N]       Apply all or N up migrations
		  migrate:down [-n N]     Apply all or N down migrations
		  migrate:force v version  Set version V but don't run migration (ignores dirty state)
		>> HTTP服务  [go run main.go http]
		>> 消息队列&定时任务  [go run main.go worker]
		>> 查看帮助  [go run main.go help]
		>> 查看版本  [go run main.go version]
    `,
	}

	All = &gcmd.Command{
		Name:        "all",
		Brief:       "start all server",
		Description: "this is the command entry for starting all server",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)
			g.Log().Debug(ctx, "starting all server")
			// 需要启动的服务
			var allServers = []*gcmd.Command{Http, Worker}

			for _, server := range allServers {
				var cmd = server
				utils.SafeGo(ctx, func(ctx context.Context) {
					if err := cmd.Func(ctx, parser); err != nil {
						g.Log().Panicf(ctx, "%v start fail:%v", cmd.Name, err)
					}
				})
			}
			// 信号监听
			signalListen(ctx, signalHandlerForOverall)
			<-serverCloseSignal
			serverWg.Wait()
			g.Log().Debug(ctx, "all service successfully closed ..")
			return
		},
	}
)

func init() {
	if err := Main.AddCommand(All, Http, Version, Worker, Unpack, MigrateUp, MigrateDown, MigrateGoto, MigrateCreate, MigrateForce, Help); err != nil {
		panic(err)
	}
}
