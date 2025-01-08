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
	"devinggo/modules/system/myerror"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	MenuController = menuController{}
)

type menuController struct {
	base.BaseController
}

func (c *menuController) Index(ctx context.Context, in *system.IndexMenuReq) (out *system.IndexMenuRes, err error) {
	out = &system.IndexMenuRes{}
	items, err := service.SystemMenu().GetTreeList(ctx, &in.SystemMenuSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemMenuTree, 0)
	}
	return
}

func (c *menuController) Recycle(ctx context.Context, in *system.RecycleMenuReq) (out *system.RecycleMenuRes, err error) {
	out = &system.RecycleMenuRes{}
	in.Recycle = true
	items, err := service.SystemMenu().GetRecycleTreeList(ctx, &in.SystemMenuSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemMenuTree, 0)
	}
	return
}

func (c *menuController) Tree(ctx context.Context, in *system.TreeMenuReq) (out *system.TreeMenuRes, err error) {
	out = &system.TreeMenuRes{}
	items, err := service.SystemMenu().GetSelectTree(ctx, c.UserId, in.OnlyMenu, in.Scope)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemDeptSelectTree, 0)
	}
	return
}

func (c *menuController) Save(ctx context.Context, in *system.SaveMenuReq) (out *system.SaveMenuRes, err error) {
	out = &system.SaveMenuRes{}
	id, err := service.SystemMenu().Save(ctx, &in.SystemMenuSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *menuController) Update(ctx context.Context, in *system.UpdateMenuReq) (out *system.UpdateMenuRes, err error) {
	out = &system.UpdateMenuRes{}
	err = service.SystemMenu().Update(ctx, &in.SystemMenuSave)
	if err != nil {
		return
	}
	return
}

func (c *menuController) Delete(ctx context.Context, in *system.DeleteMenuReq) (out *system.DeleteMenuRes, err error) {
	out = &system.DeleteMenuRes{}
	names, err := service.SystemMenu().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	if !g.IsEmpty(names) {
		return out, myerror.ValidationFailed(ctx, "菜单：【"+slice.Join(names, ",")+"】有子菜单，不能删除")
	}
	return
}

func (c *menuController) RealDelete(ctx context.Context, in *system.RealDeleteMenuReq) (out *system.RealDeleteMenuRes, err error) {
	out = &system.RealDeleteMenuRes{}
	names, err := service.SystemMenu().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}

	if !g.IsEmpty(names) {
		return out, myerror.ValidationFailed(ctx, "菜单：【"+slice.Join(names, ",")+"】有子菜单，不能删除")
	}
	return
}

func (c *menuController) Recovery(ctx context.Context, in *system.RecoveryMenuReq) (out *system.RecoveryMenuRes, err error) {
	out = &system.RecoveryMenuRes{}
	err = service.SystemMenu().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *menuController) ChangeStatus(ctx context.Context, in *system.ChangeStatusMenuReq) (out *system.ChangeStatusMenuRes, err error) {
	out = &system.ChangeStatusMenuRes{}
	err = service.SystemMenu().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *menuController) NumberOperation(ctx context.Context, in *system.NumberOperationMenuReq) (out *system.NumberOperationMenuRes, err error) {
	out = &system.NumberOperationMenuRes{}
	err = service.SystemMenu().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}
