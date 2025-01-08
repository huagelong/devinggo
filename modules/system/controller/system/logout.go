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
	LogoutController = logoutController{}
)

type logoutController struct {
	base.BaseController
}

func (c *logoutController) Logout(ctx context.Context, req *system.LogoutReq) (res *system.LogoutRes, err error) {
	err = service.Token().Logout(request.GetHttpRequest(ctx))
	return
}
