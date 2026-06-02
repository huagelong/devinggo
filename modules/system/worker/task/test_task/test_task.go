// Package test_task
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package test_task

import (
	"context"
	"time"

	"devinggo/modules/system/pkg/worker"
	"devinggo/modules/system/worker/consts"
)

// TestTaskData \u6d4b\u8bd5\u4efb\u52a1\u6570\u636e\u7ed3\u6784
type TestTaskData struct {
	Name string `json:"name"`
}

// Send \u53d1\u9001\u6d4b\u8bd5\u4efb\u52a1\uff08\u65b0\u7248\uff09
func Send(ctx context.Context, data TestTaskData) error {
	return worker.NewTaskBuilder(ctx, consts.TEST_TASK).
		WithData(data).
		Send()
}

// SendWithDelay \u53d1\u9001\u5ef6\u8fdf\u6d4b\u8bd5\u4efb\u52a1
func SendWithDelay(ctx context.Context, data TestTaskData, delay int) error {
	return worker.NewTaskBuilder(ctx, consts.TEST_TASK).
		WithData(data).
		WithDelay(time.Duration(delay) * time.Second).
		Send()
}
