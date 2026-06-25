// Package worker
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package worker

import (
	"context"

	"github.com/hibiken/asynq"
)

// WorkerBuilder Worker构建器
type WorkerBuilder struct {
	taskType string
	execute  func(ctx context.Context, t *asynq.Task) error
}

// NewWorkerBuilder 创建Worker构建器
// taskType: 任务类型
func NewWorkerBuilder(taskType string) *WorkerBuilder {
	return &WorkerBuilder{
		taskType: taskType,
	}
}

// WithExecute 设置执行函数
func (b *WorkerBuilder) WithExecute(execute func(ctx context.Context, t *asynq.Task) error) *WorkerBuilder {
	b.execute = execute
	return b
}

// Register 注册Worker到全局Manager
func (b *WorkerBuilder) Register() {
	worker := &workerImpl{
		taskType: b.taskType,
		execute:  b.execute,
	}
	defaultManager.RegisterWorker(worker)
}

// workerImpl Worker接口的实现
type workerImpl struct {
	taskType string
	execute  func(ctx context.Context, t *asynq.Task) error
}

func (w *workerImpl) GetType() string {
	return w.taskType
}

func (w *workerImpl) Execute(ctx context.Context, t *asynq.Task) error {
	if w.execute != nil {
		return w.execute(ctx, t)
	}
	return nil
}

// RegisterWorkerFunc 函数式注册Worker（最简洁的方式）
// taskType: 任务类型
// execute: 执行函数
func RegisterWorkerFunc(taskType string, execute func(ctx context.Context, t *asynq.Task) error) {
	NewWorkerBuilder(taskType).
		WithExecute(execute).
		Register()
}
