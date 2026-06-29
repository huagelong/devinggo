// Package worker
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package worker

import (
	"context"

	"devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/task"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/hibiken/asynq"
)

// CronBuilder 定时任务构建器
type CronBuilder struct {
	taskType    string
	description string
	setParams   func(ctx context.Context, payload *glob.Payload, params *gjson.Json)
}

// NewCronBuilder 创建定时任务构建器
// taskType: 任务类型
// description: 任务描述
func NewCronBuilder(taskType, description string) *CronBuilder {
	return &CronBuilder{
		taskType:    taskType,
		description: description,
	}
}

// WithParamsHandler 设置参数处理函数
// 该函数用于从params中解析数据并设置到payload.Data中
func (b *CronBuilder) WithParamsHandler(handler func(ctx context.Context, payload *glob.Payload, params *gjson.Json)) *CronBuilder {
	b.setParams = handler
	return b
}

// Register 注册定时任务到全局Manager
func (b *CronBuilder) Register() {
	cronTask := &cronTaskImpl{
		taskType:    b.taskType,
		description: b.description,
		setParams:   b.setParams,
		payload:     &glob.Payload{},
	}
	defaultManager.RegisterCronTask(cronTask)
}

// cronTaskImpl CronTask接口的实现
type cronTaskImpl struct {
	taskType    string
	description string
	payload     *glob.Payload
	setParams   func(ctx context.Context, payload *glob.Payload, params *gjson.Json)
}

func (c *cronTaskImpl) GetType() string {
	return c.taskType
}

func (c *cronTaskImpl) GetPayload() *glob.Payload {
	return c.payload
}

func (c *cronTaskImpl) GetCronTask() *asynq.Task {
	return task.GetTask(c)
}

func (c *cronTaskImpl) GetDescription() string {
	return c.description
}

func (c *cronTaskImpl) SetParams(ctx context.Context, params *gjson.Json) {
	if c.setParams != nil {
		c.setParams(ctx, c.payload, params)
	}
}

// RegisterCronFunc 函数式注册定时任务（最简洁的方式）
// taskType: 任务类型
// description: 任务描述
// paramsHandler: 参数处理函数，可为nil（从params解析数据到payload.Data）
func RegisterCronFunc(taskType, description string, paramsHandler func(ctx context.Context, payload *glob.Payload, params *gjson.Json)) {
	NewCronBuilder(taskType, description).
		WithParamsHandler(paramsHandler).
		Register()
}
