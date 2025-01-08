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
	"devinggo/modules/system/model/req"
	"devinggo/modules/system/model/res"
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	SystemModulesController = systemModulesController{}
)

type systemModulesController struct {
	base.BaseController
}

func (c *systemModulesController) Index(ctx context.Context, in *system.IndexSystemModulesReq) (out *system.IndexSystemModulesRes, err error) {
	out = &system.IndexSystemModulesRes{}
	items, totalCount, err := service.SystemModules().GetPageList(ctx, &in.PageListReq, &in.SystemModulesSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemModules, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *systemModulesController) Recycle(ctx context.Context, in *system.RecycleSystemModulesReq) (out *system.RecycleSystemModulesRes, err error) {
	out = &system.RecycleSystemModulesRes{}
	pageListReq := &in.PageListReq
	pageListReq.Recycle = true
	items, totalCount, err := service.SystemModules().GetPageList(ctx, pageListReq, &in.SystemModulesSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemModules, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *systemModulesController) List(ctx context.Context, in *system.ListSystemModulesReq) (out *system.ListSystemModulesRes, err error) {
	out = &system.ListSystemModulesRes{}
	rs, err := service.SystemModules().GetList(ctx, &in.ListReq, &req.SystemModulesSearch{})
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemModules, 0)
	}
	return
}

func (c *systemModulesController) Save(ctx context.Context, in *system.SaveSystemModulesReq) (out *system.SaveSystemModulesRes, err error) {
	out = &system.SaveSystemModulesRes{}
	id, err := service.SystemModules().Save(ctx, &in.SystemModulesSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *systemModulesController) Read(ctx context.Context, in *system.ReadSystemModulesReq) (out *system.ReadSystemModulesRes, err error) {
	out = &system.ReadSystemModulesRes{}
	info, err := service.SystemModules().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *systemModulesController) Update(ctx context.Context, in *system.UpdateSystemModulesReq) (out *system.UpdateSystemModulesRes, err error) {
	out = &system.UpdateSystemModulesRes{}
	err = service.SystemModules().Update(ctx, &in.SystemModulesUpdate)
	if err != nil {
		return
	}
	return
}

func (c *systemModulesController) Delete(ctx context.Context, in *system.DeleteSystemModulesReq) (out *system.DeleteSystemModulesRes, err error) {
	out = &system.DeleteSystemModulesRes{}
	err = service.SystemModules().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
func (c *systemModulesController) RealDelete(ctx context.Context, in *system.RealDeleteSystemModulesReq) (out *system.RealDeleteSystemModulesRes, err error) {
	out = &system.RealDeleteSystemModulesRes{}
	err = service.SystemModules().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *systemModulesController) Recovery(ctx context.Context, in *system.RecoverySystemModulesReq) (out *system.RecoverySystemModulesRes, err error) {
	out = &system.RecoverySystemModulesRes{}
	err = service.SystemModules().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *systemModulesController) ChangeStatus(ctx context.Context, in *system.ChangeStatusSystemModulesReq) (out *system.ChangeStatusSystemModulesRes, err error) {
	out = &system.ChangeStatusSystemModulesRes{}
	err = service.SystemModules().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *systemModulesController) Remote(ctx context.Context, in *system.RemoteSystemModulesReq) (out *system.RemoteSystemModulesRes, err error) {
	out = &system.RemoteSystemModulesRes{}
	r := request.GetHttpRequest(ctx)
	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.SystemModules().Model(ctx)
	var rs res.SystemModules
	remote := orm.NewRemote(m, rs)
	openPage := params.GetVar("openPage")
	items, totalCount, err := remote.GetRemote(ctx, params)
	if err != nil {
		return
	}
	if !g.IsEmpty(openPage) && openPage.Bool() {
		if !g.IsEmpty(items) {
			for _, item := range items {
				out.Items = append(out.Items, item)
			}
		} else {
			out.Items = make([]res.SystemModules, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.SystemModules, 0)
		}
	}
	return
}
