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
	PostController = postController{}
)

type postController struct {
	base.BaseController
}

func (c *postController) Index(ctx context.Context, in *system.IndexPostReq) (out *system.IndexPostRes, err error) {
	out = &system.IndexPostRes{}
	items, totalCount, err := service.SystemPost().GetPageList(ctx, &in.PageListReq, &in.SystemPostSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemPost, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *postController) Recycle(ctx context.Context, in *system.RecyclePostReq) (out *system.RecyclePostRes, err error) {
	out = &system.RecyclePostRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemPost().GetPageList(ctx, &in.PageListReq, &in.SystemPostSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemPost, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *postController) List(ctx context.Context, in *system.ListPostReq) (out *system.ListPostRes, err error) {
	out = &system.ListPostRes{}
	rs, err := service.SystemPost().GetList(ctx, &req.SystemPostSearch{})
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		for _, v := range rs {
			out.Data = append(out.Data, *v)
		}
	} else {
		out.Data = make([]res.SystemPost, 0)
	}

	return
}

func (c *postController) Save(ctx context.Context, in *system.SavePostReq) (out *system.SavePostRes, err error) {
	out = &system.SavePostRes{}
	id, err := service.SystemPost().Save(ctx, &in.SystemPostSave)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *postController) Read(ctx context.Context, in *system.ReadPostReq) (out *system.ReadPostRes, err error) {
	out = &system.ReadPostRes{}
	info, err := service.SystemPost().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *postController) Update(ctx context.Context, in *system.UpdatePostReq) (out *system.UpdatePostRes, err error) {
	out = &system.UpdatePostRes{}
	err = service.SystemPost().Update(ctx, &in.SystemPostSave)
	if err != nil {
		return
	}
	return
}

func (c *postController) Delete(ctx context.Context, in *system.DeletePostReq) (out *system.DeletePostRes, err error) {
	out = &system.DeletePostRes{}
	err = service.SystemPost().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *postController) RealDelete(ctx context.Context, in *system.RealDeletePostReq) (out *system.RealDeletePostRes, err error) {
	out = &system.RealDeletePostRes{}
	err = service.SystemPost().RealDelete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *postController) Recovery(ctx context.Context, in *system.RecoveryPostReq) (out *system.RecoveryPostRes, err error) {
	out = &system.RecoveryPostRes{}
	err = service.SystemPost().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *postController) ChangeStatus(ctx context.Context, in *system.ChangeStatusPostReq) (out *system.ChangeStatusPostRes, err error) {
	out = &system.ChangeStatusPostRes{}
	err = service.SystemPost().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *postController) NumberOperation(ctx context.Context, in *system.NumberOperationPostReq) (out *system.NumberOperationPostRes, err error) {
	out = &system.NumberOperationPostRes{}
	err = service.SystemPost().NumberOperation(ctx, in.Id, in.NumberName, in.NumberValue)
	if err != nil {
		return
	}
	return
}

func (c *postController) Remote(ctx context.Context, in *system.RemotePostReq) (out *system.RemotePostRes, err error) {
	out = &system.RemotePostRes{}
	r := request.GetHttpRequest(ctx)

	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.SystemPost().Model(ctx)
	var rs res.SystemPost
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
			out.Items = make([]res.SystemPost, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.SystemPost, 0)
		}
	}
	return
}
