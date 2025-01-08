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
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/pkg/utils/slice"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	DeptController = deptController{}
)

type deptController struct {
	base.BaseController
}

func (c *deptController) Index(ctx context.Context, in *system.IndexReq) (out *system.IndexRes, err error) {
	out = &system.IndexRes{}
	items, err := service.SystemDept().GetListTreeList(ctx, &in.SystemDeptSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemListDeptTree, 0)
	}
	return
}

func (c *deptController) Recycle(ctx context.Context, in *system.RecycleReq) (out *system.RecycleRes, err error) {
	out = &system.RecycleRes{}
	in.Recycle = true
	items, err := service.SystemDept().GetRecycleTreeList(ctx, &in.SystemDeptSearch)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemListDeptTree, 0)
	}
	return
}

func (c *deptController) Tree(ctx context.Context, in *system.TreeReq) (out *system.TreeRes, err error) {
	out = &system.TreeRes{}
	items, err := service.SystemDept().GetSelectTree(ctx, c.UserId)
	if err != nil {
		return
	}
	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Data = append(out.Data, *item)
		}
	} else {
		out.Data = make([]res.SystemDeptTree, 0)
	}
	return
}

func (c *deptController) GetLeaderList(ctx context.Context, in *system.GetLeaderListReq) (out *system.GetLeaderListRes, err error) {
	out = &system.GetLeaderListRes{}
	items, totalCount, err := service.SystemDeptLeader().GetPageList(ctx, &in.PageListReq, &in.SystemDeptLeaderSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemDeptLeaderInfo, 0)
	}
	out.PageRes.Pack(in, totalCount)

	return
}

func (c *deptController) Save(ctx context.Context, in *system.SaveReq) (out *system.SaveRes, err error) {
	out = &system.SaveRes{}
	id, err := service.SystemDept().Save(ctx, &in.SystemDeptSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *deptController) AddLeader(ctx context.Context, in *system.AddLeaderReq) (out *system.AddLeaderRes, err error) {
	out = &system.AddLeaderRes{}
	err = service.SystemDept().AddLeader(ctx, &in.SystemDeptAddLeader)
	if err != nil {
		return
	}
	return
}

func (c *deptController) DelLeader(ctx context.Context, in *system.DelLeaderReq) (out *system.DelLeaderRes, err error) {
	out = &system.DelLeaderRes{}
	err = service.SystemDept().DelLeader(ctx, in.Id, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *deptController) Update(ctx context.Context, in *system.UpdateReq) (out *system.UpdateRes, err error) {
	out = &system.UpdateRes{}
	err = service.SystemDept().Update(ctx, &in.SystemDeptSave)
	if err != nil {
		return
	}
	return
}

func (c *deptController) Delete(ctx context.Context, in *system.DeleteReq) (out *system.DeleteRes, err error) {
	out = &system.DeleteRes{}
	names, err := service.SystemDept().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	if !g.IsEmpty(names) {
		return out, myerror.ValidationFailed(ctx, "部门：【"+slice.Join(names, ",")+"】有子部门，不能删除")
	}
	return
}

func (c *deptController) RealDelete(ctx context.Context, in *system.RealDeleteReq) (out *system.RealDeleteRes, err error) {
	out = &system.RealDeleteRes{}
	names, err := service.SystemDept().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	if !g.IsEmpty(names) {
		return out, myerror.ValidationFailed(ctx, "部门：【"+slice.Join(names, ",")+"】有子部门，不能删除")
	}
	return
}

func (c *deptController) Recovery(ctx context.Context, in *system.RecoveryReq) (out *system.RecoveryRes, err error) {
	out = &system.RecoveryRes{}
	err = service.SystemDept().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *deptController) ChangeStatus(ctx context.Context, in *system.ChangeStatusReq) (out *system.ChangeStatusRes, err error) {
	out = &system.ChangeStatusRes{}
	err = service.SystemDept().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *deptController) NumberOperation(ctx context.Context, in *system.NumberOperationReq) (out *system.NumberOperationRes, err error) {
	out = &system.NumberOperationRes{}
	err = service.SystemDept().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *deptController) Remote(ctx context.Context, in *system.RemoteDeptReq) (out *system.RemoteDeptRes, err error) {
	out = &system.RemoteDeptRes{}
	r := request.GetHttpRequest(ctx)

	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.SystemDept().Model(ctx)
	var rs res.SystemDept
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
			out.Items = make([]res.SystemDept, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.SystemDept, 0)
		}
	}
	return
}
