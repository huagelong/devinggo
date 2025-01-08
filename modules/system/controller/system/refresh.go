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
	"devinggo/modules/system/pkg/utils/request"
	"devinggo/modules/system/service"
)

var (
	RefreshController = refreshController{}
)

type refreshController struct {
	base.BaseController
}

func (c *refreshController) Refresh(ctx context.Context, req *system.RefreshReq) (res *system.RefreshRes, err error) {
	res = &system.RefreshRes{}
	token, expire, err := service.Token().Refresh(request.GetHttpRequest(ctx))
	if err != nil {
		return
	}
	res.Token = token
	res.Expire = expire
	return
}
