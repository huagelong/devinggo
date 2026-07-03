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
	DataMaintainController = dataMaintainController{}
)

type dataMaintainController struct {
	base.BaseController
}

func (c *dataMaintainController) Index(ctx context.Context, in *system.IndexDataMaintainReq) (out *system.IndexDataMaintainRes, err error) {
	out = &system.IndexDataMaintainRes{}
	items, totalCount, err := service.DataMaintain().GetPageListForSearch(ctx, &in.PageListReq, &in.DataMaintainSearch)
	if err != nil {
		return
	}

	if !g.IsEmpty(items) {
		for _, item := range items {
			out.Items = append(out.Items, *item)
		}
	} else {
		out.Items = make([]res.DataMaintain, 0)
	}
	out.PageRes.Pack(in, totalCount)
	return
}

func (c *dataMaintainController) Detailed(ctx context.Context, in *system.DetailedDataMaintainReq) (out *system.DetailedDataMaintainRes, err error) {
	out = &system.DetailedDataMaintainRes{}
	out.Items, err = service.DataMaintain().GetColumnDetailList(ctx, in.GroupName, in.TableName)
	if err != nil {
		return
	}
	if g.IsEmpty(out.Items) {
		out.Items = make([]res.DataMaintainColumn, 0)
	}
	return
}
