// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	"context"

	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/hibiken/asynq"
)

func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		name := t.Type()
		startTime := gtime.Now()
		ctx = glob2.WithTaskOutput(ctx)
		payload, err := glob2.GetPayload(ctx, t)
		if err != nil {
			return err
		}
		crontabId := payload.CrontabId
		glob2.WithWorkLog().Debugf(ctx, "Start processing [%s]", name)
		err = h.ProcessTask(ctx, t)
		endTime := gtime.Now()
		output := glob2.GetTaskOutput(ctx)
		if err != nil {
			glob2.WithWorkLog().Warningf(ctx, "Failure processing [%s],Error: %s", name, err)
			if !g.IsEmpty(crontabId) {
				_ = service.SettingCrontabLog().AddLog(ctx, crontabId, 2, err.Error(), startTime, endTime, output)
			}
			return err
		}
		if !g.IsEmpty(crontabId) {
			_ = service.SettingCrontabLog().AddLog(ctx, crontabId, 1, "", startTime, endTime, output)
		}
		glob2.WithWorkLog().Debugf(ctx, "Finished processing [%s]: Elapsed Time = %v", name, endTime.Time.Sub(startTime.Time))
		return nil
	})
}
