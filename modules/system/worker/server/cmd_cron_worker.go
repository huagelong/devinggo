// Package server
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package server

import (
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/server"
	"devinggo/modules/system/worker/consts"
	"devinggo/modules/system/worker/cron"
	"context"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/hibiken/asynq"
)

var cmdCronWorker = &cCmdCronWorker{
	Type: consts.CMD_CRON,
}

type cCmdCronWorker struct {
	Type string
}

func init() {
	server.Register(cmdCronWorker)
}

func (s *cCmdCronWorker) GetType() string {
	return s.Type
}

// Execute 执行任务
func (s *cCmdCronWorker) Execute(ctx context.Context, t *asynq.Task) (err error) {
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

	return
}
