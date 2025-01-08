// Package system
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package system

import (
	"context"
	"devinggo/internal/dao"
	"devinggo/modules/system/api/system"
	"devinggo/modules/system/controller/base"
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	AppGroupController = appGroupController{}
)

type appGroupController struct {
	base.BaseController
}

func (c *appGroupController) Index(ctx context.Context, in *system.IndexAppGroupReq) (out *system.IndexAppGroupRes, err error) {
	out = &system.IndexAppGroupRes{}
	items, totalCount, err := service.SystemAppGroup().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemAppGroupSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemAppGroup, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *appGroupController) List(ctx context.Context, in *system.ListAppGroupReq) (out *system.ListAppGroupRes, err error) {
	out = &system.ListAppGroupRes{}
	rs, err := service.SystemAppGroup().GetList(ctx, &req.SystemAppGroupSearch{Status: 1})
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemAppGroup, 0)
	}

	return
}

func (c *appGroupController) Recycle(ctx context.Context, in *system.RecycleAppGroupReq) (out *system.RecycleAppGroupRes, err error) {
	out = &system.RecycleAppGroupRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemAppGroup().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemAppGroupSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemAppGroup, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *appGroupController) Save(ctx context.Context, in *system.SaveAppGroupReq) (out *system.SaveAppGroupRes, err error) {
	out = &system.SaveAppGroupRes{}
	id, err := service.SystemAppGroup().Save(ctx, &in.SystemAppGroupSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *appGroupController) Read(ctx context.Context, in *system.ReadAppGroupReq) (out *system.ReadAppGroupRes, err error) {
	out = &system.ReadAppGroupRes{}
	info, err := service.SystemAppGroup().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *appGroupController) Update(ctx context.Context, in *system.UpdateAppGroupReq) (out *system.UpdateAppGroupRes, err error) {
	err = dao.SystemAppGroup.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.UpdateAppGroupRes{}
		err = service.SystemAppGroup().Update(ctx, &in.SystemAppGroupUpdate)
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return
	}
	return
}

func (c *appGroupController) Delete(ctx context.Context, in *system.DeleteAppGroupReq) (out *system.DeleteAppGroupRes, err error) {
	out = &system.DeleteAppGroupRes{}
	err = service.SystemAppGroup().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *appGroupController) RealDelete(ctx context.Context, in *system.RealDeleteAppGroupReq) (out *system.RealDeleteAppGroupRes, err error) {
	err = dao.SystemAppGroup.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.RealDeleteAppGroupRes{}
		err = service.SystemAppGroup().RealDelete(ctx, in.Ids)
		if err != nil {
			return
		}
		return
	})
	if err != nil {
		return
	}
	return
}

func (c *appGroupController) Recovery(ctx context.Context, in *system.RecoveryAppGroupReq) (out *system.RecoveryAppGroupRes, err error) {
	out = &system.RecoveryAppGroupRes{}
	err = service.SystemAppGroup().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *appGroupController) ChangeStatus(ctx context.Context, in *system.ChangeStatusAppGroupReq) (out *system.ChangeStatusAppGroupRes, err error) {
	out = &system.ChangeStatusAppGroupRes{}
	err = service.SystemAppGroup().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}
