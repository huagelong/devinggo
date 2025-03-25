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
	"devinggo/modules/system/pkg/orm"
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
	"github.com/gogf/gf/v2/container/gmap"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	AppController = appController{}
)

type appController struct {
	base.BaseController
}

func (c *appController) GetAppId(ctx context.Context, in *system.GetAppIdReq) (out *system.GetAppIdRes, err error) {
	out = &system.GetAppIdRes{}
	rs, err := service.SystemApp().GetAppId(ctx)
	if err != nil {
		return
	}
	out.AppId = rs
	return
}

func (c *appController) GetAppSecret(ctx context.Context, in *system.GetAppSecretReq) (out *system.GetAppSecretRes, err error) {
	out = &system.GetAppSecretRes{}
	rs, err := service.SystemApp().GetAppSecret(ctx)
	if err != nil {
		return
	}
	out.AppSecret = rs
	return
}

func (c *appController) Index(ctx context.Context, in *system.IndexAppReq) (out *system.IndexAppRes, err error) {
	out = &system.IndexAppRes{}
	items, totalCount, err := service.SystemApp().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemAppSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApp, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *appController) GetApiList(ctx context.Context, in *system.GetApiListReq) (out *system.GetApiListRes, err error) {
	out = &system.GetApiListRes{}
	rs, err := service.SystemApp().GetApiList(ctx, in.Id)
	if err != nil {
		return
	}

	if !g.IsEmpty(rs) {
		out.Data = rs
	} else {
		out.Data = make([]int64, 0)
	}

	return
}

func (c *appController) Recycle(ctx context.Context, in *system.RecycleAppReq) (out *system.RecycleAppRes, err error) {
	out = &system.RecycleAppRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemApp().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemAppSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApp, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *appController) Save(ctx context.Context, in *system.SaveAppReq) (out *system.SaveAppRes, err error) {
	out = &system.SaveAppRes{}
	id, err := service.SystemApp().Save(ctx, &in.SystemAppSave, c.UserId)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *appController) Read(ctx context.Context, in *system.ReadAppReq) (out *system.ReadAppRes, err error) {
	out = &system.ReadAppRes{}
	info, err := service.SystemApp().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *appController) Update(ctx context.Context, in *system.UpdateAppReq) (out *system.UpdateAppRes, err error) {
	out = &system.UpdateAppRes{}
	err = dao.SystemApp.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = service.SystemApp().Update(ctx, &in.SystemAppUpdate)
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

func (c *appController) BindApp(ctx context.Context, in *system.BindAppReq) (out *system.BindAppRes, err error) {
	err = dao.SystemApp.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		err = service.SystemApp().BindApp(ctx, in.Id, in.ApiIds)
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

func (c *appController) Delete(ctx context.Context, in *system.DeleteAppReq) (out *system.DeleteAppRes, err error) {
	out = &system.DeleteAppRes{}
	err = service.SystemApp().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *appController) RealDelete(ctx context.Context, in *system.RealDeleteAppReq) (out *system.RealDeleteAppRes, err error) {
	err = dao.SystemApp.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.RealDeleteAppRes{}
		err = service.SystemApp().RealDelete(ctx, in.Ids)
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

func (c *appController) Recovery(ctx context.Context, in *system.RecoveryAppReq) (out *system.RecoveryAppRes, err error) {
	out = &system.RecoveryAppRes{}
	err = service.SystemApp().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *appController) ChangeStatus(ctx context.Context, in *system.ChangeStatusAppReq) (out *system.ChangeStatusAppRes, err error) {
	out = &system.ChangeStatusAppRes{}
	err = service.SystemApp().ChangeStatus(ctx, in.Id, in.Status)
	if err != nil {
		return
	}
	return
}

func (c *appController) Remote(ctx context.Context, in *system.RemoteAppReq) (out *system.RemoteAppRes, err error) {
	out = &system.RemoteAppRes{}
	r := request.GetHttpRequest(ctx)

	params := gmap.NewStrAnyMapFrom(r.GetMap())
	m := service.SystemApp().Model(ctx)
	var rs res.SystemApp
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
			out.Items = make([]res.SystemApp, 0)
		}
		out.PageRes.Pack(in, totalCount)
	} else {
		if !g.IsEmpty(items) {
			out.Data = items
		} else {
			out.Data = make([]res.SystemApp, 0)
		}
	}
	return
}
