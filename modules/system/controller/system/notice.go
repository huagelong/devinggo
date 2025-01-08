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
	NoticeController = noticeController{}
)

type noticeController struct {
	base.BaseController
}

func (c *noticeController) Index(ctx context.Context, in *system.IndexNoticeReq) (out *system.IndexNoticeRes, err error) {
	out = &system.IndexNoticeRes{}
	items, totalCount, err := service.SystemNotice().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemNoticeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemNotice, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *noticeController) Recycle(ctx context.Context, in *system.RecycleNoticeReq) (out *system.RecycleNoticeRes, err error) {
	out = &system.RecycleNoticeRes{}
	in.Recycle = true
	items, totalCount, err := service.SystemNotice().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemNoticeSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemNotice, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *noticeController) Save(ctx context.Context, in *system.SaveNoticeReq) (out *system.SaveNoticeRes, err error) {
	out = &system.SaveNoticeRes{}
	id, err := service.SystemNotice().Save(ctx, &in.SystemNoticeSave, c.UserId)
	if err != nil {
		return
	}
	out.Id = id
	return
}

func (c *noticeController) Read(ctx context.Context, in *system.ReadNoticeReq) (out *system.ReadNoticeRes, err error) {
	out = &system.ReadNoticeRes{}
	info, err := service.SystemNotice().GetById(ctx, in.Id)
	if err != nil {
		return
	}
	out.Data = *info
	return
}

func (c *noticeController) Update(ctx context.Context, in *system.UpdateNoticeReq) (out *system.UpdateNoticeRes, err error) {
	err = dao.SystemNotice.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.UpdateNoticeRes{}
		err = service.SystemNotice().Update(ctx, &in.SystemNoticeUpdate)
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

func (c *noticeController) Delete(ctx context.Context, in *system.DeleteNoticeReq) (out *system.DeleteNoticeRes, err error) {
	out = &system.DeleteNoticeRes{}
	err = service.SystemNotice().Delete(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}

func (c *noticeController) RealDelete(ctx context.Context, in *system.RealDeleteNoticeReq) (out *system.RealDeleteNoticeRes, err error) {
	err = dao.SystemNotice.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.RealDeleteNoticeRes{}
		err = service.SystemNotice().RealDelete(ctx, in.Ids)
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

func (c *noticeController) Recovery(ctx context.Context, in *system.RecoveryNoticeReq) (out *system.RecoveryNoticeRes, err error) {
	out = &system.RecoveryNoticeRes{}
	err = service.SystemNotice().Recovery(ctx, in.Ids)
	if err != nil {
		return
	}
	return
}
