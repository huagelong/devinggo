// Package cron
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cron

import (
	"devinggo/modules/system/pkg/utils/config"
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hibiken/asynq"
	"time"
)

var cronWorker = &cronWorkerManager{
	WorkerList: make(map[string]CronTaskInterface),
}

var cronListKey = "cronWorkerList"

type cronWorkerManager struct {
	WorkerList map[string]CronTaskInterface
}

type CronTaskInterface interface {
	GetType() string //获取任务类型
	GetPayload() *glob2.Payload
	GetCronTask() *asynq.Task                          //获取任务，定时任务需要该方法
	SetParams(ctx context.Context, params *gjson.Json) //设置任务参数
	GetDescription() string                            //获取任务描述
}

func Register(s CronTaskInterface) {
	name := s.GetType()
	if _, ok := cronWorker.WorkerList[name]; ok {
		glob2.WithWorkLog().Debugf(gctx.GetInitCtx(), "cronWorker.Register name:%v duplicate registration.", name)
		return
	}
	cronWorker.WorkerList[name] = s
}

func GetWorkerList(name string) CronTaskInterface {
	if data, ok := cronWorker.WorkerList[name]; !ok {
		return nil
	} else {
		return data
	}
}

func GetWorkerLists() map[string]CronTaskInterface {
	return cronWorker.WorkerList
}

func Run(ctx context.Context) error {
	//启动先清空
	local := config.GetConfigString(ctx, "worker.location", "Asia/Shanghai")
	loc, _ := time.LoadLocation(local)
	provider := &ConfigProvider{
		Ctx: ctx,
	}

	mgr, err := asynq.NewPeriodicTaskManager(
		asynq.PeriodicTaskManagerOpts{
			RedisConnOpt: glob2.GetConnConfig(ctx),
			SchedulerOpts: &asynq.SchedulerOpts{
				Logger:   glob2.NewLogger(ctx),
				LogLevel: asynq.InfoLevel,
				Location: loc,
			},
			PeriodicTaskConfigProvider: provider,
			SyncInterval:               1 * time.Second, // this field specifies how often sync should happen
		})
	if err != nil {
		return err
	}

	if err := mgr.Run(); err != nil {
		return err
	}
	return nil
}
