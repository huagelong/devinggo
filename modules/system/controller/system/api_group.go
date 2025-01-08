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
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	ApiGroupController = apiGroupController{}
)

type apiGroupController struct {
	base.BaseController
}

func (c *apiGroupController) Index(ctx context.Context, in *system.IndexApiGroupReq) (out *system.IndexApiGroupRes, err error) {
	out = &system.IndexApiGroupRes{}
	items, totalCount, err := service.SystemApiGroup().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemApiGroupSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApiGroup, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *apiGroupController) List(ctx context.Context, in *system.ListApiGroupReq) (out *system.ListApiGroupRes, err error) {
	out = &system.ListApiGroupRes{}
	search := &in.SystemApiGroupSearch
	search.Status = 1
	rs, err := service.SystemApiGroup().GetList(ctx, search)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemApiGroup, 0)
	}

	return
}

func (c *apiGroupController) Recycle(ctx context.Context, in *system.RecycleApiGroupReq) (out *system.RecycleApiGroupRes, err error) {
	out = &system.RecycleApiGroupRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemApiGroup().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemApiGroupSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApiGroup, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *apiGroupController) Save(ctx context.Context, in *system.SaveApiGroupReq) (out *system.SaveApiGroupRes, err error) {
	out = &system.SaveApiGroupRes{}
	id, err := service.SystemApiGroup().Save(ctx, &in.SystemApiGroupSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *apiGroupController) Read(ctx context.Context, in *system.ReadApiGroupReq) (out *system.ReadApiGroupRes, err error) {
	out = &system.ReadApiGroupRes{}
	info, err := service.SystemApiGroup().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *apiGroupController) Update(ctx context.Context, in *system.UpdateApiGroupReq) (out *system.UpdateApiGroupRes, err error) {
	err = dao.SystemApiGroup.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.UpdateApiGroupRes{}
		err = service.SystemApiGroup().Update(ctx, &in.SystemApiGroupUpdate)
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

func (c *apiGroupController) Delete(ctx context.Context, in *system.DeleteApiGroupReq) (out *system.DeleteApiGroupRes, err error) {
	out = &system.DeleteApiGroupRes{}
	err = service.SystemApiGroup().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *apiGroupController) RealDelete(ctx context.Context, in *system.RealDeleteApiGroupReq) (out *system.RealDeleteApiGroupRes, err error) {
	err = dao.SystemApiGroup.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.RealDeleteApiGroupRes{}
		err = service.SystemApiGroup().RealDelete(ctx, in.Ids)
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

func (c *apiGroupController) Recovery(ctx context.Context, in *system.RecoveryApiGroupReq) (out *system.RecoveryApiGroupRes, err error) {
	out = &system.RecoveryApiGroupRes{}
	err = service.SystemApiGroup().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *apiGroupController) ChangeStatus(ctx context.Context, in *system.ChangeStatusApiGroupReq) (out *system.ChangeStatusApiGroupRes, err error) {
	out = &system.ChangeStatusApiGroupRes{}
	err = service.SystemApiGroup().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}
