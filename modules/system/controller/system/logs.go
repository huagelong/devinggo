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
	LogsController = logsController{}
)

type logsController struct {
	base.BaseController
}

func (c *logsController) GetLoginLogPageList(ctx context.Context, in *system.GetLoginLogPageListReq) (out *system.GetLoginLogPageListRes, err error) {
	out = &system.GetLoginLogPageListRes{}
	items, totalCount, err := service.SystemLoginLog().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemLoginLogSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemLoginLog, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *logsController) GetOperLogPageList(ctx context.Context, in *system.GetOperLogPageListReq) (out *system.GetOperLogPageListRes, err error) {
	out = &system.GetOperLogPageListRes{}
	items, totalCount, err := service.SystemOperLog().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemOperLogSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemOperLog, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *logsController) GetApiLogPageList(ctx context.Context, in *system.GetApiLogPageListReq) (out *system.GetApiLogPageListRes, err error) {
	out = &system.GetApiLogPageListRes{}
	items, totalCount, err := service.SystemApiLog().GetPageListForSearch(ctx, &in.PageListReq, &in.SystemApiLogSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.SystemApiLog, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *logsController) DeleteOperLog(ctx context.Context, in *system.DeleteOperLogReq) (out *system.DeleteOperLogRes, err error) {
	err = dao.SystemOperLog.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.DeleteOperLogRes{}
		err = service.SystemOperLog().DeleteOperLog(ctx, in.Ids)
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

func (c *logsController) DeleteApiLog(ctx context.Context, in *system.DeleteApiLogReq) (out *system.DeleteApiLogRes, err error) {
	err = dao.SystemApiLog.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.DeleteApiLogRes{}
		err = service.SystemApiLog().DeleteApiLog(ctx, in.Ids)
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

func (c *logsController) DeleteLoginLog(ctx context.Context, in *system.DeleteLoginLogReq) (out *system.DeleteLoginLogRes, err error) {
	err = dao.SystemLoginLog.Transaction(ctx, func(ctx context.Context, tx gdb.TX) (err error) {
		out = &system.DeleteLoginLogRes{}
		err = service.SystemLoginLog().DeleteLoginLog(ctx, in.Ids)
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
