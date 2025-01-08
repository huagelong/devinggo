// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/worker/cron"
	"devinggo/modules/system/pkg/worker/glob"
	"devinggo/modules/system/service"
	"devinggo/modules/system/worker/consts"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	CrontabController = crontabController{}
)

type crontabController struct {
	base.BaseController
}

func (c *crontabController) Index(ctx context.Context, in *system.IndexCrontabReq) (out *system.IndexCrontabRes, err error) {
	out = &system.IndexCrontabRes{}
	items, totalCount, err := service.SettingCrontab().GetPageList(ctx, &in.PageListReq, &in.SettingCrontabSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SettingCrontab, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *crontabController) LogPageList(ctx context.Context, in *system.LogPageListReq) (out *system.LogPageListRes, err error) {
	out = &system.LogPageListRes{}
	items, totalCount, err := service.SettingCrontabLog().GetPageList(ctx, &in.PageListReq, &in.SettingCrontabLogSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SettingCrontabLog, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *crontabController) Save(ctx context.Context, in *system.SaveCrontabReq) (out *system.SaveCrontabRes, err error) {
	out = &system.SaveCrontabRes{}
	id, err := service.SettingCrontab().Save(ctx, &in.SettingCrontabSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *crontabController) Update(ctx context.Context, in *system.UpdateCrontabReq) (out *system.UpdateCrontabRes, err error) {
	out = &system.UpdateCrontabRes{}
	err = service.SettingCrontab().Update(ctx, &in.SettingCrontabUpdate)
	if err != nil {
		return
	}
	return
}

func (c *crontabController) Delete(ctx context.Context, in *system.DeleteCrontabReq) (out *system.DeleteCrontabRes, err error) {
	out = &system.DeleteCrontabRes{}
	err = service.SettingCrontab().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *crontabController) DeleteLog(ctx context.Context, in *system.DeleteCrontabLogReq) (out *system.DeleteCrontabLogRes, err error) {
	out = &system.DeleteCrontabLogRes{}
	err = service.SettingCrontabLog().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *crontabController) ChangeStatus(ctx context.Context, in *system.ChangeStatusCrontabReq) (out *system.ChangeStatusCrontabRes, err error) {
	out = &system.ChangeStatusCrontabRes{}
	err = service.SettingCrontab().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *crontabController) ReadCrontab(ctx context.Context, in *system.ReadCrontabReq) (out *system.ReadCrontabRes, err error) {
	out = &system.ReadCrontabRes{}
	data, err := service.SettingCrontab().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *data
	return
}

func (c *crontabController) RunCrontab(ctx context.Context, in *system.RunCrontabReq) (out *system.RunCrontabRes, err error) {
	out = &system.RunCrontabRes{}
	err = service.SettingCrontab().Run(ctx, in.Id)
	if err != nil {
		return
	}
	return
}

func (c *crontabController) GetCrontabTarget(ctx context.Context, in *system.GetCrontabTargetReq) (out *system.GetCrontabTargetRes, err error) {
	out = &system.GetCrontabTargetRes{}

	ls := cron.GetWorkerLists()
	if !g.IsEmpty(ls) {
		for _, item := range ls {
			typeName := item.GetType()
			if !g.IsEmpty(in.Type) && (in.Type == 1) && (typeName != consts.CMD_CRON) {
				continue
			}

			if !g.IsEmpty(in.Type) && (in.Type == 3) && (typeName != consts.URL_CRON) {
				continue
			}

			description := item.GetDescription()
			taskTmp := glob.TaskStruct{
				Name:        typeName,
				Description: description + "(" + typeName + ")",
			}
			out.Data = append(out.Data, taskTmp)
		}
	} else {
		out.Data = make([]glob.TaskStruct, 0)
	}
	return
}
