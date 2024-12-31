// Package cmd
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cmd

import (
	_ "devinggo/modules/_/worker"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/worker/cron"
	"devinggo/modules/system/pkg/worker/server"
	"devinggo/modules/system/pkg/worker/task"
	"context"
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
			utils.SafeGo(ctx, func(ctx context.Context) {
				server.Run(ctx)
			})

			utils.SafeGo(ctx, func(ctx context.Context) {
				cron.Run(ctx)
			})

			serverWg.Add(1)

			// 信号监听
			signalListen(ctx, signalHandlerForOverall)

			<-serverCloseSignal
			task.GetClient(ctx).Close()
			g.Log().Debug(ctx, "worker server successfully closed ..")
			serverWg.Done()
			return
		},
	}
)
