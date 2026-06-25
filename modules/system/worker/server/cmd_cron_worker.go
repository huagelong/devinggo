// Package server
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package server

import (
	"context"

	"devinggo/modules/system/pkg/worker"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/worker/consts"
	"devinggo/modules/system/worker/cron"

	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/hibiken/asynq"
)

func init() {
	// 使用新方式注册Worker
	worker.RegisterWorkerFunc(consts.CMD_CRON, executeCmdCronWorker)
}

// executeCmdCronWorker 执行命令Worker
func executeCmdCronWorker(ctx context.Context, t *asynq.Task) error {
	data, err := glob2.GetParamters[cron.CmdCronData](ctx, t)
	if err != nil {
		return err
	}
	glob2.WithWorkLog().Infof(ctx, `type:%s, jsonData:%+v`, t.Type(), data)

	r, err := gproc.ShellExec(gctx.New(), data.Cmd)
	if err != nil {
		return err
	}
	glob2.WithWorkLog().Infof(ctx, `type:%s, response:%+v`, t.Type(), r)

	return nil
}
