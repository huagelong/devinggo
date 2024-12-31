// Package api
// @Link  https://github.com/huagelong/devinggo
// @Copyright  Copyright (c) 2024 devinggo
// @Author  Kai <hpuwang@gmail.com>
// @License  https://github.com/huagelong/devinggo/blob/master/LICENSE

package controller

import (
	"devinggo/internal/controller/base"
	"devinggo/modules/api/api"
	"devinggo/modules/system/pkg/worker/cron"
	"devinggo/modules/system/service"
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var (
	TestController = testController{}
)

type testController struct {
	base.BaseController
}

func (c *testController) Index(ctx context.Context, in *api.IndexReq) (out *api.IndexRes, err error) {
	out = &api.IndexRes{}
	user, _ := service.SystemUser().GetInfoById(ctx, 1)
	g.Log().Debug(ctx, user)
	cron.GetWorkerLists()
	ls := cron.GetWorkerLists()
	g.Log().Debug(ctx, ls)
	out.Data = "Hello, World!"
	return
}
