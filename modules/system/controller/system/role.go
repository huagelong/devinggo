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
	RoleController = roleController{}
)

type roleController struct {
	base.BaseController
}

func (c *roleController) Index(ctx context.Context, in *system.IndexRoleReq) (out *system.IndexRoleRes, err error) {
	out = &system.IndexRoleRes{}
	items, totalCount, err := service.SystemRole().GetPageList(ctx, &in.PageListReq, &in.SystemRoleSearch, false)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemRole, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *roleController) Recycle(ctx context.Context, in *system.RecycleRoleReq) (out *system.RecycleRoleRes, err error) {
	out = &system.RecycleRoleRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemRole().GetPageList(ctx, &in.PageListReq, &in.SystemRoleSearch, false)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemRole, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *roleController) List(ctx context.Context, in *system.ListRoleReq) (out *system.ListRoleRes, err error) {
	out = &system.ListRoleRes{}
	rs, err := service.SystemRole().GetList(ctx, &req.SystemRoleSearch{}, true)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemRole, 0)
	}

	return
}

func (c *roleController) Save(ctx context.Context, in *system.SaveRoleReq) (out *system.SaveRoleRes, err error) {
	out = &system.SaveRoleRes{}
	id, err := service.SystemRole().Save(ctx, &in.SystemRoleSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *roleController) Update(ctx context.Context, in *system.UpdateRoleReq) (out *system.UpdateRoleRes, err error) {
	out = &system.UpdateRoleRes{}
	err = service.SystemRole().Update(ctx, &in.SystemRoleSave)
	if err != nil {
		return
	}
	return
}

func (c *roleController) Delete(ctx context.Context, in *system.DeleteRoleReq) (out *system.DeleteRoleRes, err error) {
	out = &system.DeleteRoleRes{}
	err = service.SystemRole().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *roleController) RealDelete(ctx context.Context, in *system.RealDeleteRoleReq) (out *system.RealDeleteRoleRes, err error) {
	out = &system.RealDeleteRoleRes{}
	err = service.SystemRole().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *roleController) Recovery(ctx context.Context, in *system.RecoveryRoleReq) (out *system.RecoveryRoleRes, err error) {
	out = &system.RecoveryRoleRes{}
	err = service.SystemRole().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *roleController) ChangeStatus(ctx context.Context, in *system.ChangeStatusRoleReq) (out *system.ChangeStatusRoleRes, err error) {
	out = &system.ChangeStatusRoleRes{}
	err = service.SystemRole().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *roleController) NumberOperation(ctx context.Context, in *system.NumberOperationRoleReq) (out *system.NumberOperationRoleRes, err error) {
	out = &system.NumberOperationRoleRes{}
	err = service.SystemRole().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *roleController) MenuPermissionRole(ctx context.Context, in *system.MenuPermissionRoleReq) (out *system.MenuPermissionRoleRes, err error) {
	out = &system.MenuPermissionRoleRes{}
	err = service.SystemRole().Update(ctx, &in.SystemRoleSave)
	if err != nil {
		return
	}
	return
}

func (c *roleController) DataPermissionRole(ctx context.Context, in *system.DataPermissionRoleReq) (out *system.DataPermissionRoleRes, err error) {
	out = &system.DataPermissionRoleRes{}
	err = service.SystemRole().Update(ctx, &in.SystemRoleSave)
	if err != nil {
		return
	}
	return
}

func (c *roleController) GetMenuByRole(ctx context.Context, in *system.GetMenuByRoleReq) (out *system.GetMenuByRoleRes, err error) {
	out = &system.GetMenuByRoleRes{}
	ids := make([]int64, 0)
	ids = append(ids, in.Id)
	data, err := service.SystemRole().GetMenuByRoleIds(ctx, ids)
	if err != nil {
		return
	}
	out.Data = data
	return
}

func (c *roleController) GetDeptByRole(ctx context.Context, in *system.GetDeptByRoleReq) (out *system.GetDeptByRoleRes, err error) {
	out = &system.GetDeptByRoleRes{}
	ids := make([]int64, 0)
	ids = append(ids, in.Id)
	data, err := service.SystemRole().GetDeptByRole(ctx, ids)
	if err != nil {
		return
	}
	out.Data = data
	return
}

func (c *roleController) Remote(ctx context.Context, in *system.RemoteRoleReq) (out *system.RemoteRoleRes, err error) {
	out = &system.RemoteRoleRes{}
	r := request.GetHttpRequest(ctx)

	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.SystemRole().Model(ctx)
	var rs res.SystemRole
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
			out.Items = make([]res.SystemRole, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.SystemRole, 0)
		}
	}
	return
}
