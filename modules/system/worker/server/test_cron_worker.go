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

	"github.com/hibiken/asynq"
)

func init() {
	// 使用新方式注册Worker
	worker.RegisterWorkerFunc(consts.TEST_CRON, executeTestCronWorker)
}

// executeTestCronWorker 执行测试Cron Worker
func executeTestCronWorker(ctx context.Context, t *asynq.Task) error {
	data, err := glob2.GetParamters[cron.TestCronData](ctx, t)
	if err != nil {
		return err
	}
	glob2.WithWorkLog().Infof(ctx, `jsonData:%+v`, data)
	return nil
}
