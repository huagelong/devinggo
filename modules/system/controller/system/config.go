// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"

	"devinggo/internal/dao"
	"devinggo/internal/model/do"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/utils"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"

	"github.com/gogf/gf/v2/frame/g"
)

var (
	ConfigController = configController{}
)

type configController struct {
	base.BaseController
}

func (c *configController) IndexConfigGroup(ctx context.Context, in *system.IndexConfigGroupReq) (out *system.IndexConfigGroupRes, err error) {
	out = &system.IndexConfigGroupRes{}
	rs, err := service.SettingConfigGroup().GetList(ctx)
	if utils.IsError(err) {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SettingConfigGroup, 0)
	}
	return
}

func (c *configController) SaveConfigGroup(ctx context.Context, in *system.SaveConfigGroupReq) (out *system.SaveConfigGroupRes, err error) {
	out = &system.SaveConfigGroupRes{}
	id, err := service.SettingConfigGroup().SaveConfigGroup(ctx, &in.SettingConfigGroupSave)
	if utils.IsError(err) {
		return
	}
	out.Id = id
	return
}

func (c *configController) UpdateConfigGroup(ctx context.Context, in *system.UpdateConfigGroupReq) (out *system.UpdateConfigGroupRes, err error) {
	out = &system.UpdateConfigGroupRes{}
	err = service.SettingConfigGroup().UpdateConfigGroup(ctx, &in.SettingConfigGroupUpdate)
	if utils.IsError(err) {
		return
	}
	return
}

func (c *configController) DeleteConfigGroup(ctx context.Context, in *system.DeleteConfigGroupReq) (out *system.DeleteConfigGroupRes, err error) {
	out = &system.DeleteConfigGroupRes{}
	err = service.SettingConfigGroup().DeleteConfigGroup(ctx, in.Id)
	if utils.IsError(err) {
		return
	}
	return
}

func (c *configController) IndexConfig(ctx context.Context, in *system.IndexConfigReq) (out *system.IndexConfigRes, err error) {
	out = &system.IndexConfigRes{}
	rs, err := service.SettingConfig().GetList(ctx, &in.SettingConfigSearch)
	if utils.IsError(err) {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SettingConfig, 0)
	}
	return
}

func (c *configController) SaveConfig(ctx context.Context, in *system.SaveConfigReq) (out *system.SaveConfigGroupRes, err error) {
	out = &system.SaveConfigGroupRes{}
	id, err := service.SettingConfig().SaveConfig(ctx, &in.SettingConfigSave)
	if utils.IsError(err) {
		return
	}
	out.Id = id
	return
}

func (c *configController) UpdateConfig(ctx context.Context, in *system.UpdateConfigReq) (out *system.UpdateConfigRes, err error) {
	out = &system.UpdateConfigRes{}
	err = service.SettingConfig().UpdateConfig(ctx, &in.SettingConfigUpdate)
	if utils.IsError(err) {
		return
	}
	return
}

func (c *configController) UpdateByKeysConfig(ctx context.Context, in *system.UpdateByKeysConfigReq) (out *system.UpdateByKeysConfigRes, err error) {
	r := request.GetHttpRequest(ctx)
	j, err := r.GetJson()
	if utils.IsError(err) {
		return
	}
	g.Log().Info(ctx, "j", j)
	maps := j.Map()
	for k, v := range maps {
		saveData := do.SettingConfig{
			Value: v,
		}
		_, err = service.SettingConfig().Model(ctx).Where(dao.SettingConfig.Columns().Key, k).Data(saveData).Update()
		if utils.IsError(err) {
			return
		}
	}

	out = &system.UpdateByKeysConfigRes{}
	return
}

func (c *configController) DeleteConfig(ctx context.Context, in *system.DeleteConfigReq) (out *system.DeleteConfigRes, err error) {
	out = &system.DeleteConfigRes{}
	err = service.SettingConfig().DeleteConfig(ctx, in.Ids)
	if utils.IsError(err) {
		return
	}
	return
}
