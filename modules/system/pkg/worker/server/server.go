// Package server
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package server

import (
	"devinggo/modules/system/pkg/utils/config"
	"devinggo/modules/system/pkg/utils/conv"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/pkg/worker/middleware"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hibiken/asynq"
)

var workerServer = &workerServerManager{
	WorkerList: make(map[string](func(context.Context, *asynq.Task) error)),
}

type WorkerServerItem interface {
	// 获取任务名称
	GetType() string
	// Execute 执行任务
	Execute(ctx context.Context, t *asynq.Task) (err error)
}

type workerServerManager struct {
	WorkerList map[string](func(context.Context, *asynq.Task) error)
}

func Register(s WorkerServerItem) {
	name := s.GetType()
	if _, ok := workerServer.WorkerList[name]; ok {
		glob2.WithWorkLog().Debugf(gctx.GetInitCtx(), "workerServer.Register name:%v duplicate registration.", name)
		return
	}
	workerServer.WorkerList[name] = s.Execute
}

func GetServer(ctx context.Context) *asynq.Server {
	queueConfig := config.GetConfigMap(ctx, "worker.queues", map[string]interface{}{
		"critical": 6,
		"default":  3,
		"low":      1,
	})
	mapQueueConfig := conv.ConvIntMap(queueConfig)
	return asynq.NewServer(
		glob2.GetConnConfig(ctx),
		asynq.Config{
			Concurrency: config.GetConfigInt(ctx, "worker.concurrency", 10),
			BaseContext: func() context.Context {
				return ctx
			},
			Queues:          mapQueueConfig,
			Logger:          glob2.NewLogger(ctx),
			ShutdownTimeout: config.GetConfigDuration(ctx, "worker.shutdownTimeout", "10s"),
		},
	)
}

func Run(ctx context.Context) error {
	srv := GetServer(ctx)
	mux := asynq.NewServeMux()
	mux.Use(middleware.LoggingMiddleware)
	if !g.IsEmpty(workerServer.WorkerList) {
		for workerName, workerItem := range workerServer.WorkerList {
			mux.HandleFunc(workerName, workerItem)
		}
	}
	if err := srv.Run(mux); err != nil {
		return gerror.Wrapf(err, "start worker failed:%s", err)
	}
	return nil
}
