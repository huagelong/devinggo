// Package worker
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package worker

import (
	"context"
	"time"

	"devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/task"

	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hibiken/asynq"
)

// TaskBuilder 任务构建器
type TaskBuilder struct {
	taskType string
	payload  *glob.Payload
	data     interface{}
	ctx      context.Context
}

// NewTaskBuilder 创建任务构建器
func NewTaskBuilder(ctx context.Context, taskType string) *TaskBuilder {
	return &TaskBuilder{
		taskType: taskType,
		ctx:      ctx,
		payload: &glob.Payload{
			QueueName: "default",
			TaskID:    gconv.String(gtime.TimestampNano()),
			Time:      asynq.ProcessIn(0),
		},
	}
}

// WithData 设置任务数据
func (b *TaskBuilder) WithData(data interface{}) *TaskBuilder {
	b.data = data
	return b
}

// WithQueue 设置队列名称
func (b *TaskBuilder) WithQueue(queueName string) *TaskBuilder {
	b.payload.QueueName = queueName
	return b
}

// WithTaskID 设置任务ID（用于唯一标识）
func (b *TaskBuilder) WithTaskID(taskID string) *TaskBuilder {
	b.payload.TaskID = taskID
	return b
}

// WithDelay 设置延迟执行时间
func (b *TaskBuilder) WithDelay(delay time.Duration) *TaskBuilder {
	b.payload.Time = asynq.ProcessIn(delay)
	return b
}

// WithProcessAt 设置在指定时间执行
func (b *TaskBuilder) WithProcessAt(t time.Time) *TaskBuilder {
	b.payload.Time = asynq.ProcessAt(t)
	return b
}

// WithRetention 设置任务保留时间
func (b *TaskBuilder) WithRetention(retention time.Duration) *TaskBuilder {
	b.payload.Retention = asynq.Retention(retention)
	return b
}

// WithCrontabID 设置定时任务ID
func (b *TaskBuilder) WithCrontabID(id int64) *TaskBuilder {
	b.payload.CrontabId = id
	return b
}

// Send 发送任务
func (b *TaskBuilder) Send() error {
	b.payload.Data = b.data
	simpleTask := &simpleTaskImpl{
		taskType: b.taskType,
		payload:  b.payload,
	}
	return task.NewSimpleTask(b.ctx, simpleTask)
}

// Build 构建asynq.Task（用于自定义处理）
func (b *TaskBuilder) Build() *asynq.Task {
	b.payload.Data = b.data
	simpleTask := &simpleTaskImpl{
		taskType: b.taskType,
		payload:  b.payload,
	}
	return task.GetTask(simpleTask)
}

// simpleTaskImpl SimpleTask接口的简单实现
type simpleTaskImpl struct {
	taskType string
	payload  *glob.Payload
}

func (s *simpleTaskImpl) GetType() string {
	return s.taskType
}

func (s *simpleTaskImpl) GetPayload() *glob.Payload {
	return s.payload
}
