// Package cron
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package cron

import (
	"context"
	"time"

	"devinggo/internal/dao"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/hook"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/hibiken/asynq"
)

const singletonUniqueTTL = 24 * time.Hour

type ConfigProvider struct {
	Ctx context.Context
}

func (p *ConfigProvider) GetConfigs() ([]*asynq.PeriodicTaskConfig, error) {
	bindCron := cronWorker.WorkerList
	if g.IsEmpty(bindCron) {
		return nil, nil
	}
	var dbCrons []*res.SettingCrontabOne
	err := dao.SettingCrontab.Ctx(p.Ctx).Hook(hook.Default()).Cache(orm.SetCacheOption(p.Ctx, time.Hour*24)).Where("status", 1).Scan(&dbCrons)
	if utils.IsError(err) {
		return nil, err
	}

	if g.IsEmpty(dbCrons) {
		return nil, nil
	}
	var configs []*asynq.PeriodicTaskConfig
	for _, cronItem := range bindCron {
		typeName := cronItem.GetType()
		for _, dbCron := range dbCrons {
			target := dbCron.Target
			singleton := dbCron.Singleton
			rule := dbCron.Rule
			if target == typeName {
				cronItem.SetParams(p.Ctx, dbCron.Parameter)
				cronItem.GetPayload().CrontabId = dbCron.Id
				cronItem.GetPayload().AutoTaskID = true
				cronItem.GetPayload().TaskID = typeName + "_" + gconv.String(dbCron.Id)
				if singleton == 1 {
					cronItem.GetPayload().Unique = asynq.Unique(singletonUniqueTTL)
				} else {
					cronItem.GetPayload().Unique = nil
				}
				configs = append(configs, &asynq.PeriodicTaskConfig{
					Cronspec: rule,
					Task:     cronItem.GetCronTask(),
				})
			}
		}
	}
	//glob.WithWorkLog().Debugf(context.Background(), "cron configs: %s", configs)
	return configs, nil
}
