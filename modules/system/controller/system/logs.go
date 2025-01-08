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
	"devinggo/modules/system/service"
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
