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
	"github.com/hibiken/asynq"
)

var testWorker = &cTestWorker{
	Type: consts.TEST_TASK,
}

type cTestWorker struct {
	Type string
}

func init() {
	server.Register(testWorker)
}

func (s *cTestWorker) GetType() string {
	return s.Type
}

// Execute 执行任务
func (s *cTestWorker) Execute(ctx context.Context, t *asynq.Task) (err error) {
	data, err := glob2.GetParamters[cron.TestCronData](ctx, t)
	if err != nil {
		return err
	}
	glob2.WithWorkLog().Infof(ctx, `jsonData:%+v`, data)
	return
}
