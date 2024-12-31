// Package middleware
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package middleware

import (
	glob2 "devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hibiken/asynq"
	"time"
)

func LoggingMiddleware(h asynq.Handler) asynq.Handler {
	return asynq.HandlerFunc(func(ctx context.Context, t *asynq.Task) error {
		name := t.Type()
		start := time.Now()
		payload, err := glob2.GetPayload(ctx, t)
		crontabId := payload.CrontabId
		if err != nil {
			return err
		}
		glob2.WithWorkLog().Debugf(ctx, "Start processing [%s]", name)
		err = h.ProcessTask(ctx, t)
		if err != nil {
			glob2.WithWorkLog().Warningf(ctx, "Failure processing [%s],Error: %s", name, err)
			if !g.IsEmpty(crontabId) {
				service.SettingCrontabLog().AddLog(ctx, crontabId, 2, err.Error())
			}
			return err
		}
		if !g.IsEmpty(crontabId) {
			service.SettingCrontabLog().AddLog(ctx, crontabId, 1, "")
		}
		glob2.WithWorkLog().Debugf(ctx, "Finished processing [%s]: Elapsed Time = %v", name, time.Since(start))
		return nil
	})
}
