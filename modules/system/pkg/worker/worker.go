// Package worker
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package worker

import (
	"context"

	"devinggo/modules/system/pkg/worker/cron"
	"devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/server"
	"devinggo/modules/system/pkg/worker/task"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/hibiken/asynq"
)

var (
	// defaultManager 默认的全局Worker管理器
	defaultManager *Manager
)

func init() {
	// 创建默认的全局Manager
	defaultManager = New(gctx.GetInitCtx())
}

// Manager 统一的Worker管理器
type Manager struct {
	ctx context.Context
}

// New 创建Worker管理器
func New(ctx context.Context) *Manager {
	return &Manager{ctx: ctx}
}

// RegisterCronTask 注册定时任务
func (m *Manager) RegisterCronTask(task CronTask) *Manager {
	cron.Register(task)
	return m
}

// RegisterWorker 注册Worker处理器
func (m *Manager) RegisterWorker(worker Worker) *Manager {
	server.Register(worker)
	return m
}

// Register 注册Worker到全局Manager（供各Worker文件的init使用）
func Register(worker Worker) {
	defaultManager.RegisterWorker(worker)
}

// RegisterCron 注册Cron任务到全局Manager（供各Cron文件的init使用）
func RegisterCron(task CronTask) {
	defaultManager.RegisterCronTask(task)
}

// GetDefaultManager 获取默认的全局Manager
func GetDefaultManager() *Manager {
	return defaultManager
}

// RunCron 启动定时任务调度器
func (m *Manager) RunCron() error {
	return cron.Run(m.ctx)
}

// RunServer 启动Worker服务器
func (m *Manager) RunServer() error {
	return server.Run(m.ctx)
}

// SendTask 发送任务
func (m *Manager) SendTask(taskItem Task) error {
	return task.NewTask(m.ctx, taskItem)
}

// SendSimpleTask 发送简单任务（无自定义Send方法）
func (m *Manager) SendSimpleTask(taskItem SimpleTask) error {
	return task.NewSimpleTask(m.ctx, taskItem)
}

// GetClient 获取Asynq客户端
func (m *Manager) GetClient() *asynq.Client {
	return task.GetClient(m.ctx)
}

// GetServer 获取Asynq服务器实例
func (m *Manager) GetServer() *asynq.Server {
	return server.GetServer(m.ctx)
}

// CronTask 定时任务接口
type CronTask interface {
	GetType() string
	GetPayload() *glob.Payload
	GetCronTask() *asynq.Task
	SetParams(ctx context.Context, params *gjson.Json)
	GetDescription() string
}

// Worker 任务处理器接口
type Worker interface {
	GetType() string
	Execute(ctx context.Context, t *asynq.Task) error
}

// Task 任务发送接口
type Task interface {
	GetType() string
	GetPayload() *glob.Payload
	Send(ctx context.Context, data interface{}) error
}

// SimpleTask 简单任务接口（用于快速发送任务）
type SimpleTask interface {
	GetType() string
	GetPayload() *glob.Payload
}

// GetPayload 从任务中获取Payload
func GetPayload(ctx context.Context, t *asynq.Task) (*glob.Payload, error) {
	return glob.GetPayload(ctx, t)
}

// GetParameters 从任务中获取参数并转换为指定类型
func GetParameters[T any](ctx context.Context, t *asynq.Task) (T, error) {
	return glob.GetParamters[T](ctx, t)
}

// GetLogger 获取Worker日志记录器
func GetLogger() *glog.Logger {
	return glob.WithWorkLog()
}

// GetLoggerWithContext 获取带上下文的Worker日志记录器（已弃用，请直接使用GetLogger()）
// Deprecated: 使用 GetLogger() 代替
func GetLoggerWithContext(ctx context.Context) *glog.Logger {
	return glob.WithWorkLog()
}
