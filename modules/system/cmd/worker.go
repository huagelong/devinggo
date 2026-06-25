// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	"context"

	_ "devinggo/modules/bootstrap/worker"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/worker"
	_ "devinggo/modules/system/worker"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

var (
	Worker = &gcmd.Command{
		Name:        "worker",
		Brief:       "消息队列",
		Description: ``,
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			CmdInit(ctx, parser)

			// 获取全局Worker管理器（已在init中自动注册）
			mgr := worker.GetDefaultManager()

			// 启动Worker服务器（处理任务）
			utils.SafeGo(ctx, func(ctx context.Context) {
				if err := mgr.RunServer(); err != nil {
					g.Log().Errorf(ctx, "Worker服务器错误: %v", err)
				}
			})

			// 启动Cron调度器（定时任务）
			utils.SafeGo(ctx, func(ctx context.Context) {
				if err := mgr.RunCron(); err != nil {
					g.Log().Errorf(ctx, "Cron调度器错误: %v", err)
				}
			})

			ServerWg.Add(1)

			// 信号监听
			SignalListen(ctx, SignalHandlerForOverall)

			<-ServerCloseSignal

			// 关闭Worker客户端
			if client := mgr.GetClient(); client != nil {
				client.Close()
			}
			g.Log().Debug(ctx, "worker server successfully closed ..")
			ServerWg.Done()
			return
		},
	}
)
