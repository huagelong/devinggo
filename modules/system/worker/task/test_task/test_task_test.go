// Package test_task
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package test_task

import (
	"context"
	"testing"
)

func TestTaskQueue(t *testing.T) {
	testTask := New()
	testTask.Send(context.Background(), &TestTaskData{Name: "helloworld"})
}
