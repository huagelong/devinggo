// Package test_task
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package test_task

import (
	"context"
	"devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/task"
	"devinggo/modules/system/worker/consts"
	"github.com/hibiken/asynq"
)

type ctestTask struct {
	Type    string
	Payload *glob.Payload
}

type TestTaskData struct {
	Name string `json:"name"`
}

func New() *ctestTask {
	return &ctestTask{
		Type: consts.TEST_TASK,
		Payload: &glob.Payload{
			//Time:     asynq.ProcessIn(gconv.Duration("1s")),
			/*
				asynq.Retention  执行后保留一段时间删除,唯一
				asynq.ProcessIn(time.Second * 5)  等待一段时间后执行 延迟5s执行
				asynq.ProcessAt(time.Now().Add(time.Second*10)) 指定时间执行，当前时间+10s执行
			*/
			Time:   asynq.ProcessIn(0),
			TaskID: consts.TEST_TASK,
		},
	}
}

func (s *ctestTask) GetType() string {
	return s.Type
}

func (s *ctestTask) Send(ctx context.Context, data interface{}) error {
	s.Payload.Data = data
	return task.NewTask(ctx, s)
}

func (s *ctestTask) GetPayload() *glob.Payload {
	return s.Payload
}
