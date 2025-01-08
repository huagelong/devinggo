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
	ApiController = apiController{}
)

type apiController struct {
	base.BaseController
}

func (c *apiController) Index(ctx context.Context, in *system.IndexApiReq) (out *system.IndexApiRes, err error) {
	out = &system.IndexApiRes{}
	items, totalCount, err := service.SystemApi().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemApiSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApi, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *apiController) Recycle(ctx context.Context, in *system.RecycleApiReq) (out *system.RecycleApiRes, err error) {
	out = &system.RecycleApiRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemApi().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemApiSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApi, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *apiController) Save(ctx context.Context, in *system.SaveApiReq) (out *system.SaveApiRes, err error) {
	out = &system.SaveApiRes{}
	id, err := service.SystemApi().Save(ctx, &in.SystemApiSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *apiController) Read(ctx context.Context, in *system.ReadApiReq) (out *system.ReadApiRes, err error) {
	out = &system.ReadApiRes{}
	info, err := service.SystemApi().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *apiController) Update(ctx context.Context, in *system.UpdateApiReq) (out *system.UpdateApiRes, err error) {
	err = dao.SystemApi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.UpdateApiRes{}
		err = service.SystemApi().Update(ctx, &in.SystemApiUpdate)
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

func (c *apiController) Delete(ctx context.Context, in *system.DeleteApiReq) (out *system.DeleteApiRes, err error) {
	out = &system.DeleteApiRes{}
	err = service.SystemApi().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *apiController) RealDelete(ctx context.Context, in *system.RealDeleteApiReq) (out *system.RealDeleteApiRes, err error) {
	err = dao.SystemApi.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.RealDeleteApiRes{}
		err = service.SystemApi().RealDelete(ctx, in.Ids)
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

func (c *apiController) Recovery(ctx context.Context, in *system.RecoveryApiReq) (out *system.RecoveryApiRes, err error) {
	out = &system.RecoveryApiRes{}
	err = service.SystemApi().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *apiController) ChangeStatus(ctx context.Context, in *system.ChangeStatusApiReq) (out *system.ChangeStatusApiRes, err error) {
	out = &system.ChangeStatusApiRes{}
	err = service.SystemApi().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}
